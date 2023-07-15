package main

import (
	"fmt"
	"maps"
)

func main(){

	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	// Values and Keys
	fmt.Printf("Values: %v\n", maps.Values(myMap))
	fmt.Printf("Keys: %v\n", maps.Keys(myMap))

	// Clone
	myOtherMap := maps.Clone(myMap)
	fmt.Printf("First Map %v - Clone %v\n", myMap, myOtherMap)

	// Compare
	fmt.Printf("Compare %t\n", maps.Equal(myMap, myOtherMap) )

	// Delete
	maps.DeleteFunc(myMap, func(k string, v int) bool {
		return k == "one"
	})
	fmt.Printf("with Deleted value %v\n", myMap)

	fmt.Printf("Compare %t\n", maps.Equal(myMap, myOtherMap) )

	// Copy copies all key/value pairs in src adding them to dst. When a key in src is already present in dst, the value in dst will be overwritten by the value associated with the key in src.
	myOtherMap["four"] = 44
	myOtherMap["nine"] = 9
	myMap["ten"] = 10
	fmt.Printf("First Map %v - Clone %v\n", myMap, myOtherMap)
	maps.Copy(myMap, myOtherMap)
	fmt.Printf("First Map %v - Clone %v\n", myMap, myOtherMap)


}