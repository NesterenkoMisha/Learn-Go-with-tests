package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

// Hello method returns a greeting string like: "Hello, $name"
// to the user in the required language.
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return HelloPrefix(language) + ", " + name
}

// HelloPrefix return the correct "Hello" phrase by the language inputted
func HelloPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Misha", "English"))
}
