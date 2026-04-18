package main

import "fmt"

// global scope or global variable

var globalVar = "Globar variable"



func main (){
//local scope or local variable

var localVar = "Local variable"

fmt.Println(globalVar)
fmt.Println(localVar)

}


func anotherFunc (){
	fmt.Println(globalVar) // this is access for everywhere
	// fmt.Println(localVar) // this is not allowed, this is just for block scope
}
