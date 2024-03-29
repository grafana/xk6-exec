// Package exec provides the xk6 Modules implementation for running local commands using Javascript
package exec

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/exec", new(RootModule))
}

// RootModule is the global module object type. It is instantiated once per test
// run and will be used to create `k6/x/exec` module instances for each VU.
type RootModule struct{}

// EXEC represents an instance of the EXEC module for every VU.
type EXEC struct {
	vu modules.VU
}

// CommandOptions contains the options that can be passed to command.
type CommandOptions struct {
	Dir        string
	Env        []string
	FatalError *bool `js:"fatalError"`
}

type CommandReturn struct {
	StatusCode int `js:"statusCode"`
	Stdout     string
	Stderr     string
}

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &EXEC{}
)

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &EXEC{vu: vu}
}

// Exports implements the modules.Instance interface and returns the exports
// of the JS module.
func (exec *EXEC) Exports() modules.Exports {
	return modules.Exports{Default: exec}
}

// Command is a wrapper for Go exec.Command
func (*EXEC) Command2(name string, args []string, option CommandOptions) CommandReturn {
	cmd := exec.Command(name, args...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	if option.Dir != "" {
		cmd.Dir = option.Dir
	}

	if len(option.Env) > 0 {
		cmd.Env = append(cmd.Environ(), option.Env...)
	}
	err := cmd.Run()
	statusCode := cmd.ProcessState.ExitCode()
	if err != nil {
		// Keep default behaviour backwards compatible
		if option.FatalError == nil || *option.FatalError {
			log.Fatal(err.Error() + " on command: " + name + " " + strings.Join(args, " "))
		} else {
			log.Print(err.Error() + " on command: " + name + " " + strings.Join(args, " "))
		}
	}

	return CommandReturn{
		StatusCode: statusCode,
		Stdout:     outb.String(),
		Stderr:     errb.String(),
	}
}

func (e *EXEC) Command(name string, args []string, option CommandOptions) string {
	retval := e.Command2(name, args, option)
	if retval.StatusCode == 0 {
		return retval.Stdout
	} else {
		return retval.Stderr
	}
}
