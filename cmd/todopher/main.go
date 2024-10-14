package main

import (
	"fmt"
	"os"

	"github.com/nendix/Todopher/cmd/todopher/cli"
	"github.com/nendix/Todopher/cmd/todopher/tui"
)

func main() {
	if len(os.Args) < 2 {
		if err := tui.StartTUI(); err != nil {
			fmt.Println("Error launching TUI:", err)
			os.Exit(1)
		}
		return
	}
	cli.HandleCLI()
}
