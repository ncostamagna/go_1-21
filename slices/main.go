package main

import (
	"fmt"
	"slices"
	"strconv"
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
	fmt.Printf("replaceTest variable: %v\n", replaceTest)
	replaceTest = slices.Replace(replaceTest, 3, 3, 10, 11, 12)
	fmt.Printf("Replace: %v\n", replaceTest)

	// Clone
	cloneTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("cloneTest value: %v\n", cloneTest)
	clonedTest := slices.Clone(cloneTest)
	cloneTest[2] = 5 // change same value
	fmt.Printf("Clone: %v & %v\n", cloneTest, clonedTest)

	// Compare 
	// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2
	compare1Test := []int{3, 2, 4, 3, 1, 2,4,6}
	compare2Test := []int{3, 2, 4, 3, 1, 2,4,6}

	r := slices.Compare(compare1Test, compare2Test)
	fmt.Printf("Compare: %d - %v & %v\n", r, compare1Test, compare2Test)

	compare2Test[2] = 5
	r = slices.Compare(compare1Test, compare2Test)
	fmt.Printf("Compare: %d - %v & %v\n", r, compare1Test, compare2Test)


	// Max 
	maxTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("Max: %d\n", slices.Max(maxTest))

	// Min
	minTest := []int{3, 2, 4, 3, 1, 2,4,6}
	fmt.Printf("Min: %d\n", slices.Min(minTest))


	// BinarySearch
	binarySrcTest := []int{3, 2, 4, 7,3, 1, 2,4,6}
	i, found := slices.BinarySearch(binarySrcTest, 7)
	fmt.Printf("BinarySearch - found: %t | position: %d\n", found, i)
	
	// the slice must be ordered
	slices.Sort(binarySrcTest)
	
	i, found = slices.BinarySearch(binarySrcTest, 7)
	fmt.Printf("BinarySearch - found: %t | position: %d\n", found, i)


	// Compact 
	// Compact replaces consecutive runs of equal elements with a single copy.
	compactTest := []int{1,1,2,2,9,9,3,3,2,1,10,10,5,1}
	fmt.Printf("Compact: %v\n", slices.Compact(compactTest))

	slices.Sort(compactTest)
	fmt.Printf("Compact: %v\n", slices.Compact(compactTest))


	// Index returns the index of the first occurrence of v in s, or -1 if not present.
	indexTest := []int{1,2,3,1,2,8}
	fmt.Printf("Index: %d\n", slices.Index(indexTest, 8))
	fmt.Printf("Index: %d\n", slices.Index(indexTest, 2))
	fmt.Printf("Index: %d\n", slices.Index(indexTest, 9))
	
	isSortedTest := []int{1,2,3,1,2,8}
	fmt.Printf("IsSorted: %t\n", slices.IsSorted(isSortedTest))
	slices.Sort(isSortedTest)
	fmt.Printf("IsSorted: %t\n", slices.IsSorted(isSortedTest))

	// Insert inserts the values v... into  at index i, returning the modified slice.
	insertTest := []int{1,2,3,3,1,2,8,1}
	fmt.Printf("Insert: %v\n", slices.Insert(insertTest, 4, 10, 20, 22))
	fmt.Printf("Insert: %v\n", slices.Insert(insertTest, 6, 11, 11, 11))


	containsTest := []int{1,2,3,3,1,2,8,1}
	fmt.Printf("Contains: %v\n", slices.Contains(containsTest, 3))
	fmt.Printf("Contains: %v\n", slices.Contains(containsTest, 9))

	equalTest := []int{40,1,5,1,3}
	fmt.Printf("Equal: %t\n", slices.Equal(equalTest, []int{40,5,1,1,3}))
	fmt.Printf("Equal: %t\n", slices.Equal(equalTest, []int{4,3}))
	fmt.Printf("Equal: %t\n", slices.Equal(equalTest, []int{40,1,5,1,3}))


	numbers := []int{0, 42, 8}
	strings := []string{"000", "42", "8"}
	equal := slices.EqualFunc(numbers, strings, func(n int, s string) bool {
		sn, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return false
		}
		return n == int(sn)
	})
	fmt.Printf("EqualFunc: %t\n", equal)

	fmt.Println()
}