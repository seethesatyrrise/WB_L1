package main

import "fmt"

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(doSet(input))
}

func doSet(in []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, val := range in {
		set[val] = struct{}{}
	}
	return set
}
