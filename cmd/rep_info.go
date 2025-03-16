package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/eth"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/manager"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
)

const repInfoShortMsg = "Get the current reputation of a trust node (by EOA)"

var repInfoCmd = &cobra.Command{
	Use:                   "rep-info",
	Short:                 repInfoShortMsg,
	Long:                  title + "\n\n" + "Info:\n  " + repInfoShortMsg,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize and configure trust manager
		ethc, _, tmi := manager.Init(ctx, args[0], false, true, "")

		// Get target's reputation info
		if len(args) == 2 && utils.ValidEthAddress(args[1]) {
			target := common.HexToAddress(args[1])
			tni, err := eth.GetTrustNodeInstance(ethc, tmi, target)
			if err == nil {
				fmt.Println("--> Total evidences (txns):", eth.GetEvidencesCount(tni))
				fmt.Println("--> Reputation (Ri):", utils.ConvertBigIntToFloat(eth.GetReputationScore(tmi, target)))

				// Get subjective opinions
				for _, addr := range eth.GetRegisteredTrustNodes(tmi) {
					if addr == target {
						continue
					}
					dtsi := eth.GetDirectTrustScoreInfo(tmi, addr, target)
					fmt.Print("	- ", addr, " -->",
						" α:", utils.ConvertBigIntToFloat(dtsi.A),
						" β:", utils.ConvertBigIntToFloat(dtsi.B),
						" Nα:", dtsi.Na,
						" Nβ:", dtsi.Nb,
						" DTS:", utils.ConvertBigIntToFloat(dtsi.Dts),
						" Rj:", utils.ConvertBigIntToFloat(dtsi.R),
						" C':", utils.ConvertBigIntToFloat(dtsi.Cp), "\n")

					//na, nb := eth.GetCurrentIntervalNCount(tmi, addr, target)
					//fmt.Println(na, nb, dtsi.NCounts)
				}
			}
		}
	},
}
