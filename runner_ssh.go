package hotomata

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"golang.org/x/crypto/ssh"
)

type SSHRunner struct {
}

// TODO(kiasaki) Respect "ignore_errors"
// TODO(kiasaki) Respect "skip"
func (r *SSHRunner) Run(machine Machine, command string) *TaskResponse {
	var response = &TaskResponse{
		Log:    &bytes.Buffer{},
		Action: TaskActionContinue,
		Status: TaskStatusSuccess,
	}

	// Local execution
	if machine.Hostname == "127.0.0.1" && machine.Port == 0 {
		cmd := exec.Command("/bin/sh", "-c", command)
		cmd.Stdout = &b
		cmd.Stderr = &b
		err := cmd.Run()
		if err != nil {
			response.Log.WriteString(err.Error() + "\n")
			response.Action = TaskActionAbort
			response.Status = TaskStatusError
		}
		return response
	}

	// Remote execution
	hostname := machine.Hostname + ":" + strconv.Itoa(machine.Port)
	client, err := ssh.Dial("tcp", hostname, machine.SSHConfig)
	if err != nil {
		fmt.Printf("Failed to dial: %s: %s\n", hostname, err.Error())
		os.Exit(1)
	}
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Unable to connect: %s: %s\n", hostname, err.Error())
		os.Exit(1)
	}
	defer session.Close()
	defer client.Close()

	modes := ssh.TerminalModes{
		ECHO:          0,
		TTY_OP_ISPEED: 14400,
		TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		fmt.Printf("Request for terminal failed: %s: %s\n", hostname, err.Error())
		os.Exit(1)
	}
	session.Stdout = &b
	session.Stderr = &b
	err = session.Run(command)
	if err != nil {
		response.Log.WriteString(err.Error() + "\n")
		response.Action = TaskActionAbort
		response.Status = TaskStatusError
	}

	return response
}
