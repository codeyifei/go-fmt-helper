package go_fmt_helper

import "fmt"

func CursorUp() {
	fmt.Print("\033[1A")
}

func CursorUpLines(line int) {
	fmt.Printf("\033[%dA", line)
}

func CursorDown() {
	fmt.Println()
}

func CursorDownLines(line int) {
	if line >= 0 {
		for i := 0; i < line; i++ {
			fmt.Println()
		}
	} else {
		CursorUpLines(-line)
	}
}

func EmptyLine() {
	fmt.Print("\r\033[K")
}
