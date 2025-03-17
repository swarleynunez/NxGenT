package manager

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swarleynunez/NxGenT/core/bindings"
	"github.com/swarleynunez/NxGenT/core/eth"
	"github.com/swarleynunez/NxGenT/core/types"
	"github.com/swarleynunez/NxGenT/core/utils"
	"github.com/swarleynunez/NxGenT/inputs"
	"math"
	"net"
	"reflect"
	"strconv"
	"sync"
)

var (
	// Unexported and read-only global variables
	_ethc *ethclient.Client
	_auth bind.TransactOpts
	_tmi  *bindings.TrustManager

	// Unexported and writable global variables
	_evs sync.Map // Store for local evidences

	// Experiments
	_rpcport string

	// Errors
	errUnknownEvidenceType = errors.New("unknown evidence type")
	errUnknownTrustNode    = errors.New("unknown trust node")
)

func Init(ctx context.Context, nodeDir string, deploying, testing bool, rpcport string) (*ethclient.Client, bind.TransactOpts, *bindings.TrustManager) {

	// Load environment variables
	utils.LoadEnv()
	epTime, err := strconv.ParseUint(utils.GetEnv("EPOCH_TIME"), 10, 64)
	utils.CheckError(err, utils.ErrorMode)
	srTime, err := strconv.ParseUint(utils.GetEnv("SEARCH_RANGE_TIME"), 10, 64)
	utils.CheckError(err, utils.ErrorMode)

	// Connect to an Ethereum node
	if testing {
		_ethc = eth.Connect("http://localhost:8545")
	} else {
		_ethc = eth.Connect(utils.FormatPath(nodeDir, "geth.ipc"))
	}

	// Load account and create transactor
	ks := eth.LoadKeystore(utils.FormatPath(nodeDir, "keystore"))
	_auth = *eth.Transactor(ctx, _ethc, ks, eth.LoadAccount(ks, utils.GetEnv("ETH_NODE_PASS")))

	// Get trust manager instance
	if deploying {
		_tmi = eth.DeployTrustManagerContract(_ethc, _auth, epTime, srTime)
	} else {
		_tmi = eth.GetTrustManagerInstance(_ethc)
	}

	// Experiments
	_rpcport = rpcport

	return _ethc, _auth, _tmi
}

func aggregateEpochEvidences(addr common.Address, counter uint64) *types.EvidenceInfo {

	// Get local evidences
	evs := getEpochEvidences(addr)
	if len(evs) == 0 {
		return &types.EvidenceInfo{}
	}

	// Clear local evidences for the next epoch
	_evs.Store(addr, []*types.RPCEvidence{})

	// Aggregate local evidences
	totalPkts := 0
	var lats []float64
	evi := types.EvidenceInfo{}
	for _, ev := range evs {
		switch ev.Type {
		case types.LatencyEvidence:
			totalPkts++
			if !ev.Info.LostPacket {
				lats = append(lats, ev.Info.PacketLatency)
			}
		case types.ResourcesEvidence: // Once per epoch
			cores := 0.0
			freq := 0.0
			for _, f := range ev.Info.CPUCores {
				cores++
				freq += f / 1000
			}
			evi.CPUCores = cores
			evi.CPUFrequency = freq / cores
			evi.MemoryTotal = float64(ev.Info.MemoryTotal) / 1024 / 1024 / 1024
			evi.DiskTotal = float64(ev.Info.DiskTotal) / 1024 / 1024 / 1024
		case types.IDSEvidence: // Once per epoch
			if ev.Info.IDSAlert {
				evi.IDSAlert = 1
			}
		default:
			utils.CheckError(errUnknownEvidenceType, utils.WarningMode)
		}
	}

	if len(lats) > 0 {
		// Aggregate data to calculate availability
		evi.Availability = (float64(len(lats)) / float64(totalPkts)) * 100

		// Aggregate data to calculate average latency
		var totalLat float64
		for _, lat := range lats {
			totalLat += lat
		}
		evi.Latency = totalLat / float64(len(lats))

		// Aggregate data to calculate jitter
		var sd float64 // Squared deviation
		for _, lat := range lats {
			sd += math.Pow(lat-evi.Latency, 2)
		}
		evi.Jitter = math.Sqrt(sd / float64(len(lats)))
	}

	// Experiments
	evi = inputs.EV1

	return &evi
}

func parseEvidenceMetrics(evi *types.EvidenceInfo, tni *bindings.TrustNode, slaId uint64) (metrics []bindings.TypesEvidenceMetric) {

	// For each struct field
	v := reflect.ValueOf(*evi)
	for i := 0; i < v.NumField(); i++ {
		mid := v.Type().Field(i).Tag.Get("json") // Key
		mv := v.Field(i).Interface().(float64)   // Value

		// Check metric
		if eth.HasSLAMetric(tni, slaId, mid) && mv > 0 {
			// Convert metric value to type SD59x18 (Solidity float)
			mstrv := utils.ConvertFloatToBigInt(mv)

			// Save metric
			metrics = append(metrics, bindings.TypesEvidenceMetric{
				Id:    mid,
				Value: mstrv,
			})
		}
	}

	return
}

