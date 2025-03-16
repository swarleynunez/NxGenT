package eth

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/bindings"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
	"net"
)

const (
	SendEvidenceAction = "sendEvidence"
	SetSLAAction       = "setSLA"
	AcceptSLAAction    = "acceptSLA"
)

func DeployTrustManagerContract(ethc *ethclient.Client, auth bind.TransactOpts, epTime, srTime uint64) (tmi *bindings.TrustManager) {

	// Configure transactor
	auth.Nonce = GetNextNonce()

	// Send txn
	addr, _, tmi, err := bindings.DeployTrustManager(&auth, ethc, epTime, srTime)
	utils.CheckError(err, utils.ErrorMode)

	// Persist contract address
	utils.SetEnv("TM_CONTRACT", addr.String())

	return
}

func RegisterTrustNode(auth bind.TransactOpts, tmi *bindings.TrustManager, ip net.IP) {

	// Configure transactor
	auth.Nonce = GetNextNonce()

	// Send txn
	_, err := tmi.RegisterTrustNode(&auth, ip.String())
	utils.CheckError(err, utils.ErrorMode)
}

func SetSLA(ethc *ethclient.Client, auth bind.TransactOpts, tmi *bindings.TrustManager, customer common.Address, metrics []bindings.TypesSLAMetric) error {

	// Get my trust node contract instance (provider)
	tni, err := GetTrustNodeInstance(ethc, tmi, auth.From)
	if err != nil {
		return err
	}

	// Checking zone
	if auth.From != customer && IsTrustNodeRegistered(tmi, customer) && len(metrics) > 0 {
		ms := make(map[string]bool)
		for _, metric := range metrics {
			if !existSLAMetric(tmi, metric.Id) || ms[metric.Id] {
				return errors.New(SetSLAAction + ": transaction not sent")
			}
			ms[metric.Id] = true
		}

		// Configure transactor
		auth.Nonce = GetNextNonce()

		// Send txn
		_, err = tni.SetSLA(&auth, customer, metrics)

		return err
	} else {
		return errors.New(SetSLAAction + ": transaction not sent")
	}
}

func AcceptSLA(ethc *ethclient.Client, auth bind.TransactOpts, tmi *bindings.TrustManager, provider common.Address, slaId uint64) error {

	// Get provider contract instance
	tni, err := GetTrustNodeInstance(ethc, tmi, provider)
	if err != nil {
		return err
	}

	// Checking zone
	if !canAcceptSLA(tni, slaId, auth.From) {
		return errors.New(AcceptSLAAction + ": transaction not sent")
	}

	// Configure transactor
	auth.Nonce = GetNextNonce()

	// Send txn
	_, err = tni.AcceptSLA(&auth, slaId)

	return err
}

func SendEvidence(ethc *ethclient.Client, auth bind.TransactOpts, tmi *bindings.TrustManager, target common.Address, slaId uint64, metrics []bindings.TypesEvidenceMetric) error {

	// Get target contract instance
	tni, err := GetTrustNodeInstance(ethc, tmi, target)
	if err != nil {
		return err
	}

	// Checking zone
	if auth.From != target &&
		IsTrustNodeRegistered(tmi, auth.From) &&
		IsTrustNodeRegistered(tmi, target) &&
		isSLACustomer(tni, slaId, auth.From) &&
		isSLAActive(tni, slaId) &&
		//CanSendEvidence(tmi, auth.From, target, slaId) &&
		len(metrics) > 0 {
		ms := make(map[string]bool)
		for _, metric := range metrics {
			if !existSLAMetric(tmi, metric.Id) || !HasSLAMetric(tni, slaId, metric.Id) || ms[metric.Id] {
				return errors.New(SendEvidenceAction + ": transaction not sent")
			}
			ms[metric.Id] = true
		}

		// Configure transactor
		auth.Nonce = GetNextNonce()

		// Send txn
		_, err = tmi.SendEvidence(&auth, target, slaId, metrics)

		return err
	} else {
		return errors.New(SendEvidenceAction + ": transaction not sent")
	}
}
