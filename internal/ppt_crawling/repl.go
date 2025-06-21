package ppt_crawling

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Client      Client
	DownloadURL map[string]string
	StartPages  int
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"search": {
			name:        "search",
			description: fmt.Sprintf("search praise download url (count: 10)"),
			callback:    commandSearch,
		},
		"download": {
			name:        "download",
			description: "download all praise",
			callback:    commandDownload,
		},
		"list": {
			name:        "list",
			description: "list all praise",
			callback:    commandList,
		},
		"exit": {
			name:        "exit",
			description: "exit repl",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func StartRepl(cfg *Config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("ppt-crawler > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if commandName == "s" {
			commandName = "search"
		}

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exist := getCommands()[commandName]
		if !exist {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	text_lower := strings.ToLower(text)
	words := strings.Fields(text_lower)
	return words
}
