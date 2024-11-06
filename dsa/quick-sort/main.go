package main

import "fmt"

func main() {
	array := []int{4, 5, 1, 3, 6, 8, 1, 7, 3}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	var i, j int
	for i, j = low-1, low; j < high; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	i++
	a[i], a[high] = a[high], a[i]
	return i
}

func quickSort(a []int, low, high int) []int {
	if low < high {

		pivot := partition(a, low, high)

		quickSort(a, low, pivot-1)
		quickSort(a, pivot+1, high)
	}
	return a
}
