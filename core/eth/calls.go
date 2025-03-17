package eth

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swarleynunez/NxGenT/core/bindings"
	"github.com/swarleynunez/NxGenT/core/utils"
	"math/big"
	"net"
)

var (
	ErrMalformedAddr = errors.New("malformed address")
)

// Getters //
func GetTrustManagerInstance(ethc *ethclient.Client) (tmi *bindings.TrustManager) {

	// Get contract address
	addr := utils.GetEnv("TM_CONTRACT")

	// Get contract instance
	if utils.ValidEthAddress(addr) {
		i, err := bindings.NewTrustManager(common.HexToAddress(addr), ethc)
		utils.CheckError(err, utils.ErrorMode)
		tmi = i
	} else {
		utils.CheckError(ErrMalformedAddr, utils.ErrorMode)
	}

	return
}

func GetTrustNodeInstance(ethc *ethclient.Client, tmi *bindings.TrustManager, addr common.Address) (tni *bindings.TrustNode, err error) {

	// Get trust node contract address
	caddr, err := tmi.Nodes(&bind.CallOpts{}, addr)
	if err == nil {
		// Get trust node instance
		tni, err = bindings.NewTrustNode(caddr, ethc)
	}

	return
}

func GetNetworkEpochTime(tmi *bindings.TrustManager) uint64 {

	conf, err := tmi.Config(&bind.CallOpts{})
	utils.CheckError(err, utils.WarningMode)

	return conf.EpochTime
}

// TrustManager getters //
func GetRegisteredTrustNodes(tmi *bindings.TrustManager) (addrs []common.Address) {

	addrs, err := tmi.GetRegisteredTrustNodes(&bind.CallOpts{})
	utils.CheckError(err, utils.WarningMode)

	return
}

func GetTrustNodeAddressFromIP(tmi *bindings.TrustManager, ip net.IP) (addr common.Address) {

	addr, err := tmi.NodeIPs(&bind.CallOpts{}, ip.String())
	utils.CheckError(err, utils.WarningMode)

	return
}

func GetDirectTrustScoreInfo(tmi *bindings.TrustManager, addr, target common.Address) *bindings.TypesDTSInfo {

	ds, err := tmi.GetDTSInfo(&bind.CallOpts{}, addr, target)
	utils.CheckError(err, utils.WarningMode)

	return &ds
}

func GetReputationScore(tmi *bindings.TrustManager, addr common.Address) (r *big.Int) {

	r = new(big.Int)
	if GetReputationScoreCount(tmi, addr) > 1 {
		rs, err := tmi.GetRScore(&bind.CallOpts{}, addr)
		utils.CheckError(err, utils.WarningMode)
		r = rs
	}

	return
}

func GetReputationScoreCount(tmi *bindings.TrustManager, addr common.Address) uint64 {

	rsc, err := tmi.GetRScoreCount(&bind.CallOpts{}, addr)
	utils.CheckError(err, utils.WarningMode)

	return rsc.Uint64()
}

func GetCurrentIntervalNCount(tmi *bindings.TrustManager, addr, target common.Address) (uint64, uint64) {

	nCount, err := tmi.GetCurrentIntervalNCount(&bind.CallOpts{}, addr, target)
	utils.CheckError(err, utils.WarningMode)

	return nCount.NaIn, nCount.NbIn
}

// TrustNode getters //
func GetTrustNodeIP(tni *bindings.TrustNode) (ip net.IP) {

	strIP, err := tni.GetIP(&bind.CallOpts{})
	utils.CheckError(err, utils.WarningMode)

	return net.ParseIP(strIP)
}

func GetLastEvidenceTime(tni *bindings.TrustNode, target common.Address, slaId uint64) uint64 {

	t, err := tni.GetLastEvidenceTime(&bind.CallOpts{}, target, slaId)
	utils.CheckError(err, utils.WarningMode)

	return t.Uint64()
}

func GetEvidencesCount(tni *bindings.TrustNode) (r *big.Int) {

	r, err := tni.GetEvidencesCount(&bind.CallOpts{})
	utils.CheckError(err, utils.WarningMode)

	return
}

/*func GetSLA(tni *bindings.TrustNode, slaId uint64) *types.SLA {

	// Get SLA header
	h, err := tni.SLAs(&bind.CallOpts{}, slaId)
	utils.CheckError(err, utils.WarningMode)

	// Get SLA metrics
	ms, err := tni.GetSLAMetrics(&bind.CallOpts{}, slaId)
	utils.CheckError(err, utils.WarningMode)

	return &types.SLA{
		Customer:     h.Customer,
		Metrics:      ms,
		SetAt:        h.SetAt,
		AcceptedAt:   h.AcceptedAt,
		TerminatedAt: h.TerminatedAt,
	}
}*/

// TrustManager helpers //
func IsTrustNodeRegistered(tmi *bindings.TrustManager, addr common.Address) (r bool) {

	r, err := tmi.IsTrustNodeRegistered(&bind.CallOpts{}, addr)
	utils.CheckError(err, utils.WarningMode)

	return
}

func isIPAvailable(tmi *bindings.TrustManager, ip net.IP) (r bool) {

	r, err := tmi.IsIPAvailable(&bind.CallOpts{}, ip.String())
	utils.CheckError(err, utils.WarningMode)

	return
}

func CanSendEvidence(tmi *bindings.TrustManager, sender, target common.Address, slaId uint64) (r bool) {

	r, err := tmi.CanSendEvidence(&bind.CallOpts{}, sender, target, slaId)
	utils.CheckError(err, utils.WarningMode)

	return
}

func existSLAMetric(tmi *bindings.TrustManager, metricId string) (r bool) {

	r, err := tmi.ExistMetric(&bind.CallOpts{}, metricId)
	utils.CheckError(err, utils.WarningMode)

	return
}

// TrustNode helpers //
func isSLACustomer(tni *bindings.TrustNode, id uint64, addr common.Address) (r bool) {

	r, err := tni.IsSLACustomer(&bind.CallOpts{}, id, addr)
	utils.CheckError(err, utils.WarningMode)

	return
}

func isSLAActive(tni *bindings.TrustNode, id uint64) (r bool) {

	r, err := tni.IsSLAActive(&bind.CallOpts{}, id)
	utils.CheckError(err, utils.WarningMode)

	return
}

func canAcceptSLA(tni *bindings.TrustNode, id uint64, addr common.Address) (r bool) {

	r, err := tni.CanAcceptSLA(&bind.CallOpts{}, id, addr)
	utils.CheckError(err, utils.WarningMode)

	return
}

/*func canTerminateSLA(tni *bindings.TrustNode, id uint64, addr common.Address) (r bool) {

	r, err := tni.CanTerminateSLA(&bind.CallOpts{}, id, addr)
	utils.CheckError(err, utils.WarningMode)

	return
}*/

func HasSLAMetric(tni *bindings.TrustNode, slaId uint64, metricId string) (r bool) {

	r, err := tni.HasSLAMetric(&bind.CallOpts{}, slaId, metricId)
	utils.CheckError(err, utils.WarningMode)

	return
}
