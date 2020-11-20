/*
Blockchain


*/


package main

import (
	"os"

	"github.com/AmirYunus/go_blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
