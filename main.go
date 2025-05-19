package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!\n", name)
}

func add(a, b int) int {
	return a + b
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	return fibonacci(num - 2) + fibonacci(num - 1)
}

func main() {
	fmt.Println("Hello, Git!")
	fmt.Println(greet("Alibi"))
	fmt.Println("2 + 3 =", add(2,3))
	fmt.Println("The 7th number in fibonacci sequence is", fibonacci(7))
}
