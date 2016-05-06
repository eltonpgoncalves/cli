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
	"strings"
)

type (
	CommandAction func(CommandFlags) error

	Commands []*Command

	Command struct {
		Name        string
		Description string
		// Flags are not the arguments was given by the user, but the flags that developer sets to this command
		Flags       CommandFlags
		action      CommandAction
		Subcommands Commands
		flagset     *flag.FlagSet
	}
)

func DefaultAction(cmdName string) CommandAction {
	return func(a CommandFlags) error { Printf(ErrNoAction, cmdName); return nil }
}

func NewCommand(name string, description string) *Command {
	name = strings.Replace(name, "-", "", -1) //removes all - if present, --help -> help
	fset := flag.NewFlagSet(name, flag.PanicOnError)
	return &Command{Name: name, Description: description, Flags: CommandFlags{}, action: DefaultAction(name), flagset: fset}
}

// Subcommand adds a child command (subcommand)
func (c *Command) Subcommand(subCommand *Command) *Command {
	if c.Subcommands == nil {
		c.Subcommands = Commands{}
	}

	c.Subcommands = append(c.Subcommands, subCommand)
	return c
}

func (c *Command) requestFlagValue(name string, defaultValue interface{}, usage string) interface{} {
	switch defaultValue.(type) {
	case int:
		{
			valPointer := c.flagset.Int(name, defaultValue.(int), usage)

			// it's not h (-h) for example but it's host, then assign it's alias also
			if len(name) > 1 {
				alias := name[0:1]
				c.flagset.IntVar(valPointer, alias, defaultValue.(int), usage)
			}
			return valPointer
		}
	case bool:
		{
			valPointer := c.flagset.Bool(name, defaultValue.(bool), usage)

			// it's not h (-h) for example but it's host, then assign it's alias also
			if len(name) > 1 {
				alias := name[0:1]
				c.flagset.BoolVar(valPointer, alias, defaultValue.(bool), usage)
			}
			return valPointer
		}
	default:
		valPointer := c.flagset.String(name, defaultValue.(string), usage)

		// it's not h (-h) for example but it's host, then assign it's alias also
		if len(name) > 1 {
			alias := name[0:1]
			c.flagset.StringVar(valPointer, alias, defaultValue.(string), usage)
		}

		return valPointer

	}
}

func (c *Command) Flag(name string, defaultValue interface{}, usage string) *Command {
	if c.Flags == nil {
		c.Flags = CommandFlags{}
	}
	valPointer := c.requestFlagValue(name, defaultValue, usage)

	newFlag := &CommandFlag{name, defaultValue, usage, valPointer}
	c.Flags = append(c.Flags, newFlag)
	return c
}

func (c *Command) Action(action CommandAction) *Command {
	c.action = action
	return c
}

// Execute returns true if this command has been executed
func (c *Command) Execute(parentflagset *flag.FlagSet) bool {
	var index = -1
	// check if this command has been called from app's arguments
	for idx, a := range parentflagset.Args() {
		if c.Name == a {
			index = idx + 1
		}
	}

	// this command hasn't been called from the user
	if index == -1 {
		return false
	}

	if !c.flagset.Parsed() {

		if err := c.flagset.Parse(parentflagset.Args()[index:]); err != nil {
			panic("Panic on command.Execute: " + err.Error())
		}
	}

	if err := c.Flags.Validate(); err == nil {
		c.action(c.Flags)

		for idx := range c.Subcommands {
			if c.Subcommands[idx].Execute(c.flagset) {
				break
			}

		}
	} else {
		Printf(err.Error())
	}
	return true

}
