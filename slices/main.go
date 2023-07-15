package main

import (
	"fmt"
	"slices"
)

func main(){
	mySlice := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Println("My Slice", mySlice)

	// remove from index 1 to index 4
	fmt.Printf("\n\nDelete: %v", slices.Delete(mySlice, 1,4))

	// reverse slice
	slices.Reverse(mySlice)
	fmt.Printf("\nReverse: %v", mySlice)
	
	// sort
	slices.Sort(mySlice)
	fmt.Printf("\nSort: %v", mySlice)

	// replace
	mySlice = slices.Replace(mySlice, 1, 3, 10, 11, 12)
	fmt.Printf("\nReplace: %v", mySlice)

	// Clone
	myOtherSlice := slices.Clone(mySlice)
	mySlice[2] = 5 // changue same value
	fmt.Printf("\nClone: %v & %v", mySlice, myOtherSlice)

	// Compare 
	// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2
	fmt.Printf("\nCompare: %d - %v & %v", slices.Compare(mySlice, myOtherSlice), mySlice, myOtherSlice)
	myOtherSlice[2] = 5
	fmt.Printf("\nCompare: %d - %v & %v", slices.Compare(mySlice, myOtherSlice), mySlice, myOtherSlice)




	// Max 
	fmt.Printf("\nMax: %d", slices.Max(mySlice))

	// Min
	fmt.Printf("\nMin: %d", slices.Min(mySlice))


	// BinarySearch
	v, found := slices.BinarySearch(mySlice, 5)
	fmt.Printf("\nBinarySearch - found: %t | value: %d", found, v)
	// the slice must be ordered
	slices.Sort(mySlice)
	v, found = slices.BinarySearch(mySlice, 5)
	fmt.Printf("\nBinarySearch - found: %t | value: %d", found, v)


	// Compact 
	// Compact replaces consecutive runs of equal elements with a single copy.
	mySlice3 := []int{1,1,2,2,9,9,3,3,2,1,10,10,5,1}
	fmt.Printf("\nCompact: %v", slices.Compact(mySlice3))
	slices.Sort(mySlice3)
	fmt.Printf("\nCompact: %v", slices.Compact(mySlice3))

	// Grow
	//s := slices.Grow(mySlice3, 500)
	//fmt.Printf("\n %v %v %d %d", mySlice3, s, cap(mySlice3), cap(s))

	// Index returns the index of the first occurrence of v in s, or -1 if not present.
	mySlice4 := []int{1,2,3,1,2,8}
	fmt.Printf("\nIndex: %d", slices.Index(mySlice4, 8))
	fmt.Printf("\nIndex: %d", slices.Index(mySlice4, 2))
	fmt.Printf("\nIndex: %d", slices.Index(mySlice4, 9))
	
	mySliceIsSortedTest := []int{1,2,3,1,2,8}
	fmt.Printf("\nIsSorted: %t", slices.IsSorted(mySliceIsSortedTest))
	slices.Sort(mySliceIsSortedTest)
	fmt.Printf("\nIsSorted: %t", slices.IsSorted(mySliceIsSortedTest))

	// Insert inserts the values v... into mySlice at index i, returning the modified slice.
	mySliceInsertTest := []int{1,2,3,3,1,2,8,1}
	fmt.Printf("\nInsert: %v", slices.Insert(mySliceInsertTest, 4, 10, 20, 22))
	fmt.Printf("\nInsert: %v", slices.Insert(mySliceInsertTest, 6, 11, 11, 11))


	mySliceContainsTest := []int{1,2,3,3,1,2,8,1}
	fmt.Printf("\nContains: %v", slices.Contains(mySliceContainsTest, 3))
	fmt.Printf("\nContains: %v", slices.Contains(mySliceContainsTest, 9))

	fmt.Println()
}