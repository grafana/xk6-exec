// Package exec provides the xk6 Modules implementation for running local commands using Javascript
package exec

import (
	"errors"
	"log"
	"os"
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
	Dir                  string
	ContinueOnError      bool
	IncludeStdoutOnError bool
	Env                  []string
	CombineOutput        bool
}

type myExitError struct {
	ProcessState *os.ProcessState
	Stderr       []byte
	Stdout       []byte
}

func (e *myExitError) Error() string {
	return e.ProcessState.String()
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
func (*EXEC) Command(name string, args []string, option CommandOptions) (string, error) {
	var out []byte
	var err error

	cmd := exec.Command(name, args...)

	if option.Dir != "" {
		cmd.Dir = option.Dir
	}

	if len(option.Env) > 0 {
		cmd.Env = append(cmd.Environ(), option.Env...)
	}

	if option.CombineOutput {
		out, err = cmd.CombinedOutput()
	} else {
		out, err = cmd.Output()
	}

	if err != nil && !option.ContinueOnError {
		log.Fatal(err.Error() + " on command: " + name + " " + strings.Join(args, " "))
	}

	if err != nil && option.IncludeStdoutOnError {
		var exitErr *exec.ExitError
		var myExitErr myExitError
		if errors.As(err, &exitErr) {
			myExitErr.Stderr = exitErr.Stderr
			myExitErr.Stdout = out
			myExitErr.ProcessState = exitErr.ProcessState
			return string(out), &myExitErr
		}
	}

	return string(out), err
}
