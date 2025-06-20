package sns

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yanmoyy/our-tools/internal/sns/cli"
	"github.com/yanmoyy/our-tools/internal/sns/kakao"
)

func printGreeting() {
	fmt.Println("=========== SNS-Sender ==========")
	fmt.Println("Sending your SNS message on terminal.")
}

func StartRepl(cfg *config) {
	printGreeting()
	reader := bufio.NewScanner(os.Stdin)
	for {
		if cfg.snsType == Default {
			fmt.Print("SNS > ")
		} else {
			fmt.Printf("%s > ", cfg.snsType.Upper())
		}
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
		command, exists := getCommands(cfg)[commandName]
		if !exists {
			fmt.Println("Invalid command")
			continue
		}
		err := command.Callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			if strings.Contains(err.Error(), "argument") {
				if command.Helper == nil {
					fmt.Println("No help message")
					continue
				}
				command.Helper()
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands(cfg *config) cli.Commands[config] {
	commands := getDefaultCommands()
	for _, cmd := range getSNSCommands(cfg) {
		commands[cmd.Name] = cmd
	}
	return commands
}

func getDefaultCommands() cli.Commands[config] {
	return cli.Commands[config]{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the SNS-Sender",
			Callback:    commandExit,
		},
		"mode": {
			Name:        "mode",
			Description: "Change target SNS Mode (or default)",
			Callback:    commandMode,
			Helper:      printModeHelp,
		},
	}
}

func getSNSCommands(cfg *config) cli.Commands[config] {
	commands := cli.Commands[config]{}
	switch cfg.snsType {
	case Kakao:
		kakaoCmds := kakao.GetCommands()
		for _, cmd := range kakaoCmds {
			commands[cmd.Name] = convertCommandToDefault(cfg.kakao, cmd)
		}
	}
	return commands
}

func convertCommandToDefault[T any](cfg *T, cmd cli.Command[T]) cli.Command[config] {
	return cli.Command[config]{
		Name:        cmd.Name,
		Description: cmd.Description,
		Callback: func(commonCfg *config, args ...string) error {
			return cmd.Callback(cfg, args...)
		},
		Helper: cmd.Helper,
	}
}
