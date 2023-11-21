package main

import "fmt"

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	arr := []int{4, 11, 4, 2, 1, 1, 2}
	fmt.Println(singleNumber(arr))
}
