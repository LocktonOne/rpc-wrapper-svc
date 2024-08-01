package main

import (
	"gitlab.com/tokene/lockton-one/rpc-wrapper-svc/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
