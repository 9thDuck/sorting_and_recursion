package main

import (
	"fmt"

	"example.com/recursion/sort"
)

// func recursiveDifferenceCal(arrA, arrB *[]int, idx int, differenceSum float64) float64 {
// 	if len(*arrA) == idx || len(*arrB) == idx {
// 		return differenceSum
// 	}

// 	differenceSum += math.Abs(float64((*arrA)[idx] - (*arrB)[idx]))
// 	return recursiveDifferenceCal(arrA, arrB, idx+1, differenceSum)
// }

// func basicRecurion() {
// 	sensorA := []int{15, -4, 56, 10, -23}
// 	sensorB := []int{14, -9, 56, 14, -23}

// 	var differenceSum float64

// 	// for i := 0; i < 5; i++ {
// 	// 	differenceSum += math.Abs(float64(sensorA[i] - sensorB[i]))
// 	// }

// 	differenceSum = recursiveDifferenceCal(&sensorA, &sensorB, 0, differenceSum)
// 	fmt.Println(differenceSum)
// }

func main() {
	arr := []int{5, 150, 3, 1, 99, 6, 42, 11, 9}
	// sorted := sort.MergeSort(arr)
	sorted := *sort.QuickSort[int](&arr)
	fmt.Println(sorted)
}
