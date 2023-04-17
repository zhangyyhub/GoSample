package print

import "fmt"

// 方法：彩色打印
func PrintColuor(code string, str string, line string) {
	fmt.Printf("\033[1;%sm%s\033[0m%s", code, str, line)
}
