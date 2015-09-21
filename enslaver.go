package enslaver

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

const (
	exitNormal            = 0
	exitUnknownErr        = 125
	exitCommandNotInvoked = 126
	exitCommandNotFound   = 127
)

func Command(name string, args ...string) *Slave {
	return &Slave{
		cmdName: name,
		cmdArgs: args,
	}
}

type Slave struct {
	cmdName string
	cmdArgs []string
}

func (sl *Slave) Run() int {
	for {
		sl.labor()
	}
}

func (sl *Slave) newCommand() *exec.Cmd {
	return exec.Command(sl.cmdName, sl.cmdArgs...)
}

func (sl *Slave) labor() int {
	cmd := sl.newCommand()
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitUnknownErr
	}
	defer stdoutPipe.Close()

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitUnknownErr
	}
	defer stderrPipe.Close()

	if err := cmd.Start(); err != nil {
		switch {
		case os.IsNotExist(err):
			return exitCommandNotFound
		case os.IsPermission(err):
			return exitCommandNotInvoked
		default:
			return exitUnknownErr
		}
	}
	go io.Copy(os.Stdout, stdoutPipe)
	go io.Copy(os.Stderr, stderrPipe)

	return <- getExitChan(cmd)
}

func getExitChan(cmd *exec.Cmd) chan int {
	ch := make(chan int)
	go func() {
		err := cmd.Wait()
		ch <- resolveExitCode(err)
	}()
	return ch
}

func resolveExitCode(err error) int {
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus()
			}
		}
		// The exit codes in some platforms aren't integer. e.g. plan9.
		return -1
	}
	return exitNormal
}
