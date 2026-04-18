package main

import "fmt"

func main() {
	const a int = 5

	fmt.Println(a)

	//iota

	const (
		sunday = iota
		monday
		tuesday
	)

	fmt.Println(sunday,monday,tuesday)
}