// Setters //
func StoreEvidenceLocally(ev *types.RPCEvidence) {

	addr := eth.GetTrustNodeAddressFromIP(_tmi, ev.Target) // Get trust node's EOA from its IP
	if !utils.EmptyEthAddress(addr.String()) {
		evs, _ := _evs.Load(addr)
		if evs == nil {
			_evs.Store(addr, []*types.RPCEvidence{ev})
		} else {
			_evs.Store(addr, append(evs.([]*types.RPCEvidence), ev))
		}
	}
}

// Getters //
func GetNodeIP() net.IP {

	conn, err := net.Dial("udp", "192.168.1.1:80")
	utils.CheckError(err, utils.ErrorMode)
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP
}

func getEpochEvidences(addr common.Address) []*types.RPCEvidence {

	evs, _ := _evs.Load(addr)
	if evs != nil {
		return evs.([]*types.RPCEvidence)
	} else {
		return []*types.RPCEvidence{}
	}
}

func GetReputationScore(ip net.IP) (rs float64, err error) {

	addr := eth.GetTrustNodeAddressFromIP(_tmi, ip) // Get trust node's EOA from its IP
	if !utils.EmptyEthAddress(addr.String()) {
		rs = utils.ConvertBigIntToFloat(eth.GetReputationScore(_tmi, addr))
	} else {
		err = errUnknownTrustNode
	}

	return
}

// Experiment_4_a
/*
	if addr.String() == "0xBFe66A4F8b3208c5b780a11407644DEe2132ad28" {
		if counter > 150 {
			var idsa float64
			if rand.Intn(10) > 4 {
				idsa = 1
			}
			evi = types.EvidenceInfo{
				Availability: float64(rand.Intn(103-94+1) + 94),
				Latency:      float64(rand.Intn(20-11+1) + 11),
				Jitter:       float64(rand.Intn(15-6+1) + 6),
				CPUCores:     float64(rand.Intn(14-5+1) + 5),
				CPUFrequency: float64(rand.Intn(14-5+1) + 5),
				MemoryTotal:  float64(rand.Intn(20-11+1) + 11),
				DiskTotal:    float64(rand.Intn(54-45+1) + 45),
				IDSAlert:     idsa,
			}
		} else {
			evi = inputs.EV1
		}
	} else {
		evi = inputs.EV1
	}
*/

// Experiment_4_b
/*
	if addr.String() == "0xBFe66A4F8b3208c5b780a11407644DEe2132ad28" {
		if counter > 150 {
			evi = inputs.EV4_2
		} else {
			evi = inputs.EV1
		}
	} else {
		evi = inputs.EV1
	}
*/

// Experiment_5
/*
	target := addr.String()
	if target == "0xBFe66A4F8b3208c5b780a11407644DEe2132ad28" {
		me := _auth.From.String()
		if me == "0xa4C22Dcd46c387F518C2a7032Db7EBeE1Cf620c1" ||
			me == "0xD5C6257bAf02B484f0aB968FCe16198110846F94" {
			var idsa float64
			if rand.Intn(10) > 1 {
				idsa = 1
			}
			evi = types.EvidenceInfo{
				Availability: float64(rand.Intn(100-91+1) + 91),
				Latency:      float64(rand.Intn(23-14+1) + 14),
				Jitter:       float64(rand.Intn(18-9+1) + 9),
				CPUCores:     float64(rand.Intn(11-2+1) + 2),
				CPUFrequency: float64(rand.Intn(11-2+1) + 2),
				MemoryTotal:  float64(rand.Intn(17-8+1) + 8),
				DiskTotal:    float64(rand.Intn(51-42+1) + 42),
				IDSAlert:     idsa,
			}
		} else {
			evi = inputs.EV1
		}
	} else {
		if target == "0xa4C22Dcd46c387F518C2a7032Db7EBeE1Cf620c1" ||
			target == "0xD5C6257bAf02B484f0aB968FCe16198110846F94" {
			var idsa float64
			if rand.Intn(10) > 4 {
				idsa = 1
			}
			evi = types.EvidenceInfo{
				Availability: float64(rand.Intn(103-94+1) + 94),
				Latency:      float64(rand.Intn(20-11+1) + 11),
				Jitter:       float64(rand.Intn(15-6+1) + 6),
				CPUCores:     float64(rand.Intn(14-5+1) + 5),
				CPUFrequency: float64(rand.Intn(14-5+1) + 5),
				MemoryTotal:  float64(rand.Intn(20-11+1) + 11),
				DiskTotal:    float64(rand.Intn(54-45+1) + 45),
				IDSAlert:     idsa,
			}
		} else {
			evi = inputs.EV1
		}
	}
*/
