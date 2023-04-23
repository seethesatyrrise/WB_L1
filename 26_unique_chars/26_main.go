package main

import "fmt"

func main() {
	fmt.Println(isUnique("BB"))
}

func isUnique(s string) bool {
	m := make(map[int32]struct{})
	// наполнение множества символами строки
	for _, char := range s {
		fmt.Println(int(char))
		if char < 'a' {
			char = char + 'a' - 'A'
		}
		// если символ уже находится в множестве, возвращаем false
		if _, ok := m[char]; ok {
			return false
		}
		m[char] = struct{}{}
	}
	// если не произошло повторов символов
	return true
}
