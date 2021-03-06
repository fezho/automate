package main

import (
	"fmt"
	"os"

	"github.com/chef/automate/components/infra-proxy-service/cmd/infra-proxy-service/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
