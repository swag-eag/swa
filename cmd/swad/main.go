package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/swag-eag/swa/v2/app"
	"github.com/swag-eag/swa/v2/cmd/swad/cmd"
)

func main() {
	// Watch os.Args to see the command line arguments
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, cmd.EnvPrefix, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
