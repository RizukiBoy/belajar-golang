package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "jawa Barat", "indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
