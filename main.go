package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!")
}

func farewell(name string) string {
	return fmt.Sprintf("Goodbye, %s!")
}

func main() {
	fmt.Println("Hello, Git!")
	fmt.Println(greet("Alibi"))
	fmt.Println(farewell("Alibi"))
}
