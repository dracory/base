package cmd

import (
	"runtime"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	// Platform-specific test setup
	var command string
	var args []string

	if runtime.GOOS == "windows" {
		command = "cmd"
		args = []string{"/c", "echo hello"}
	} else {
		command = "echo"
		args = []string{"hello"}
	}

	tests := []struct {
		name     string
		command  string
		args     []string
		wantOut  string
		wantErr  bool
	}{
		{
			name:    "successful command",
			command: command,
			args:    args,
			wantOut: "hello",
			wantErr: false,
		},
		{
			name:    "non-existent command",
			command: "nonexistentcommand",
			args:    []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Exec(tt.command, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(got, tt.wantOut) {
				t.Errorf("Exec() = %q, want to contain %q", got, tt.wantOut)
			}
		})
	}
}

func TestExecSeparated(t *testing.T) {
	t.Run("stdout and stderr separation", func(t *testing.T) {
		var command string
		var args []string

		if runtime.GOOS == "windows" {
			command = "cmd"
			args = []string{"/c", "echo stdout && echo stderr 1>&2"}
		} else {
			command = "sh"
			args = []string{"-c", "echo stdout && echo stderr 1>&2"}
		}

		stdout, stderr, err := ExecSeparated(command, args...)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !strings.Contains(stdout, "stdout") {
			t.Errorf("Expected stdout to contain 'stdout', got: %q", stdout)
		}
		if !strings.Contains(stderr, "stderr") {
			t.Errorf("Expected stderr to contain 'stderr', got: %q", stderr)
		}
	})

	t.Run("command error", func(t *testing.T) {
		_, _, err := ExecSeparated("nonexistentcommand", "arg")
		if err == nil {
			t.Errorf("Expected error for non-existent command, got nil")
		}
		// We don't care about the content of stdout/stderr here, just that an error was returned
	})
}

func TestExecLine(t *testing.T) {
	var cmdLine string
	if runtime.GOOS == "windows" {
		cmdLine = "cmd /c echo hello"
	} else {
		cmdLine = "echo hello"
	}

	tests := []struct {
		name    string
		cmd     string
		wantOut string
		wantErr bool
	}{
		{
			name:    "valid command",
			cmd:     cmdLine,
			wantOut: "hello",
			wantErr: false,
		},
		{
			name:    "empty command",
			cmd:     "",
			wantErr: true,
		},
		{
			name:    "invalid command format",
			cmd:     "   ",  // Multiple spaces will split to empty strings
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecLine(tt.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(got, tt.wantOut) {
				t.Errorf("ExecLine() = %q, want to contain %q", got, tt.wantOut)
			}
		})
	}
}

func TestExecLineSeparated(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		wantErr bool
	}{
		{
			name:    "command with arguments",
			cmd:     getCommandLine(),
			wantErr: false,
		},
		{
			name:    "empty command",
			cmd:     "",
			wantErr: true,
		},
		{
			name:    "invalid command format",
			cmd:     "   ",  // Multiple spaces will split to empty strings
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, err := ExecLineSeparated(tt.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecLineSeparated() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if !tt.wantErr {
				if !strings.Contains(stdout, "hello world") {
					t.Errorf("Expected stdout to contain 'hello world', got: %q", stdout)
				}
				if stderr != "" {
					t.Errorf("Expected empty stderr, got: %q", stderr)
				}
			}
		})
	}
}

// Helper function to get platform-specific command line
func getCommandLine() string {
	if runtime.GOOS == "windows" {
		return "cmd /c echo hello world"
	}
	return "echo hello world"
}
