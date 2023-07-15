package main

import (
	"cmp"
	"math"
	"fmt"
)
func main(){
	myNanValue := math.NaN()
	var myFloatValue float64 = 20

	if myNanValue < myFloatValue {
		fmt.Println("test")
	}
	if myNanValue > myFloatValue {
		fmt.Println("test")
	}
	if myNanValue == myFloatValue {
		fmt.Println("test")
	}

	// For floating-point types, a NaN is considered less than any non-NaN, a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
	fmt.Printf("Numbers (%f, %f) - Compare: %d", myNanValue,myFloatValue, cmp.Compare(myNanValue, myFloatValue) )

	// Less reports whether x is less than y. For floating-point types, a NaN is considered less than any non-NaN, and -0.0 is not less than (is equal to) 0.0.
	fmt.Printf("\nNumbers (%f, %f) - Less: %t", myNanValue,myFloatValue, cmp.Less(myNanValue, myFloatValue) )

	fmt.Println()


}