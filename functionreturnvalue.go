package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("ahmad")
	fmt.Println(result)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Riski"))

}
