package main

import "fmt"

func main() {
	result := add(5, 3)
	fmt.Printf("Result: %d\n", result)
	
	str := reverse("hello")
	fmt.Printf("Reversed: %s\n", str)
}

func add(a, b int) int {
	return a + b
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
