package main

import "fmt"

func main() {
	slice := makeSlice(0, 10)

	for _, v := range slice {
		fmt.Println(v, "is", evenOrOdd(v))
	}
}

func makeSlice(start int, end int) []int {
	slice := []int{}
	for i := start; i < end+1; i++ {
		slice = append(slice, i)
	}
	return slice
}

func evenOrOdd(value int) string {
	if value%2 == 0 {
		return "even"
	}

	return "odd"

}
