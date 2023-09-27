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

	// Those methods were removed: https://github.com/golang/go/commit/b25266c58debb831cd880a7372b197466e75b833
	// fmt.Printf("Values: %v\n", maps.Values(myMap))
	// fmt.Printf("Keys: %v\n", maps.Keys(myMap))

	myOtherMap := maps.Clone(myMap)
	fmt.Printf("First Map %v \nClone %v\n", myMap, myOtherMap)

	r := maps.Equal(myMap, myOtherMap)

	fmt.Printf("myMap:      %v \nmyOtherMap: %v\n", myMap, myOtherMap)
	fmt.Printf("Compare %t\n", r )

	//myMap["one"] = 11
	r = maps.Equal(myMap, myOtherMap)

	fmt.Printf("\n\nmyMap:      %v \nmyOtherMap: %v\n", myMap, myOtherMap)
	fmt.Printf("Compare %t\n", r )

	fmt.Printf("Original Map: %v\n", myMap)
	maps.DeleteFunc(myMap, func(k string, v int) bool {
		return k == "one"
	})
	fmt.Printf("with Deleted value %v\n", myMap)

	fmt.Printf("Compare %t\n", maps.Equal(myMap, myOtherMap) )

	fmt.Printf("\n---- Copy ----\n")

	dst := map[string]int{
		"four":4, 
		"ten":10,
		"three":3,
		"two":2,
	}

	src := map[string]int {
		"four":44, 
		"nine":9,
		"one":1,
		"three":3,
		"two":2,
	}

	fmt.Printf("dst: %v \nsrc: %v\n", dst, src)

	maps.Copy(dst, src)
	
	fmt.Printf("\ncopied variable (dst): %v\n", dst)


}