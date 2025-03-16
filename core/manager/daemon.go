package manager

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/bindings"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/eth"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/inputs"
	"math/big"
	"time"
)

func RunTrustManager() {

	rtns := eth.GetRegisteredTrustNodes(_tmi)

	// Experiments (mixing registered trust nodes)
	/*for i := len(rtns) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		rtns[i], rtns[j] = rtns[j], rtns[i]
	}*/

	// Watchers to receive blockchain events
	for _, addr := range rtns {
		if addr != _auth.From {
			go WatchNewSLA(addr)
		}
		go WatchAcceptedSLA(addr)
	}

	// Time to set up watchers
	time.Sleep(3 * time.Second)

	//i := 0                               // Experiments
	//ni := rand.Intn((len(rtns) / 2) + 1) // Experiments
	for _, addr := range rtns {
		if addr != _auth.From {
			// Experiments
			/*if i >= ni {
				break
			}*/

			// Setting SLAs
			err := eth.SetSLA(_ethc, _auth, _tmi, addr, inputs.SLA1)
			utils.CheckError(err, utils.WarningMode)
			//i++ // Experiments
		}
	}

	// Experiments
	//ctx := context.Background()
	//time.Sleep(30 * time.Second)
	//fmt.Print("[", time.Now().Unix(), "] Starting experiment...\n")
	//firstBlock, _ := _ethc.BlockByNumber(ctx, nil)
	//fmt.Print("[", time.Now().Unix(), "] First block: ", firstBlock.Number(), firstBlock.Time(), "\n")
	//time.Sleep((180 - 1) * time.Second)
	//fmt.Print("[", time.Now().Unix(), "] Experiment completed!\n")
	//lastBlock, _ := _ethc.BlockByNumber(ctx, nil)
	//fmt.Print("[", time.Now().Unix(), "] Last block: ", lastBlock.Number(), lastBlock.Time(), "\n")
	//getBlockchainStorage(ctx, firstBlock, lastBlock)

	/*for {
		r := utils.ConvertBigIntToFloat(eth.GetReputationScore(_tmi, _auth.From))
		if r > 0 {
			fmt.Println(r)
		}
		time.Sleep(10 * time.Second)
	}*/

	// Sleep main goroutine forever
	select {}
}

func runEvidenceManager(target common.Address, slaId uint64) {

	// Get parameters
	epTime := eth.GetNetworkEpochTime(_tmi)
	tni, err := eth.GetTrustNodeInstance(_ethc, _tmi, target)
	if err == nil {
		// Pending evidence metrics
		var pms []bindings.TypesEvidenceMetric

		// Experiments
		var counter uint64

		// Infinite loop
		for {
			time.Sleep(time.Duration(epTime) * time.Second)

			// Check if there are pending evidence metrics
			var ms []bindings.TypesEvidenceMetric
			if len(pms) == 0 {
				// Experiments
				counter++

				// Aggregate local evidences
				ms = parseEvidenceMetrics(aggregateEpochEvidences(target, counter), tni, slaId)
			} else {
				ms = pms
				pms = []bindings.TypesEvidenceMetric{}
			}

			// Send an aggregated evidence to the blockchain
			if len(ms) > 0 {
				err = eth.SendEvidence(_ethc, _auth, _tmi, target, slaId, ms)
				if err == nil {
					// Debug
					//fmt.Print("[", time.Now().Unix(), "] Evidence sent for ", target, " (", eth.GetTrustNodeIP(tni), ")\n")
				} else {
					pms = ms
					utils.CheckError(err, utils.WarningMode)
				}
			}
		}
	} else {
		utils.CheckError(err, utils.WarningMode)
	}
}

// Experiments
func getBlockchainStorage(ctx context.Context, firstBlock *types.Block, lastBlock *types.Block) {

	var tsize uint64
	for i := firstBlock.Number().Uint64(); i <= lastBlock.Number().Uint64(); i++ {
		block, _ := _ethc.BlockByNumber(ctx, new(big.Int).SetUint64(i))
		if len(block.Transactions()) > 0 {
			tsize += block.Size()
		}
	}

	fmt.Println("TOTAL:", tsize)
}
