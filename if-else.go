package main

import "fmt"

func main() {

	num := 10

	if num > 0 {
		fmt.Println(num, "Is positive")
	}
	if num <10  {
		fmt.Println(num, "Num is not large than this number")
	}


// if with function

if isPositive (num){
	fmt.Println(num , "is positive")
} else{
	fmt.Println(num, "is negative number")
}


}


func isPositive (num int) bool {
}
