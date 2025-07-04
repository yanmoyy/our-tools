package select_item

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func Repl() {

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
		if n == 3 && buf[0] == 27 && buf[1] == 91 {
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
		} else if n == 1 && buf[0] == 13 {
			break
		}
		moveCursorUp(len(options))
		renderOptions(selected)
		moveCursorFront()

	}

	if err := clearScreen(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear screen: %v\n", err)
	}
	fmt.Printf("\x1b[HYou selected: \x1b[32m%s\x1b[0m\n", options[selected])
	fmt.Println("\x1b[EGoodbye!")
	fmt.Print("\x1b[E")
}
