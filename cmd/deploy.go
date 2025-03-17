package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarleynunez/NxGenT/core/manager"
	"github.com/swarleynunez/NxGenT/core/utils"
)

const deployShortMsg = "Deploy and configure a new NxGenT smart contract"

var deployCmd = &cobra.Command{
	Use:                   "deploy",
	Short:                 deployShortMsg,
	Long:                  title + "\n\n" + "Info:\n  " + deployShortMsg,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize and configure trust manager
		manager.Init(ctx, args[0], true, false, "")

		// Debug
		fmt.Println("--> NxGenT smart contract:", utils.GetEnv("TM_CONTRACT"))
	},
}
