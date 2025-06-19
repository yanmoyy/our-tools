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
		err := setMode(cfg, Kakao)
		if err != nil {
			return err
		}
		return cfg.kakaoConfig.StartMode()
	default:
		return fmt.Errorf("unknown mode argument: %s", mode)
	}
}

func printModeHelp() {
	fmt.Println("Usage: mode [mode]")
	fmt.Println("Available modes:")
	for _, mode := range getAvaliableModes() {
		fmt.Printf("  - %s (%s)\n", mode.snsType, mode.description)
	}
}

// call setMode from config, and print line
func setMode(cfg *config, mode snsType) error {
	err := cfg.setMode(mode)
	if err != nil {
		return fmt.Errorf("failed to set mode: %s", err)
	}
	fmt.Printf("Set mode to %s\n", mode)
	return nil
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
