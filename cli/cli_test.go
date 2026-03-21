package cli_test

import (
	"errors"
	"testing"

	"github.com/dracory/base/cli"
)

// MockRegistry represents a mock registry for testing
type MockRegistry struct {
	name string
}

// MockCommandHandler creates a mock command handler for testing
func MockCommandHandler(expectedResult error) cli.CommandHandler[MockRegistry] {
	return func(registry MockRegistry, args []string) error {
		return expectedResult
	}
}

func TestNewDispatcher(t *testing.T) {
	dispatcher := cli.NewDispatcher[MockRegistry]()
	if dispatcher == nil {
		t.Fatal("NewDispatcher returned nil")
	}
}

func TestRegisterCommand(t *testing.T) {
	dispatcher := cli.NewDispatcher[MockRegistry]()

	// Test successful registration
	err := dispatcher.RegisterCommand("test", "Test command", MockCommandHandler(nil))
	if err != nil {
		t.Fatal("Failed to register command:", err)
	}

	// Test duplicate registration
	err = dispatcher.RegisterCommand("test", "Another test command", MockCommandHandler(nil))
	if err == nil {
		t.Fatal("Expected error when registering duplicate command")
	}

	// Test empty name registration
	err = dispatcher.RegisterCommand("", "Empty name command", MockCommandHandler(nil))
	if err == nil {
		t.Fatal("Expected error when registering command with empty name")
	}
}

func TestExecuteCommand(t *testing.T) {
	dispatcher := cli.NewDispatcher[MockRegistry]()
	registry := MockRegistry{name: "test"}

	// Test successful command execution
	dispatcher.RegisterCommand("success", "Success command", MockCommandHandler(nil))
	err := dispatcher.ExecuteCommand(registry, []string{"success"})
	if err != nil {
		t.Fatal("Expected success, got error:", err)
	}

	// Test command execution with error
	dispatcher.RegisterCommand("error", "Error command", MockCommandHandler(errors.New("test error")))
	err = dispatcher.ExecuteCommand(registry, []string{"error"})
	if err == nil {
		t.Fatal("Expected error, got success")
	}
	if err.Error() != "test error" {
		t.Fatalf("Expected 'test error', got '%s'", err.Error())
	}

	// Test unrecognized command
	err = dispatcher.ExecuteCommand(registry, []string{"unknown"})
	if err == nil {
		t.Fatal("Expected error for unknown command")
	}
	if err.Error() != "unrecognized command: unknown" {
		t.Fatalf("Expected 'unrecognized command: unknown', got '%s'", err.Error())
	}

	// Test no command provided
	err = dispatcher.ExecuteCommand(registry, []string{})
	if err == nil {
		t.Fatal("Expected error for no command")
	}
	if err.Error() != "no command provided" {
		t.Fatalf("Expected 'no command provided', got '%s'", err.Error())
	}
}

func TestListCommands(t *testing.T) {
	dispatcher := cli.NewDispatcher[MockRegistry]()

	// Test empty command list
	commands := dispatcher.ListCommands()
	if len(commands) != 0 {
		t.Fatalf("Expected empty command list, got %d commands", len(commands))
	}

	// Test non-empty command list
	dispatcher.RegisterCommand("cmd1", "Command 1", MockCommandHandler(nil))
	dispatcher.RegisterCommand("cmd2", "Command 2", MockCommandHandler(nil))
	commands = dispatcher.ListCommands()
	if len(commands) != 2 {
		t.Fatalf("Expected 2 commands, got %d", len(commands))
	}

	// Verify command details
	foundCmd1, foundCmd2 := false, false
	for _, cmd := range commands {
		if cmd.Name == "cmd1" && cmd.Description == "Command 1" {
			foundCmd1 = true
		}
		if cmd.Name == "cmd2" && cmd.Description == "Command 2" {
			foundCmd2 = true
		}
	}
	if !foundCmd1 {
		t.Fatal("Command 'cmd1' not found in list")
	}
	if !foundCmd2 {
		t.Fatal("Command 'cmd2' not found in list")
	}
}

func TestGetCommand(t *testing.T) {
	dispatcher := cli.NewDispatcher[MockRegistry]()

	// Test getting non-existent command
	cmd := dispatcher.GetCommand("nonexistent")
	if cmd != nil {
		t.Fatal("Expected nil for non-existent command")
	}

	// Test getting existing command
	dispatcher.RegisterCommand("test", "Test command", MockCommandHandler(nil))
	cmd = dispatcher.GetCommand("test")
	if cmd == nil {
		t.Fatal("Expected command for 'test'")
	}
	if cmd.Name != "test" || cmd.Description != "Test command" {
		t.Fatalf("Expected name 'test' and description 'Test command', got '%s' and '%s'", cmd.Name, cmd.Description)
	}
}

func TestHasCommand(t *testing.T) {
	dispatcher := cli.NewDispatcher[MockRegistry]()

	// Test non-existent command
	if dispatcher.HasCommand("nonexistent") {
		t.Fatal("Expected false for non-existent command")
	}

	// Test existing command
	dispatcher.RegisterCommand("test", "Test command", MockCommandHandler(nil))
	if !dispatcher.HasCommand("test") {
		t.Fatal("Expected true for existing command")
	}
}

func TestExecuteCommandWithArgs(t *testing.T) {
	dispatcher := cli.NewDispatcher[MockRegistry]()
	registry := MockRegistry{name: "test"}

	// Test command with arguments
	var receivedArgs []string
	handler := func(registry MockRegistry, args []string) error {
		receivedArgs = args
		return nil
	}

	dispatcher.RegisterCommand("args", "Args command", handler)
	err := dispatcher.ExecuteCommand(registry, []string{"args", "arg1", "arg2", "arg3"})
	if err != nil {
		t.Fatal("Expected success, got error:", err)
	}

	if len(receivedArgs) != 3 {
		t.Fatalf("Expected 3 arguments, got %d", len(receivedArgs))
	}
	if receivedArgs[0] != "arg1" || receivedArgs[1] != "arg2" || receivedArgs[2] != "arg3" {
		t.Fatalf("Expected [arg1, arg2, arg3], got %v", receivedArgs)
	}
}
