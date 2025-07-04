package select_item

import (
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	} else {
		cmd = exec.Command("clear") // For Linux
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
