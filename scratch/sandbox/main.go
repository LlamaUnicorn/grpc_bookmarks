package main

import "fmt"

func main() {
	fmt.Println("Hey")

	var p *int
	fmt.Println(p)
	// fmt.Println(*p) // Nil-pointer dereference
}
