package slice

import "fmt"

func lenthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {
	fmt.Println(lenthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lenthOfNonRepeatingSubStr("bbb"))
	fmt.Println(lenthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lenthOfNonRepeatingSubStr(""))
	fmt.Println(lenthOfNonRepeatingSubStr("b"))
	fmt.Println(lenthOfNonRepeatingSubStr("abcdefg"))
	fmt.Println(lenthOfNonRepeatingSubStr("一二三三二一"))
	fmt.Println(lenthOfNonRepeatingSubStr("这里是上海"))
}
