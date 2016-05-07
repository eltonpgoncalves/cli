## CLI

Build small command line apps with ease

## Usage

```go

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

```

### Second approach

```go

package main

import (
	"strconv"
	"github.com/kataras/cli"
)

func main() {

	var globalFlags = cli.Flags{
		cli.Flag("directory", "C:/users/myfiles", "specify a current working directory"),
	}

	cli.App{"httpserver", "converts current directory into http server", "0.0.1", cli.Commands{
		cli.Command("listen", "starts the server").
			Flag("host", "127.0.0.1", "specify an address listener").
			Flag("port", 8080, "specify a port to listen").
			Flag("dir", "", "current working directory"). // It's defaults to empty string it is not required flag
			Flag("req", nil, "a required flag because nil default given"). // required flag because nil default given
			Action(listen),
	}, globalFlags}.Run(run)

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

```



## Versioning

Current: **v0.0.1**


Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions


## Contributors

Thanks goes to the people who have contributed code to this package, see the

- [Contributors page](https://github.com/kataras/cli/graphs/contributors).


## License

This project is licensed under the [BSD 3-Clause License](https://opensource.org/licenses/BSD-3-Clause).
License can be found [here](https://github.com/kataras/iris/blob/master/LICENSE).
