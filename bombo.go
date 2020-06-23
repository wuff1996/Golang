package main

import "fmt"

func main() {

	arr := []int{4, 5, 6, 7, 3, 453, 5, 7, 8, 33, 44}
	for i, _ := range arr {
		for j := 0; j+i < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

}
