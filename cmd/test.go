package cmd

import (
	"github.com/spf13/cobra"
	"github.com/swarleynunez/NxGenT/core/manager"
)

const testShortMsg = ""

var testCmd = &cobra.Command{
	Use:                   "test",
	Short:                 testShortMsg,
	Long:                  title + "\n\n" + "Info:\n  " + testShortMsg,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize and configure trust manager
		_, _, _ = manager.Init(ctx, args[0], false, true, "")

		// Test 1
		/*tni, err := eth.GetTrustNodeInstance(ethc, tmi, common.HexToAddress("0xBFe66A4F8b3208c5b780a11407644DEe2132ad28"))
		if err == nil {
			fmt.Println(eth.GetSLA(tni, 0))
		}*/

		// Test 2
		/*firstBlock, _ := ethc.BlockByNumber(ctx, nil)
		fmt.Print("[", time.Now().Unix(), "] Init: ", firstBlock.Time(), "\n")
		time.Sleep(time.Duration(180) * time.Second)
		lastBlock, _ := ethc.BlockByNumber(ctx, nil)
		fmt.Print("[", time.Now().Unix(), "] End: ", lastBlock.Time(), "\n")

		var tsize uint64
		for i := firstBlock.Number().Uint64(); i <= lastBlock.Number().Uint64(); i++ {
			block, _ := ethc.BlockByNumber(ctx, new(big.Int).SetUint64(i))
			if len(block.Transactions()) > 0 {
				fmt.Println(block.Number(), block.Size(), len(block.Transactions()))
				tsize += block.Size()
			}
		}

		fmt.Println("TOTAL:", tsize)*/

		// Test 3
		/*msg, err := json.EncodeClientRequest("RPC.GetReputationScore", net.ParseIP("172.18.0.3"))
		if err == nil {
			resp, err := http.Post(
				"http://localhost:"+utils.GetEnv("TM_RPC_PORT")+"/rpc",
				"application/json",
				bytes.NewBuffer(msg),
			)
			utils.CheckError(err, utils.WarningMode)
			defer resp.Body.Close()

			// Debug
			fmt.Println(string(msg))
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		} else {
			utils.CheckError(err, utils.WarningMode)
		}*/
	},
}
