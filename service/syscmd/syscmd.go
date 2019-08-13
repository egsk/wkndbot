package syscmd

import (
	"os/exec"
	"wkndbot/cnf"
)


func Exec(command cnf.Command) string {

	cmd := exec.Command(command.Provider, command.Command...)
	stdout, err := cmd.Output()
	if err != nil {
		stdout = []byte(err.Error())
	}

	return string(stdout)
}
