package main

import (
	// "github.com/jessebotx/home/cmd/ex"
	"github.com/jessebotx/home/cmd/hello"
	"github.com/jessebotx/home/common"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const program = "home"

var commands = map[string]*common.Command{
	"hello":   hello.Cmd,
	// "ex":      ex.Cmd,
	// "extract": ex.Cmd,
}

func main() {
	// assume program called as `<command> [subcommand]`
	name := filepath.Base(strings.TrimSuffix(os.Args[0], ".exe"))
	args := os.Args[1:]

	// check if program called as `home <command> [subcommand]`
	if name == program {
		if len(args) <= 0 {
			log.Fatal("Missing command")
		}

		name = os.Args[1]
		args = os.Args[2:]

	}

	if name == "help" {
		usage(args)
		return
	}

	cmd, ok := commands[name]
	if !ok {
		log.Fatalf("Unrecognized command `%v`\n", name)
	}

	cmd.Run(args)
}

func usage(args []string) {
	if len(args) <= 0 {
		log.Fatal("Missing arguments.")
	}

	name := args[0]
	args = args[1:]

	cmd, ok := commands[name]
	if !ok {
		log.Fatal("Missing arguments.")
	}

	cmd.Usage()
}
