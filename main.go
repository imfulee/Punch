package main

import (
	"fmt"
	"os"

	"github.com/imfulee/punch/cmd"
)

func main() {
	cmdRoot := cmd.CmdRoot()

	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
