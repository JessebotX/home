package main

import (
	"github.com/jessebotx/home/cmd/ex"
	"github.com/jessebotx/home/common"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const program = "home"

var commands = map[string]*common.Command{
	"ex":      ex.Cmd,
	"extract": ex.Cmd,
}

func main() {
	name := filepath.Base(strings.TrimSuffix(os.Args[0], ".exe"))
	args := os.Args[1:]

	// TODO: implement help command for self-documentation
	// eg. home help extract => prints info for command "extract"

	if name == program {
		if len(args) <= 0 {
			log.Fatal("Missing command")
		}

		name = os.Args[1]
		args = os.Args[2:]
	}

	cmd, ok := commands[name]
	if !ok {
		log.Fatalf("Unrecognized command %v\n", name)
	}

	cmd.Run(args)
}
