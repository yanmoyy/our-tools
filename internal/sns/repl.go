package sns

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printGreeting() {
	fmt.Println("=========== SNS-Sender ==========")
	fmt.Println("Sending your SNS message on terminal.")
}

func StartRepl(cfg *config) {
	printGreeting()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("SNS > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			if command.helper != nil {
				command.helper()
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
	helper      func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the SNS-Sender",
			callback:    commandExit,
		},
		"mode": {
			name:        "mode",
			description: "Change (or show) target SNS Mode (args: [mode])",
			callback:    commandMode,
			helper:      printModeHelp,
		},
	}
}
