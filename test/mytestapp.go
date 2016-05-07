package main

import (
	"strconv"

	"github.com/kataras/cli"
)

func main() {
	app := cli.NewApp("httpserver", "converts current directory into http server", "0.0.1")

	app.Flag("directory", "C:/users/myfiles", "specify a current working directory")

	listenCommand := cli.Command("listen", "starts the server")

	listenCommand.Flag("host", "127.0.0.1", "specify an address listener")
	listenCommand.Flag("port", 8080, "specify a port to listen")
	listenCommand.Flag("dir", "", "current working directory")
	listenCommand.Flag("req", nil, "a required flag because nil default given")
	listenCommand.Flag("key", "", "key file for https")

	app.Run(run)
}

func run(args cli.Flags) error {
	println("Running ONLY APP with -d = " + args.String("directory"))
	return nil
}

func listen(args cli.Flags) error {
	println("EXECUTE ONLY 'listen' with args\n1 host: ", args.String("host"))
	println("2 port: ", strconv.Itoa(args.Int("port")))
	return nil
}

/* OR */
/*
func main() {

	var globalFlags = cli.Flags{
		cli.Flag("directory", "C:/users/myfiles", "specify a current working directory"),
	}

	cli.App{"httpserver", "converts current directory into http server", "0.0.1", cli.Commands{
		cli.Command("listen", "starts the server").
			Flag("host", "127.0.0.1", "specify an address listener").
			Flag("port", 8080, "specify a port to listen").
			Flag("req", nil, "a required flag because no default given").
			Action(listen),
	}, globalFlags}.Run(app)

}

func app(args cli.Flags) error {
	println("Running ONLY APP with -d = " + args.String("directory"))
	return nil
}

func listen(args cli.Flags) error {
	println("EXECUTE ONLY 'listen' with args\n1 host: ", args.String("host"))
	println("2 port: ", strconv.Itoa(args.Int("port")))
	return nil
}
*/
