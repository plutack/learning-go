package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{6, 7, 1, 3, 7, 4, 2}))
}

func bubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
