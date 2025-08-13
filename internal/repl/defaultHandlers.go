package repl

import "fmt"


func getHelpHandler(commandList *replCommandList) *replCommand {
	return &replCommand{
		name: "help",
		description: "List all available commands.",
		callback: func(name string, args ...string) error {
			fmt.Println()
			for _, cmd := range commandList.commands {
				fmt.Printf("'%s' -> '%s'\n", cmd.name, cmd.description)
			}

			return nil
		},
	}
}

func getClearHandler(commandList *replCommandList) *replCommand {
	return &replCommand{
		name: "clear",
		description: "Clear screen.",
		callback: func(name string, args ...string) error {
			fmt.Println("\033[2J\033[H")
			commandList.printReplHint()

			return nil
		},
	}
}
