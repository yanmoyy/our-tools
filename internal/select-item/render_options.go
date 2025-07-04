package select_item

import "fmt"

func renderOptions(selected int) {
	fmt.Print("\x1b[2J\x1b[H") // top-left and clear screen
	for i, opt := range options {
		fmt.Print("\x1b[2K\r") // clear line and move cursor to top-left
		if i == selected {
			fmt.Printf("\x1b[32m> %s\x1b[0m\n", opt) // ">" with green color
		} else {
			fmt.Printf("%s\n", opt) // non selected option
		}
	}
	fmt.Printf("\x1b[%d;0H", len(options)+1) // move cursor to selected option
}
