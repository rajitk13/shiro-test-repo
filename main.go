package main

import "fmt"

func main() {
	result := add(5, 3)
	fmt.Printf("Result: %d\n", result)

	str := reverse("hello")
	fmt.Printf("Reversed: %s\n", str)

	// New function with potential issues
	num := divide(10, 0)
	fmt.Printf("Division: %d\n", num)

	// Additional test case
	slice := []int{1, 2, 3}
	val := unsafeSliceAccess(slice, 5)
	fmt.Printf("Slice value: %d\n", val)
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

func divide(a, b int) int {
	// Potential division by zero issue
	return a / b
}

func unsafeSliceAccess(arr []int, index int) int {
	// No bounds checking - potential panic
	return arr[index]
}
