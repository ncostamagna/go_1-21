package main

import (
	"fmt"
	"slices"
)

func main(){


	// remove from index 1 to index 4
	deleteTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("deleteTest variable: %v\n", deleteTest)
	fmt.Printf("Delete: %v\n", slices.Delete(deleteTest, 1,4))

	// reverse slice
	reverseTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("reverseTest variable: %v\n", reverseTest)
	slices.Reverse(reverseTest)
	fmt.Printf("Reverse: %v\n", reverseTest)
	
	// sort
	sortTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("sortTest variable: %v\n", sortTest)
	slices.Sort(sortTest)
	fmt.Printf("Sort: %v\n", sortTest)

	// replace
	replaceTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Println("replaceTest: ", replaceTest)
	replaceTest = slices.Replace(replaceTest, 1, 3, 10, 11, 12)
	fmt.Printf("\nReplace: %v", replaceTest)

	// Clone
	cloneTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Println("cloneTest: ", cloneTest)
	clonedTest := slices.Clone(cloneTest)
	cloneTest[2] = 5 // changue same value
	fmt.Printf("\nClone: %v & %v", cloneTest, clonedTest)

	// Compare 
	// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2
	compare1Test := []int{3, 2, 4, 3, 1, 2,4,6}
	compare2Test := []int{3, 2, 4, 3, 1, 2,4,6}

	r := slices.Compare(compare1Test, compare2Test)
	fmt.Printf("\nCompare: %d - %v & %v", r, compare1Test, compare2Test)

	compare2Test[2] = 5
	r = slices.Compare(compare1Test, compare2Test)
	fmt.Printf("\nCompare: %d - %v & %v", r, compare1Test, compare2Test)


	// Max 
	maxTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("\nMax: %d", slices.Max(maxTest))

	// Min
	minTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("\nMin: %d", slices.Min(minTest))


	// BinarySearch
	binarySrcTest := []int{3, 2, 4, 3, 1, 2,4,6}
	v, found := slices.BinarySearch(binarySrcTest, 5)
	fmt.Printf("\nBinarySearch - found: %t | value: %d", found, v)
	
	// the slice must be ordered
	slices.Sort(binarySrcTest)
	v, found = slices.BinarySearch(binarySrcTest, 5)
	fmt.Printf("\nBinarySearch - found: %t | value: %d", found, v)


	// Compact 
	// Compact replaces consecutive runs of equal elements with a single copy.
	compactTest := []int{1,1,2,2,9,9,3,3,2,1,10,10,5,1}
	fmt.Printf("\nCompact: %v", slices.Compact(compactTest))
	slices.Sort(compactTest)
	fmt.Printf("\nCompact: %v", slices.Compact(compactTest))


	// Index returns the index of the first occurrence of v in s, or -1 if not present.
	indexTest := []int{1,2,3,1,2,8}
	fmt.Printf("\nIndex: %d", slices.Index(indexTest, 8))
	fmt.Printf("\nIndex: %d", slices.Index(indexTest, 2))
	fmt.Printf("\nIndex: %d", slices.Index(indexTest, 9))
	
	isSortedTest := []int{1,2,3,1,2,8}
	fmt.Printf("\nIsSorted: %t", slices.IsSorted(isSortedTest))
	slices.Sort(isSortedTest)
	fmt.Printf("\nIsSorted: %t", slices.IsSorted(isSortedTest))

	// Insert inserts the values v... into  at index i, returning the modified slice.
	insertTest := []int{1,2,3,3,1,2,8,1}
	fmt.Printf("\nInsert: %v", slices.Insert(insertTest, 4, 10, 20, 22))
	fmt.Printf("\nInsert: %v", slices.Insert(insertTest, 6, 11, 11, 11))


	containsTest := []int{1,2,3,3,1,2,8,1}
	fmt.Printf("\nContains: %v", slices.Contains(containsTest, 3))
	fmt.Printf("\nContains: %v", slices.Contains(containsTest, 9))

	fmt.Println()
}