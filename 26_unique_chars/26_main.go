package main

import "fmt"

func main() {
	fmt.Println(isUnique("BB"))
}

func isUnique(s string) bool {
	m := make(map[int32]struct{})
	for _, char := range s {
		fmt.Println(int(char))
		if char < 'a' {
			char = char + 'a' - 'A'
			fmt.Println(int(char))
		}
		if _, ok := m[char]; ok {
			return false
		}
		m[char] = struct{}{}
	}
	return true
}
