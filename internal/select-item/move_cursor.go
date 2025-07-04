package select_item

import "fmt"

func moveCursorUp(n int) {
	fmt.Printf("\x1b[%dA", n)
}

func moveCursorFront() {
	fmt.Printf("\x1b[%d;0H", len(options)+1)
}
