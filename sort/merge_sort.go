package sort

import (
	"example.com/recursion/utils"
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Integer](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	// fmt.Println(arr)
	left, right := utils.SplitArrInTheMiddle[T](arr)
	// fmt.Println(left)
	// fmt.Println(right)
	return merge(MergeSort[T](left), MergeSort[T](right))
}

func merge[T constraints.Integer](left, right []T) []T {
	i, j := 0, 0

	leftLen := len(left)
	rightLen := len(right)
	// fmt.Println("merge:")
	// fmt.Println("left", left)
	// fmt.Println("right", right)

	result := make([]T, 0, leftLen+rightLen)
	for i < leftLen && j < rightLen {
		// fmt.Print("iter start", result, " ")

		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
		// fmt.Println("iter end", result)
	}
	result = append(result, left[i:]...)
	// fmt.Println("spread left", result)

	result = append(result, right[j:]...)
	// fmt.Println("spread right", result)

	// fmt.Print("result", result, "\n\n")

	return result
}
