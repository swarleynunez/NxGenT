package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorShortMsg = "Only for monitoring/testing purposes"

var monitorCmd = &cobra.Command{
	Use:                   "monitor",
	Short:                 monitorShortMsg,
	Long:                  title + "\n\n" + "Info:\n  " + monitorShortMsg,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		for {
			b, err := os.ReadFile(args[1])
			utils.CheckError(err, utils.ErrorMode)

			if strings.Contains(string(b), "Starting experiment...") {
				pid, err := strconv.ParseInt(args[0], 10, 32)
				utils.CheckError(err, utils.ErrorMode)

				p, err := process.NewProcess(int32(pid))
				utils.CheckError(err, utils.ErrorMode)

				var cpuUsages []float64
				var memUsages []uint64

				for {
					cpu, err := p.CPUPercent()
					utils.CheckError(err, utils.ErrorMode)

					mem, err := p.MemoryInfo()
					utils.CheckError(err, utils.ErrorMode)

					cpuUsages = append(cpuUsages, cpu)
					memUsages = append(memUsages, mem.RSS)

					b, err := os.ReadFile(args[1])
					utils.CheckError(err, utils.ErrorMode)

					if strings.Contains(string(b), "Experiment completed!") {
						var cpuTotal float64
						for _, usage := range cpuUsages {
							cpuTotal += usage
						}
						cpuMean := cpuTotal / float64(len(cpuUsages))

						var memTotal uint64
						for _, usage := range memUsages {
							memTotal += usage
						}
						memMean := memTotal / uint64(len(memUsages))

						//fmt.Println(args[2], cpuMean, memMean)
						fmt.Println(cpuMean, memMean)
						return
					}

					time.Sleep(1000 * time.Millisecond)
				}
			}

			time.Sleep(100 * time.Millisecond)
		}
	},
}
