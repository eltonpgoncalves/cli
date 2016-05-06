package main

import (
	"strconv"

	"github.com/kataras/cli"
)

func main() {
	cli.App{"httpserver", "converts current directory into http server", "0.0.1", cli.Commands{
		cli.NewCommand("listen", "starts the server").
			Flag("host", "127.0.0.1", "specify an address listener").
			Flag("port", 8080, "specify a port to listen").
			Flag("req", "", "a required flag because no default given").
			Action(listen),
	}}.Run()

}

func listen(args cli.CommandFlags) error {
	println("EXECUTE 'listen' with args\n1 host: ", args.String("host"))
	println("2 port: ", strconv.Itoa(args.Int("port")))
	return nil
}
