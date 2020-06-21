package go_fmt_helper

import "fmt"

// 使控制台光标上移一行（自动重置光标到首列）
func CursorUp() {
	CursorUpLines(1)
}

// 使控制台光标上移多行，如果line为负数，则进行下移其绝对值行（自动重置光标到首列）
func CursorUpLines(line int) {
	if line > 0 {
		fmt.Printf("\033[%dA\r", line)
	} else if line < 0 {
		CursorDownLines(-line)
	}
}

// 使控制台光标下移一行（自动重置光标到首列）
func CursorDown() {
	CursorDownLines(1)
}

// 使控制台光标下移多行，如果line为负数，则进行上移其绝对值行（自动重置光标到首列）
func CursorDownLines(line int) {
	if line > 0 {
		for i := 0; i < line; i++ {
			fmt.Println()
		}
	} else if line < 0 {
		CursorUpLines(-line)
	}
}

// 重置光标到首列
func ResetFirstColumn() {
	fmt.Print("\r")
}

// 清空该行全部内容
func ClearLine() {
	fmt.Print("\r\033[K")
}

// 清屏
func ClearScreen() {
	fmt.Print("\033[2J")
}

// 跳转到指定行列
func DumpCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

// 光标左移指定行数
func CursorLeftLines(line int) {
	if line > 0 {
		fmt.Printf("\033[%dD", line)
	} else if line < 0 {
		CursorRightLines(-line)
	}
}

// 光标右移指定位置
func CursorRightLines(line int) {
	if line > 0 {
		fmt.Printf("\033[%dC", line)
	} else if line < 0 {
		CursorLeftLines(-line)
	}
}

// 保存光标位置
func SaveCursorLocation() {
	fmt.Print("\033[s")
}

// 恢复光标位置
func RestoreCursorLocation() {
	fmt.Print("\033[u")
}
