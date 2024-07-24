package main

import (
	"os"

  "gitlab.com/tokene/lockton-one/rpc-wrapper-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
