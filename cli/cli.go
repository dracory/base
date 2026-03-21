package cli

import (
	"errors"
	"fmt"
)

// CommandHandler defines the function signature for command handlers.
// It accepts a generic registry interface and command arguments.
type CommandHandler[T any] func(registry T, args []string) error

// Command represents a CLI command with its handler and description.
type Command[T any] struct {
	Name        string
	Description string
	Handler     CommandHandler[T]
}

// Dispatcher manages CLI command registration and execution.
type Dispatcher[T any] struct {
	commands map[string]Command[T]
}

// NewDispatcher creates a new CLI command dispatcher.
func NewDispatcher[T any]() *Dispatcher[T] {
	return &Dispatcher[T]{
		commands: make(map[string]Command[T]),
	}
}

// RegisterCommand registers a new command with the dispatcher.
//
// Parameters:
// - name: The command name
// - description: Description of what the command does
// - handler: Function to handle the command
//
// Returns:
// - error: If command name is empty or already registered
func (d *Dispatcher[T]) RegisterCommand(name, description string, handler CommandHandler[T]) error {
	if name == "" {
		return errors.New("command name cannot be empty")
	}

	if _, exists := d.commands[name]; exists {
		return fmt.Errorf("command '%s' is already registered", name)
	}

	d.commands[name] = Command[T]{
		Name:        name,
		Description: description,
		Handler:     handler,
	}

	return nil
}

// ExecuteCommand executes a CLI command based on the provided arguments.
//
// Business logic:
// 1. Logs the command being executed.
// 2. Validates that at least one argument (the command) is provided.
// 3. Looks up the command in the registry.
// 4. If a handler is found, executes it with the remaining arguments.
// 5. If no handler is found, returns an "unrecognized command" error.
//
// Parameters:
// - registry: The registry instance to be passed to command handlers
// - args: The command line arguments (excluding the program name)
//
// Returns:
// - error: An error if the command execution fails or is invalid, otherwise nil
func (d *Dispatcher[T]) ExecuteCommand(registry T, args []string) error {
	fmt.Println("Executing command:", args)

	if len(args) == 0 {
		fmt.Println("No command provided.")
		return errors.New("no command provided")
	}

	command := args[0]
	remainingArgs := args[1:] // Arguments after the main command

	// Look up the command
	cmd, found := d.commands[command]
	if !found {
		err := fmt.Errorf("unrecognized command: %s", command)
		fmt.Println(err.Error())
		return err
	}

	// Execute the found handler with registry
	return cmd.Handler(registry, remainingArgs)
}

// ListCommands returns a list of all registered commands with their descriptions.
func (d *Dispatcher[T]) ListCommands() []Command[T] {
	var commands []Command[T]
	for _, cmd := range d.commands {
		commands = append(commands, cmd)
	}
	return commands
}

// GetCommand returns a command by name, or nil if not found.
func (d *Dispatcher[T]) GetCommand(name string) *Command[T] {
	if cmd, exists := d.commands[name]; exists {
		return &cmd
	}
	return nil
}

// HasCommand checks if a command is registered.
func (d *Dispatcher[T]) HasCommand(name string) bool {
	_, exists := d.commands[name]
	return exists
}

// PrintUsage prints usage information for all registered commands.
func (d *Dispatcher[T]) PrintUsage() {
	fmt.Println("Available commands:")
	for _, cmd := range d.ListCommands() {
		fmt.Printf("  %-15s - %s\n", cmd.Name, cmd.Description)
	}
}
