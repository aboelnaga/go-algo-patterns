package main

import "fmt"

func sliceBasics() {
	slice := []int{-1, 2, 3, 4, 5}
	sliceCopy := doubleSliceElements(slice)
	fmt.Println("Original slice:", slice)
	fmt.Println("Doubled slice:", sliceCopy)
	sliceCopy1 := doubleSliceElements(nil)
	fmt.Println("Doubled slice 1:", sliceCopy1)
}

func doubleSliceElements(s []int) (output []int) {
	// create a new empty slice to hold results it is defined in the function signature as a named return value
	output = make([]int, 0, len(s))
	// for each element in the input slice
	for _, v := range s {
		//   double it and append to the new slice
		output = append(output, v*2)
	}
	// return the new slice
	return output
}
