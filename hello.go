package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello method to return
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Misha", "English"))
}
