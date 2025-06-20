package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	crawling_api "github.com/yanmoyy/our-tools/internal/crawling-api"
)

type config struct {
	client      crawling_api.Client
	downloadURL map[string]string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"title": {
			name:        "title",
			description: fmt.Sprintf("search praise download url in %s", TistoryTitle),
			callback:    commandTitle,
		},
		"num": {
			name:        "num",
			description: fmt.Sprintf("search praise download url in %s", TistoryNum),
			callback:    commandNum,
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

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("ppt-crawler > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

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
