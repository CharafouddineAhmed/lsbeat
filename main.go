package main

import (
	"os"

	"github.com/CharafouddineAhmed/lsbeat/cmd"

	_ "github.com/CharafouddineAhmed/lsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
