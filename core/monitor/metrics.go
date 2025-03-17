package monitor

import (
	probing "github.com/prometheus-community/pro-bing"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/swarleynunez/NxGenT/core/types"
	"github.com/swarleynunez/NxGenT/core/utils"
	"math/rand"
	"net"
	"time"
)

func monitorLatency(target net.IP, interval uint64, rpcport string) {

	// Infinite loop
	for {
		// Set and configure pinger
		pinger, err := probing.NewPinger(target.String())
		//pinger, err := probing.NewPinger("127.0.0.1") // Experiments
		pinger.Timeout = time.Duration(interval) * time.Second
		pinger.Count = 1
		if err != nil {
			utils.CheckError(err, utils.WarningMode)
			break
		}

		// Run pinger (blocks until finished)
		err = pinger.Run()
		if err == nil {
			// Set evidence and metrics
			stats := pinger.Statistics()
			evi := types.RPCEvidenceInfo{
				LostPacket:    stats.PacketsSent-stats.PacketsRecv > 0,
				PacketLatency: stats.AvgRtt.Seconds() * 1000,
			}

			// Send evidence to the manager (via RPC server)
			sendEvidence(types.LatencyEvidence, target, &evi, rpcport)

			// Synchronize monitor
			if stats.AvgRtt > 0 {
				time.Sleep(time.Duration(interval) * time.Second)
			}
		} else {
			utils.CheckError(err, utils.WarningMode)
			break
		}
	}
}

func monitorResources(target net.IP, interval uint64, rpcport string) {

	// Infinite loop
	for {
		// Get total available resources
		ci, err := cpu.Info()
		//utils.CheckError(err, utils.WarningMode)
		cores := map[uint64]float64{}
		for _, core := range ci {
			//cores[uint64(core.CPU)] = core.Mhz
			cores[uint64(core.CPU)] = 3300 // TODO: Debug
		}

		vm, err := mem.VirtualMemory()
		utils.CheckError(err, utils.WarningMode)

		du, err := disk.Usage("/") // File system root path
		utils.CheckError(err, utils.WarningMode)

		// Set evidence and metrics
		evi := types.RPCEvidenceInfo{
			CPUCores:    cores,
			MemoryTotal: vm.Total,
			DiskTotal:   du.Total,
		}

		// Send evidence to the manager (via RPC server)
		sendEvidence(types.ResourcesEvidence, target, &evi, rpcport)

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func simulateIDSAlert(target net.IP, interval uint64, rpcport string) {

	// Infinite loop
	for {
		if rand.Intn(4) == 3 { // 25% chance
			// Set evidence and metrics
			evi := types.RPCEvidenceInfo{
				IDSAlert: true,
			}

			// Send evidence to the manager (via RPC server)
			sendEvidence(types.IDSEvidence, target, &evi, rpcport)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
