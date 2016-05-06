// Copyright (c) 2016, Gerasimos Maropoulos
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//	  this list of conditions and the following disclaimer
//    in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse
//    or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER AND CONTRIBUTOR, GERASIMOS MAROPOULOS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package cli

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var Output = os.Stdout // the output is the same for all Apps, atm.

func Printf(format string, args ...interface{}) {
	fmt.Fprintf(Output, format, args...)
}

func HelpMe(app App) {
	tmpl, err := template.New(app.Name).Parse(appTmpl)
	if err != nil {
		panic("Panic: " + err.Error())
	}

	tmpl.Execute(Output, app)

}

type App struct {
	Name        string
	Description string
	Version     string
	Commands    Commands
}

func NewApp(name string, description string, version string) *App {
	return &App{name, description, version, nil}
}

// Add adds a  command to the app
func (a *App) Add(cmd *Command) *App {
	if a.Commands == nil {
		a.Commands = Commands{}
	}

	a.Commands = append(a.Commands, cmd)
	return a
}

func (a App) help() {
	HelpMe(a)
	os.Exit(1)
}

func (a App) Run() {
	flagset := flag.NewFlagSet(a.Name, flag.PanicOnError)
	flagset.SetOutput(Output)

	if err := flagset.Parse(os.Args); err != nil {
		a.help()
	}

	if len(flagset.Args()) <= 1 { //for now we only support myapp something -h, no myapp -h - I will fix it tomorrow.
		a.help()
	} else {
		if len(a.Commands) == 0 {
			Printf("No commands found in this app")
			return
		}
		var ok bool
		for idx := range a.Commands {
			if ok = a.Commands[idx].Execute(flagset); ok {
				break
			}
		}

		if !ok {
			a.help()
		}

	}

}

func (a *App) Printf(format string, args ...interface{}) {
	Printf(format, args...)
}

var appTmpl = `NAME:
   {{.Name}} - {{.Description}}

USAGE:
   {{.Name}} command [command options] [arguments...]

VERSION:
   {{.Version}}

COMMANDS:
{{ range $index, $cmd := .Commands }}
   {{$cmd.Name }} [{{$cmd.Flags.ToString}}]        {{$cmd.Description}}
     {{ range $index, $subcmd := .Subcommands }}
     {{$subcmd.Name}}        {{$subcmd.Description}}
	 {{ end }}
{{ end }}`
