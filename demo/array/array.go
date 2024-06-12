package main

import (
	"fmt"
)

func exampleArrayAndSlice() {
	var arr [5]int
	for i :=0; i < len(arr); i++ {
		arr[i] = i
	}

	// 切片左闭右开
	var slice1 []int = arr[2:5]
	// 等同于slice1 := arr[2:5]?
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice1[i])
	}

	slice2 := append(slice1, 6)
	slice2[0] = 0
	arr[2] = 100
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice2[i])
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice1[i])
	}
}

func main() {
	exampleArrayAndSlice()
}