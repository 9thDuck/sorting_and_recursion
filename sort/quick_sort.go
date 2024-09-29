package sort

import (
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Integer](arrP *[]T) *[]T {
	arr := *arrP
	if len(arr) <= 1 {
		return arrP
	}
	i, j, pivotIdx := -1, 0, len(arr)-1

	for j <= pivotIdx {
		var temp T
		if j == pivotIdx {
			temp = arr[pivotIdx]
			i++
			arr[pivotIdx] = arr[i]
			arr[i] = temp
			pivotIdx = i
			j++
		} else if arr[j] >= arr[pivotIdx] {
			j++
		} else {
			i++
			temp = arr[j]
			arr[j] = arr[i]
			arr[i] = temp
			j++
		}
	}

	leftArrP := arr[:pivotIdx]
	rightArrP := arr[pivotIdx+1:]

	left := *QuickSort(&leftArrP)
	right := *QuickSort(&rightArrP)

	res := append(append(left, arr[pivotIdx]), right...)
	return &res
}
