package main

import (
	"github.com/kataras/cli"
)

func main() {
	cli.App{"httpserver", "converts current directory into http server", "0.0.1", cli.Commands{
		cli.NewCommand("--help", "show help").Add(cli.NewCommand("--host", "show help for host")),
	}}.Run()

}
