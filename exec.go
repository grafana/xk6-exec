package exec

import (
	"log"
	"strings"
	"os/exec"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/exec", new(EXEC))
}

// EXEC is the k6 say extension.
type EXEC struct{}

// Command is a wrapper for Go exec.Command
func (*EXEC) Command(name string, args []string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Fatal(err.Error() + " on command: " + name + " " + strings.Join(args, " "))
	}
	return string(out)
}
