package main

import "fmt"


func main (){

	//arithmetic operators
	a := 5
	b:=4

	fmt.Println(a +b)
	fmt.Println(a -b)
	fmt.Println(a *b)
	fmt.Println(a /b)
	fmt.Println(a %b)

	a ++ // increment by 1
	a -- // decrement by 1



	// && logical and operator
	fmt.Print(a && b)

	// logical || or operator
fmt.Println(a || b)

	// logical not operator !

	fmt.Print(!a)

}
