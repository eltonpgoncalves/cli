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
	"strings"
)

type (
	CommandArgs struct {
		Args map[string]interface{}
	}

	CommandAction func(args CommandArgs)

	Commands []*Command

	Command struct {
		Name        string
		SmallName   string
		Description string
		Action      CommandAction
		Subcommands Commands
	}
)

func (c *CommandArgs) String(name string) string {
	return c.Args[name].(string)
}

func (c *CommandArgs) Bool(name string) bool {
	return c.Args[name].(bool)
}

func (c *CommandArgs) Int(name string) int {
	return c.Args[name].(int)
}

func NewCommand(name string, description string) *Command {
	minusLen := strings.Count(name, "-")
	// help -> --help & -help -> --help
	if minusLen < 2 {
		name = "--" + name
	}
	// --help  -> -h
	smallName := name[1:3]
	defaultAction := func(a CommandArgs) {
		Printf(ErrNoAction, name)
	}
	return &Command{name, smallName, description, defaultAction, nil}
}

// GetName  removes all - from the Name, --help -> help
func (c *Command) GetName() string {
	return strings.Replace(c.Name, "-", "", -1)
}

// Add adds a child command (subcommand)
func (c *Command) Add(subCommand *Command) *Command {
	if c.Subcommands == nil {
		c.Subcommands = Commands{}
	}

	c.Subcommands = append(c.Subcommands, subCommand)
	return c
}
