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
	"fmt"
	"reflect"
	"strings"
)

type CommandFlags []*CommandFlag

type CommandFlag struct {
	Name    string
	Default interface{}
	Usage   string
	Value   interface{}
}

// Get returns a flag by it's name, if flag not found returns nil
func (c CommandFlags) Get(name string) *CommandFlag {
	for idx, v := range c {
		if v.Name == name {
			return c[idx]
		}
	}

	return nil
}

// String returns the flag's value as string by it's name, if not found returns empty string ""
func (c CommandFlags) String(name string) string {
	f := c.Get(name)
	if f == nil {
		return ""
	}
	return *f.Value.(*string) //*f.Value if string
}

// Bool returns the flag's value as bool by it's name, if not found returns false
func (c CommandFlags) Bool(name string) bool {
	f := c.Get(name)
	if f != nil {
		return *f.Value.(*bool)
	}
	return false
}

// Int returns the flag's value as int by it's name, if can't parse int then returns -1
func (c CommandFlags) Int(name string) int {
	f := c.Get(name)
	if f == nil {
		return -1
	}
	return *f.Value.(*int)
}

// IsValid returns true if flags are valid, otherwise false
func (c CommandFlags) IsValid() bool {
	if c.Validate() != nil {
		return false
	}
	return true
}

// Validate returns nil if this flags are valid, otherwise returns an error message
func (c CommandFlags) Validate() error {
	var notFilled []string
	for _, v := range c {
		// if no value given for required flag then it is not valid
		defaultVal := reflect.ValueOf(v.Default).String()
		val := reflect.ValueOf(v.Value).Elem().String()
		if defaultVal == "" && val == "" {
			notFilled = append(notFilled, v.Name)
		}
	}

	if len(notFilled) > 0 {
		if len(notFilled) == 1 {
			return fmt.Errorf("Command is not valid. Required flag [-%s] is missing", notFilled[0])
		} else {
			return fmt.Errorf("Command is not valid. Required flags [%s] are missing", strings.Join(notFilled, ","))
		}

	}
	return nil

}

func (c CommandFlags) ToString() (summary string) {
	for idx, v := range c {
		summary += "-" + v.Name
		if idx < len(c)-1 {
			summary += ", "
		}
	}

	return
}
