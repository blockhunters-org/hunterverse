package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/blockhunters-org/hunterverse/app"
	"github.com/blockhunters-org/hunterverse/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"hunterverse",
		"cosmos",
		app.DefaultNodeHome,
		"hunterverse-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
