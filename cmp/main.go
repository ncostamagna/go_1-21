package main

import (
	"cmp"
	"math"
	"fmt"
)
func main(){
	myNanValue := math.NaN()
	var myFloatValue float64 = -20000

	if myNanValue < myFloatValue {
		fmt.Println("test")
	}
	if myNanValue > myFloatValue {
		fmt.Println("test")
	}
	if myNanValue == myFloatValue {
		fmt.Println("test")
	}

	fmt.Printf("Numbers (%f, %f) - Compare: %d\n", myNanValue,myFloatValue, cmp.Compare(myNanValue, myFloatValue) )

	fmt.Printf("Numbers (%f, %f) - Less: %t\n", myNanValue,myFloatValue, cmp.Less(myNanValue, myFloatValue) )

fmt.Println()
}