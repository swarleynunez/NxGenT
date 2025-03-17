package main

import (
	"github.com/swarleynunez/NxGenT/cmd"
	"github.com/swarleynunez/NxGenT/core/utils"
)

func main() {

	// CLI entrypoint
	err := cmd.Execute()
	utils.CheckError(err, utils.ErrorMode)
}
