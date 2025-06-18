package sns

import "fmt"

type mode struct {
	snsType     snsType
	description string
}

func commandMode(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("need to specify mode")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}
	mode := args[0]

	switch snsType(mode) {
	case Kakao:
		setMode(cfg, Kakao)
	default:
		return fmt.Errorf("unknown mode: %s", mode)
	}
	return nil
}

func printModeHelp() {
	fmt.Println("Usage: sns mode <mode>")
	fmt.Println("Available modes:")
	for _, mode := range getAvaliableModes() {
		fmt.Printf("  - %s (%s)\n", mode.snsType, mode.description)
	}
}

func setMode(cfg *config, mode snsType) {
	cfg.snsType = mode
	fmt.Printf("Set mode to %s\n", mode)
}

func getAvaliableModes() []mode {
	modes := []mode{
		{
			snsType:     Kakao,
			description: "Kakao Talk",
		},
	}
	return modes
}
