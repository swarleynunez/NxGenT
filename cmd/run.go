package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/eth"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/manager"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/rpc"
	"net/http"
	"os"
)

const runShortMsg = "Run NxGenT daemon (manager and monitor)"

var runCmd = &cobra.Command{
	Use:                   "run",
	Short:                 runShortMsg,
	Long:                  title + "\n\n" + "Info:\n  " + runShortMsg,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize and configure trust manager
		_, auth, tmi := manager.Init(ctx, args[0], false, false, args[1])

		// Check if trust node is registered
		if !eth.IsTrustNodeRegistered(tmi, auth.From) {
			fmt.Println("--> Trust node not registered")
			os.Exit(0)
		}

		// Debug
		fmt.Print("--> ", "Loaded EOA: ", auth.From, "\n")

		// Run RPC server
		go func() { // Anonymous goroutine
			//err := http.ListenAndServe(":"+utils.GetEnv("TM_RPC_PORT"), rpc.NewRouter())
			err := http.ListenAndServe(":"+args[1], rpc.NewRouter()) // Experiments
			utils.CheckError(err, utils.ErrorMode)
		}()

		// Run manager
		manager.RunTrustManager()
	},
}
