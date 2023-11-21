package main

import "fmt"

func generate(numRows int) []int {
	if numRows == 0 {
		return []int{1}
	}
	result := make([]int, numRows+1)

	for i := 1; i <= numRows+1; i++ {
		item := make([]int, i)
		for k := range item {
			if k == 0 || k == len(item)-1 {
				item[k] = 1
			} else {
				arr := result
				item[k] = arr[k] + arr[k-1]
			}
		}
		copy(result, item)
	}
	return result
}

func main() {
	fmt.Println(generate(3))
}
