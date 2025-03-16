package monitor

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/rpc/v2/json"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/bindings"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/eth"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/types"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func RunTrustMonitor(ethc *ethclient.Client, tmi *bindings.TrustManager, target common.Address, rpcport string) {

	// Get target contract instance
	tni, err := eth.GetTrustNodeInstance(ethc, tmi, target)
	if err == nil {
		ip := eth.GetTrustNodeIP(tni)
		epTime := eth.GetNetworkEpochTime(tmi)

		// Metric 1: packet loss + packet latency
		go monitorLatency(ip, 1, rpcport)

		// Metric 2: resources (CPU, memory and disk)
		go monitorResources(ip, epTime, rpcport)

		// Metric 3: IDS alerts
		go simulateIDSAlert(ip, epTime, rpcport)
	} else {
		utils.CheckError(err, utils.WarningMode)
	}
}

func sendEvidence(evt types.EvidenceType, target net.IP, evi *types.RPCEvidenceInfo, rpcport string) {

	msg, err := json.EncodeClientRequest("RPC.SendEvidence", types.RPCEvidence{Type: evt, Target: target, Info: evi})
	if err == nil {
		resp, err := http.Post(
			//"http://localhost:"+utils.GetEnv("TM_RPC_PORT")+"/rpc",
			"http://localhost:"+rpcport+"/rpc",
			"application/json",
			bytes.NewBuffer(msg),
		)
		utils.CheckError(err, utils.WarningMode)
		defer resp.Body.Close()

		// Debug
		/*fmt.Println(string(msg))
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))*/
	} else {
		utils.CheckError(err, utils.WarningMode)
	}
}

// Mock functions //
func RunMockEvidenceGenerator(ethc *ethclient.Client, tmi *bindings.TrustManager, target common.Address, rpcport string) {

	// Get trust node contract instance
	tni, err := eth.GetTrustNodeInstance(ethc, tmi, target)
	if err == nil {
		ip := eth.GetTrustNodeIP(tni)
		fmt.Print("[", time.Now().Unix(), "] Running evidence generator for ", target, " (", ip, ")\n")
		for {
			sendEvidence(types.LatencyEvidence, ip, mockEvidenceInfo(), rpcport)
			time.Sleep(1 * time.Second)
		}
	} else {
		utils.CheckError(err, utils.WarningMode)
	}
}

func mockEvidenceInfo() *types.RPCEvidenceInfo {

	return &types.RPCEvidenceInfo{
		PacketLatency: float64(rand.Intn(50 + 1)),
	}
}
