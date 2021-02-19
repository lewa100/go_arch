package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func BinarySearch(slice []int, value int) (int, int) {
	steps := 0
	start_index := 0
	end_index := len(slice) - 1

	for start_index <= end_index {

		median := (start_index + end_index) / 2

		if slice[median] < value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}
		steps++
	}

	if start_index == len(slice) || slice[start_index] != value {
		return -1, steps
	} else {
		return start_index, steps
	}

}

//Сложность log n
//Алгоритм не рабочий, если не соблюдается порядок в наборе данных

func main() {

	searchArray1 := []int{1, 22, 48, 50, 60, 65, 70, 90, 100, 120, 121, 131, 141, 151, 161, 162}
	searchArray2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	searchArray3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64}
	searchArray4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128}

	testSearch(searchArray1)
	testSearch(searchArray2)
	testSearch(searchArray3)
	testSearch(searchArray4)
}

func testSearch(searchArray []int) {
	var countSteps []float64
	for _, v := range searchArray {
		_, steps := BinarySearch(searchArray, v)
		countSteps = append(countSteps, float64(steps))
	}
	_, steps := BinarySearch(searchArray, 0)
	countSteps = append(countSteps, float64(steps))

	fmt.Printf("searchArray: %d\n", len(searchArray))
	graph := asciigraph.Plot(countSteps)
	fmt.Println(graph, "\n")
}
