package sns

import (
	"fmt"
	"sort"

	"github.com/yanmoyy/our-tools/internal/sns/cli"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the SNS-Sender!")
	fmt.Println("Usage: <command> [args]")
	if cfg.snsType != Default {
		fmt.Printf("Current Mode: %s\n", cfg.snsType.Upper())
		fmt.Println()
		fmt.Printf("%s:\n", cfg.snsType.Upper())
		sortAndPrintCommands(getSNSCommands(cfg))
		fmt.Println()
	}
	fmt.Println("Default:")
	sortAndPrintCommands(getDefaultCommands())
	fmt.Println()
	return nil
}

func sortAndPrintCommands(commands cli.Commands[config]) {
	keys := make([]string, 0, len(commands))
	for key := range commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cmd := commands[key]
		fmt.Printf("  - %s: %s\n", cmd.Name, cmd.Description)
	}
}
