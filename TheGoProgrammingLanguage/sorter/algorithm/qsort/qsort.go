package qsort

func getIndex(values []int, low, high int) int {
	temp := values[low]

	for low < high {
		for low < high && values[high] >= temp {
			high--
		}
		values[low] = values[high]
		for low < high && values[low] <= temp {
			low++
		}
		values[high] = values[low]
	}
	values[low] = temp
	return low
}

func quickSort(values []int, low, high int) {
	if low < high {
		index := getIndex(values, low, high)
		quickSort(values, low, index-1)
		quickSort(values, index+1, high)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
