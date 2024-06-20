package main

import "fmt"

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("Eko", spamFilter)

	filter := spamFilter

	sayHelloFilter("anjing", filter)
}
