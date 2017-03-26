package main

import "fmt"

func getRow(rowIndex int) []int {
	result := []int{}

	for i := 0; i <= rowIndex; i++ {
		result = append([]int{1},result...)
		for j := 1; j < len(result) - 1; j++ {
			result[j] += result[j+1]
		}
	}

	return result
}

func main()  {
	result := getRow(4)

	for _,v := range result {
		fmt.Print(v, " ")
	}
}