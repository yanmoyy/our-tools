package sns

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
