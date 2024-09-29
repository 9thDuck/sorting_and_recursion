package utils

import "golang.org/x/exp/constraints"

func SplitArrInTheMiddle[T constraints.Integer](arr []T) ([]T, []T) {
	middle := len(arr) / 2
	return arr[:middle], arr[middle:]
}
