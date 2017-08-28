package main

import (
	"os"

	"github.com/bytom/cmd/blockchain/commands"
	"github.com/tendermint/tmlibs/cli"
)

func main() {
	cmd := cli.PrepareBaseCmd(commands.RootCmd, "TM", os.ExpandEnv("./.blockchain"))
	cmd.Execute()
}