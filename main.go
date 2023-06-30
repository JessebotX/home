package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const prefix = "home"

func main() {
	name := filepath.Base(strings.TrimSuffix(os.Args[0], ".exe"))
	args := os.Args[1:]

	if name == prefix {
		if len(args) <= 0 {
			return
		}

		name = os.Args[1]
		args = os.Args[2:]
	}

	fmt.Println(name, args)
}
