package hello

import (
	"github.com/jessebotx/home/common"
	"fmt"
)

var Cmd = &common.Command{
	Run: Run,
	Usage: Usage,
}

func Run(args []string) {
	fmt.Println("Hello, world!")
}

func Usage() {
	fmt.Print(
		"COMMAND\n",
		"    hello\n",
		"\n",
		"DESCRIPTION\n",
		"    Prints `Hello, world!`\n",
	)
}
