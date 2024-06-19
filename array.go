package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ahmad"
	names[1] = "Okta"
	names[2] = "Rizky"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		80,
		503,
		534,
		23,
		12,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])

	fmt.Println(len(values))

}
