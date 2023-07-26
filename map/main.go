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

	fmt.Printf("Values: %v\n", maps.Values(myMap))
	fmt.Printf("Keys: %v\n", maps.Keys(myMap))

	myOtherMap := maps.Clone(myMap)
	fmt.Printf("First Map %v - Clone %v\n", myMap, myOtherMap)

	fmt.Printf("Compare %t\n", maps.Equal(myMap, myOtherMap) )

	maps.DeleteFunc(myMap, func(k string, v int) bool {
		return k == "one"
	})
	fmt.Printf("with Deleted value %v\n", myMap)

	fmt.Printf("Compare %t\n", maps.Equal(myMap, myOtherMap) )

	myOtherMap["four"] = 44
	myOtherMap["nine"] = 9
	myMap["ten"] = 10
	fmt.Printf("First Map %v - Clone %v\n", myMap, myOtherMap)
	maps.Copy(myMap, myOtherMap)
	fmt.Printf("First Map %v - Clone %v\n", myMap, myOtherMap)


}