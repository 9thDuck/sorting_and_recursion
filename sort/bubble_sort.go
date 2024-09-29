package sort

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Integer](arrP *[]T) *[]T {
	arr := *arrP
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			nextIdx := j + 1
			if arr[j] > arr[nextIdx] {
				arr[j], arr[nextIdx] = arr[nextIdx], arr[j]
			}
		}
	}

	return arrP
}
