package repl

import (
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	type testCase struct {
		input string
		expectedCommand string
		expectedArguments []string
	}

	printCase := func (command string, arguments []string) {
		fmt.Println("Case {")
		fmt.Printf("    %s\n", command)
		fmt.Printf("    [\n")
		for _, a := range arguments {
			fmt.Printf("        %s,\n", a)
		}
		fmt.Printf("    ]\n")
		fmt.Println("}")
	}

	printBoth := func(item testCase, actualCommand string, actualArguments []string) {
			fmt.Println("Expected:")
			printCase(item.expectedCommand, item.expectedArguments)
			fmt.Println()
			fmt.Println("Actual:")
			printCase(actualCommand, actualArguments)
	}

	runCase := func (item testCase) error {
		command, args := parseInput(item.input)
		if command != item.expectedCommand {
			return fmt.Errorf("Command parsing failed: %s != %s", command, item.expectedCommand)
		}
		if len(args) != len(item.expectedArguments) {
			printBoth(item, command, args)
			return fmt.Errorf("Argument parsing failed, incorrect number of items: %d != %d", len(args), len(item.expectedArguments))
		}
		
		for index, arg := range args {
			if arg != item.expectedArguments[index] {
				printBoth(item, command, args)
				return fmt.Errorf("Argument parsing failed, arguments differ: (%d) %s != %s", index, arg, item.expectedArguments[index])
			}
		}

		return nil
	}

	cases := []testCase{
		{
			input: "add person1 person2 somethingelse",
			expectedCommand: "add",
			expectedArguments: []string{
				"person1",
				"person2",
				"somethingelse",
			},
		},
		{
			input: "",
			expectedCommand: "",
			expectedArguments: nil,
		},
		{
			input: "command",
			expectedCommand: "command",
			expectedArguments: nil,
		},
		{
			input: "command \" with some \"quotes",
			expectedCommand: "command",
			expectedArguments: []string{
				"\"",
				"with",
				"some",
				"\"quotes",
			},
		},
	}

	for i, c := range cases {
		if err := runCase(c); err != nil {
			t.Errorf("Test case %d of %d failed: %v", i + 1, len(cases) + 1, err)
		}
	}
}
