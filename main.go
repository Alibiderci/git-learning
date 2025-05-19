package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!\n", name)
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, Git!")
	fmt.Println(greet("Alibi"))
	fmt.Println("2 + 3 =", add(2,3))
}
