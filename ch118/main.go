package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 1; i <= numRows; i++ {
		result = append(result, make([]int, i))
		item := make([]int, i)
		for k := range item {
			if k == 0 || k == len(item)-1 {
				item[k] = 1
			} else {
				arr := result[i-2]
				item[k] = arr[k] + arr[k-1]
			}
		}
		result[i-1] = item
	}
	return result
}

func main() {
	fmt.Println(generate(5))
}
