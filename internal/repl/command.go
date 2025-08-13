package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ReplCommandHandlerFunc func(name string, args ...string) error

type Repl interface {
	AddCommand(name, description string, callback ReplCommandHandlerFunc) error
	Start() error
}

type ReplConfig struct {
	Prompt string
}

type replCommand struct {
	name string
	description string
	callback ReplCommandHandlerFunc
}

type replCommandList struct {
	config ReplConfig
	commands map[string]*replCommand
}

func NewReplCommandList(config ReplConfig) Repl {
	if config.Prompt == "" {
		config.Prompt = "  >"
	}
	newList := &replCommandList{
		config: config,
		commands: make(map[string]*replCommand),
	}

	newList.addCommand(getHelpHandler(newList))
	newList.addCommand(getClearHandler(newList))
	return newList
}

func newReplCommand(name, description string, handler ReplCommandHandlerFunc) *replCommand {
	return &replCommand{
		name: name,
		description: description,
		callback: handler,
	}
}

func (r *replCommandList) printReplHint() {
	fmt.Println()
	fmt.Println("Enter 'exit' to quit")
	fmt.Println("Use 'help' to list available commands.")
}

func (r *replCommandList) AddCommand(name, description string, callback ReplCommandHandlerFunc) error {
	return r.addCommand(newReplCommand(name, description, callback))
}

func (r *replCommandList) addCommand(cmd *replCommand) error {
	if _, exists := r.commands[cmd.name]; exists {
		return fmt.Errorf("A command, named '%s', already exists", cmd.name)
	}

	r.commands[cmd.name] = cmd
	return nil
}

func (r *replCommandList) Start() error {
	reader := bufio.NewScanner(os.Stdin)
	r.printReplHint()

	for {
		fmt.Println()
		fmt.Printf("%s ", r.config.Prompt)

		if !reader.Scan() {
			if err := reader.Err(); err != nil {
				return fmt.Errorf("Error reading from StdIn: %s", err)
			}
			continue
		}
		userInput := reader.Text()

		// continue if user entered nothing.
		if len(userInput) == 0 {
			continue
		}
		// check for exit command.
		if strings.ToLower(userInput) == "exit" {
			break
		}

		commandName, args := parseInput(userInput)
		
		command, exists := r.commands[commandName]
		if !exists {
			fmt.Printf("'%s' is not a valid command\n", commandName)
			continue
		}

		if err := command.callback(commandName, args...); err != nil {
			fmt.Printf("Error running command: %s\n", err)
			continue
		}
	}

	return nil
}
