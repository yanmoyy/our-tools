package sns

import "fmt"

type mode struct {
	snsType     snsType
	description string
}

func commandMode(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Current mode:", cfg.snsType)
		printModeHelp()
		return nil
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
	fmt.Println("Usage: mode [mode]")
	fmt.Println("Available modes:")
	for _, mode := range getAvaliableModes() {
		fmt.Printf("  - %s (%s)\n", mode.snsType, mode.description)
	}
}

func setMode(cfg *config, mode snsType) {
	err := cfg.setMode(mode)
	if err != nil {
		fmt.Printf("Failed to set mode: %s\n", err)
		return
	}
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
