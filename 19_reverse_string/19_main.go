package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(reverseString(str))
}

func reverseString(str string) string {
	strRunes := []rune(str)
	strLen := len(strRunes) - 1
	for i := 0; i < strLen/2; i++ {
		strRunes[i], strRunes[strLen-i] = strRunes[strLen-i], strRunes[i]
	}
	return string(strRunes)
}
