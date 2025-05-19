package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!\n", name)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}

func main() {
	fmt.Println("Hello, Git!")
	fmt.Println(greet("Alibi"))
	fmt.Println("Factorial of 5 is", factorial(5))
}
