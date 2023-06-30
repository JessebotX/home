package ex

import (
	"fmt"
	"github.com/jessebotx/home/common"
)

var Cmd = &common.Command{
	Run:   Run,
	Usage: Usage,
}

func Run(args []string) {
	fmt.Println("Hello world")
}

func Usage() {
	fmt.Println("ex USAGE.")
}
