package main

import "fmt"

const englishHeloPrefix = "Hello, "

func Hello() string {
	return englishHeloPrefix + "world"
}

func HelloName(name string) string {
	return englishHeloPrefix + name
}

func main() {
	fmt.Println(Hello())
}