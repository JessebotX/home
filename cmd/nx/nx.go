package nx

import (
	"fmt"
	"github.com/jessebotx/home/common"
	"log"
	"os"
)

var Cmd = &common.Command{
	Run:   Run,
	Usage: Usage,
}

func Run(args []string) {
	if len(args) <= 0 {
		log.Fatal("Missing arguments")
	}

	fileName := args[0]

	_, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(fileName, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func Usage() {
	fmt.Print(
		"NAME\n",
		"    nx <fileName>\n",
		"\n",
		"DESCRIPTION\n",
		"    Create a new executable shell script named <fileName>\n",
	)
}
