package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("hc-grpc", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"golang": func() (cli.Command, error) {
			return &Golang{}, nil
		},
		//"bar": barCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}

type Golang struct{}

func (*Golang) Help() string {
	return "hi golang"
}

func (h *Golang) Synopsis() string {
	return h.Help()
}

func (*Golang) Run(args []string) int {
	fmt.Printf("hi golang with args: %s", strings.Join(args, ", "))
	return 0
}
