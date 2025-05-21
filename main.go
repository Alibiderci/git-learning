package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func farewell(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}

func reverse(word string) string {
	if len(word) == 1 {
		return word
	}
	return string(word[len(word)-1]) + reverse(word[:len(word)-1])
}

func main() {
	fmt.Println("Hello, Git!")
	fmt.Println(greet("Alibi"))
	rev := reverse("alibI")
	fmt.Println(farewell(rev))
}
