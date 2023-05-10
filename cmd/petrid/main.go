package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/petri-labs/petri/app"
)

func main() {
	config := app.GetDefaultConfig()
	config.Seal()

	rootCmd, _ := NewRootCmd()

	rootCmd.AddCommand(AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
