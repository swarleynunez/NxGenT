package main

import (
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/cmd"
	"gitlab.uoclabs.uoc.es/kison/6GENABLERS/core/utils"
)

func main() {

	// CLI entrypoint
	err := cmd.Execute()
	utils.CheckError(err, utils.ErrorMode)
}
