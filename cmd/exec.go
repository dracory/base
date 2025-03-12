package cmd

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// Exec executes a command with arguments and captures
// its combined stdout and stderr.
//
// Parameters:
//   - name: the command to execute
//   - args: the arguments to pass to the command
//
// Returns:
//   - output: the combined stdout and stderr of the command
//   - error: any error that occurred during execution
func Exec(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// ExecSeparated executes a command with arguments
// and captures its stdout and stderr separately.
//
// Parameters:
//   - name: the command to execute
//   - args: the arguments to pass to the command
//
// Returns:
//   - stdout: the standard output of the command
//   - stderr: the standard error of the command
//   - error: any error that occurred during execution
func ExecSeparated(name string, args ...string) (string, string, error) {
	cmd := exec.Command(name, args...)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	return stdoutBuf.String(), stderrBuf.String(), err
}

// ExecLine executes a full system command line string
// and returns its combined output and any error that occurs.
//
// Parameters:
//   - cmd: the command line string to execute
//
// Returns:
//   - output: the combined stdout and stderr of the command
//   - error: any error that occurred during execution
func ExecLine(cmd string) (string, error) {
	if cmd == "" {
		return "", errors.New("a blank command")
	}

	cs := strings.Split(cmd, " ")
	if len(cs) == 0 {
		return "", errors.New("invalid command format")
	}

	return Exec(cs[0], cs[1:]...)
}

// ExecLineSeparated executes a full system command line
// string and returns its stdout, stderr, and any error
// that occurs.
//
// Parameters:
//   - cmd: the command line string to execute
//
// Returns:
//   - stdout: the standard output of the command
//   - stderr: the standard error of the command
//   - error: any error that occurred during execution
func ExecLineSeparated(cmd string) (string, string, error) {
	if cmd == "" {
		return "", "", errors.New("a blank command")
	}

	cs := strings.Split(cmd, " ")
	if len(cs) == 0 {
		return "", "", errors.New("invalid command format")
	}

	return ExecSeparated(cs[0], cs[1:]...)
}
