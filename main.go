package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!\n", name)
}

func main() {
	fmt.Println("Hello, Git!")
	fmt.Println(greet("Alibi"))
}
