package repl

import "strings"

func parseInput(input string) (command string, arguments []string) {
	lowercase := strings.ToLower(input)

	parts := strings.Fields(lowercase)
	if len(parts) == 0 {
		return "", nil
	}

	return parts[0], parts[1:]
}
