package slice

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱你!" //utf-8编码
	fmt.Println(s)
	fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s))
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
