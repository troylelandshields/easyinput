package easyinput

import (
	"bufio"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func TakeInput() string {
	scanner.Scan()
	return scanner.Text()
}
