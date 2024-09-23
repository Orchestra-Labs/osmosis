package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	app "github.com/osmosis-labs/osmosis/v26/app"
	"github.com/osmosis-labs/osmosis/v26/app/params"
	"github.com/osmosis-labs/osmosis/v26/cmd/symphonyd/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "SYMPHONYD", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
