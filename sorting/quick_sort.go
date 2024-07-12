package sorting

func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
