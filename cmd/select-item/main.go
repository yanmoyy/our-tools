package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"golang.org/x/term"
)

var options = []string{
	"Send Email",
	"Check Inbox",
	"Exit",
}

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

func moveCursorUp(n int) {
	fmt.Printf("\x1b[%dA", n)
}

func renderOptions(selected int) {
	fmt.Print("\x1b[2J\x1b[H") // Move to top-left and clear screen
	for i, opt := range options {
		fmt.Print("\x1b[2K\r") // Clear line and move cursor to start
		if i == selected {
			fmt.Printf("\x1b[32m> %s\x1b[0m\n", opt) // ">" with green color
		} else {
			fmt.Printf("%s\n", opt) // No color
		}
	}
	fmt.Printf("\x1b[%d;0H", len(options)+1) // Move cursor to selected option
}

func main() {
	// Set terminal to raw mode
	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set raw mode: %v\n", err)
		os.Exit(1)
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	selected := 0
	if err := clearScreen(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear screen: %v\n", err)
	}
	renderOptions(selected)

	buf := make([]byte, 3)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read input: %v\n", err)
			break
		}
		if n == 3 && buf[0] == 27 && buf[1] == 91 { // Arrow key escape sequence
			switch buf[2] {
			case 65: // Up arrow
				if selected > 0 {
					selected--
				}
			case 66: // Down arrow
				if selected < len(options)-1 {
					selected++
				}
			}
		} else if n == 1 && buf[0] == 13 { // Enter
			break
		}
		moveCursorUp(len(options))
		renderOptions(selected)
		fmt.Printf("\x1b[%d;0H", len(options)+1) // Move cursor to selected option
	}

	if err := clearScreen(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear screen: %v\n", err)
	}
	fmt.Printf("\x1b[HYou selected: \x1b[32m%s\x1b[0m\n", options[selected])
	fmt.Println("\x1b[EGoodbye!")
	fmt.Print("\x1b[E")
}
