package manager

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/bindings"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/eth"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/monitor"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
	"strconv"
	"time"
)

// Executed by SLA customers
func WatchNewSLA(addr common.Address) {

	// Get trust node contract instance
	tni, err := eth.GetTrustNodeInstance(_ethc, _tmi, addr)
	utils.CheckError(err, utils.WarningMode)

	// Log channel
	logs := make(chan *bindings.TrustNodeNewSLA)

	// Subscription to the event
	sub, err := tni.WatchNewSLA(nil, logs)
	utils.CheckError(err, utils.WarningMode)

	// Cache to avoid duplicated logs
	lcache := map[string]bool{}

	// Infinite loop
	for {
		select {
		case log := <-logs:
			// Check if a log has already been received
			logId := log.From.String() + "_" + strconv.FormatUint(log.Id, 10)
			if !log.Raw.Removed && !lcache[logId] {
				lcache[logId] = true

				// Am I the SLA customer?
				if log.To == _auth.From {
					// Debug
					fmt.Print("[", time.Now().Unix(), "] New SLA --> ID:", log.Id, ", Provider:", log.From.String(), "\n")

					// Anonymous goroutine
					go func() {
						err = eth.AcceptSLA(_ethc, _auth, _tmi, log.From, log.Id)
						utils.CheckError(err, utils.WarningMode)
					}()
				}
			}
		case err = <-sub.Err():
			utils.CheckError(err, utils.WarningMode)
		}
	}
}

// Executed by SLA providers
func WatchAcceptedSLA(addr common.Address) {

	// Get trust node contract instance
	tni, err := eth.GetTrustNodeInstance(_ethc, _tmi, addr)
	utils.CheckError(err, utils.WarningMode)

	// Log channel
	logs := make(chan *bindings.TrustNodeAcceptedSLA)

	// Subscription to the event
	sub, err := tni.WatchAcceptedSLA(nil, logs)
	utils.CheckError(err, utils.WarningMode)

	// Cache to avoid duplicated logs
	lcache := map[string]bool{}

	// Infinite loop
	for {
		select {
		case log := <-logs:
			// Check if a log has already been received
			logId := log.From.String() + "_" + strconv.FormatUint(log.Id, 10)
			if !log.Raw.Removed && !lcache[logId] {
				lcache[logId] = true

				// Am I the SLA customer?
				if log.To == _auth.From {
					// Experiments
					//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
					//fmt.Print("[", time.Now().Unix(), "] Starting...\n")

					// Start monitoring the target (i.e. the provider)
					go monitor.RunTrustMonitor(_ethc, _tmi, log.From, _rpcport)

					// Start evidence manager for the target (i.e. the provider)
					go runEvidenceManager(log.From, log.Id)
				}

				// Am I the SLA provider?
				if log.From == _auth.From {
					// Debug
					fmt.Print("[", time.Now().Unix(), "] Accepted SLA --> ID:", log.Id, ", Customer:", log.To.String(), "\n")

					// Start monitoring the target (i.e. the customer)
					//go monitor.RunTrustMonitor(_ethc, _tmi, log.To)

					// Start evidence manager for the target (i.e. the customer)
					//go runEvidenceManager(log.To, log.Id)
				}
			}
		case err = <-sub.Err():
			utils.CheckError(err, utils.WarningMode)
		}
	}
}
