package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/eth"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/manager"
)

const registerShortMsg = "Register trust node in the configured NxGenT smart contract"

var registerCmd = &cobra.Command{
	Use:                   "register",
	Short:                 registerShortMsg,
	Long:                  title + "\n\n" + "Info:\n  " + registerShortMsg,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize and configure trust manager
		_, auth, tmi := manager.Init(ctx, args[0], false, false, "")

		// Register trust node
		if !eth.IsTrustNodeRegistered(tmi, auth.From) {
			ip := manager.GetNodeIP()
			//ip := net.ParseIP(args[1]) // Experiments
			eth.RegisterTrustNode(auth, tmi, ip)

			// Debug
			fmt.Println("--> Trust node registered:", auth.From, "("+ip.String()+")")
		} else {
			// Debug
			fmt.Println("--> Trust node is already registered")
		}
	},
}
