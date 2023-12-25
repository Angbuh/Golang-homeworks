package binarysearch

import "cmp"

func BinarySearch[E cmp.Ordered](arr []E, elem E) (int, bool) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		switch {
		case guess == elem:
			return mid, true
		case guess > elem:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	return 0, false
}
