package lib

import (
	"math"
	"math/rand"
)

func CallFunc() interface{} {
	// return BubbleSort([]int{6, 5, 4, 3, 2, 1})
	// return BubbleSortWhileLoop([]int{6, 5, 4, 3, 2, 1})
	// return SelectionSort([]int{6, 5, 4, 3, 2, 1})
	// return SelectionSortWhileLoop([]int{6, 5, 4, 3, 2, 1})
	// return InsertionSort([]int{6, 5, 4, 3, 2, 1})
	// return InsertionSortWhileLoop([]int{6, 5, 4, 3, 2, 1})
	// return MergeSort([]int{6, 5, 4, 3, 2, 1})
	// return QuickSortInit([]int{6, 5, 4, 3, 2, 1})
	// return QuickSortRandom([]int{6, 5, 4, 3, 2, 1})
	return RadixSort([]int{6, 5, 4, 3, 2, 1})
	// return true
	// golang sort
	// golang slices.Sort
}

func RadixSort(arr []int) []int {
	maxDigit := mostDigits(arr)
	temp := make([]int, len(arr))
	for k := 0; k < maxDigit; k++ {
		buckets := [10][]int{}
		for i := 0; i < len(arr); i++ {
			digit := getDigit(arr[i], k)
			buckets[digit] = append(buckets[digit], arr[i])
		}
		j := 0
		for _, bubuckets := range buckets {
			copy(temp[j:], bubuckets)
			j += len(bubuckets)
		}
		copy(arr, temp)
	}

	return arr
}

func mostDigits(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		max = int(math.Max(float64(max), float64(digitCount(arr[i]))))
	}
	return max
}

func digitCount(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(num)))) + 1)
}

func getDigit(num, i int) int {
	return num / int(math.Pow(10, float64(i))) % 10
}

func QuickSortRandom(arr []int) []int {
	return QuickSortStartRandom(arr, 0, len(arr)-1) //with random pivot
}

func QuickSortStartRandom(arr []int, left, right int) []int {
	if left < right {
		pivotIndex := rand.Intn(right-left) + left
		pivotIndex = partition(arr, left, right, pivotIndex)
		QuickSortStartRandom(arr, left, pivotIndex-1)
		QuickSortStartRandom(arr, pivotIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right, pivotIndex int) int {
	pivot := arr[pivotIndex]
	// swap pivot to the end
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	swapIndex := left
	for i := left; i < right; i++ {
		if arr[i] <= pivot {
			arr[i], arr[swapIndex] = arr[swapIndex], arr[i]
			swapIndex++
		}
	}
	arr[right], arr[swapIndex] = arr[swapIndex], arr[right]
	return swapIndex
}

func QuickSortInit(arr []int) []int {
	return quickSortStartInit(arr, 0, len(arr)-1) //with 0 as start
}

func quickSortStartInit(arr []int, left, right int) []int {
	if left < right {
		pivotIndex := pivot(arr, left, right)
		quickSortStartInit(arr, left, pivotIndex-1)
		quickSortStartInit(arr, pivotIndex+1, right)
	}

	return arr
}

func pivot(arr []int, start, end int) int {
	pivot := arr[start]
	swapIndex := start
	for i := start + 1; i <= end; i++ {
		if pivot > arr[i] {
			swapIndex++
			arr[swapIndex], arr[i] = arr[i], arr[swapIndex]
		}
	}
	// swap value at start to the end of values less than it
	arr[start], arr[swapIndex] = arr[swapIndex], arr[start]
	return swapIndex
}

func MergeSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	return merge(MergeSort(arr[:l/2]), MergeSort(arr[l/2:]))
}

func merge(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	result := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	result = append(append(result, arr1[i:]...), arr2[j:]...)
	return result
}

func InsertionSortWhileLoop(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	i := 0
	for i < l {
		j := i - 1
		for j >= 0 {
			currentVal := arr[j]
			nextVal := arr[j+1]
			if currentVal > nextVal {
				arr[j], arr[j+1] = nextVal, currentVal
				j--
			} else {
				break
			}
		}
		i++
	}

	return arr
}

func InsertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 {
			currentVal := arr[j]
			nextVal := arr[j+1]
			if currentVal > nextVal {
				arr[j], arr[j+1] = nextVal, currentVal
				j--
			} else {
				break
			}
		}
	}
	return arr
}

func SelectionSortWhileLoop(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	i := 0
	for i < l-1 {
		minIndex := i
		j := i + 1
		for j < l {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
			j++
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
		i++
	}

	return arr
}

func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

func BubbleSortWhileLoop(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	isSwap := true
	for isSwap {
		isSwap = false
		i := 0
		for i < l-1 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSwap = true
			}
			i++
		}
		l--
	}

	return arr
}

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// first inner swap will solve arr[len(arr)-1], so (n-1) iterations are enough
	for i := 0; i < len(arr)-1; i++ {
		// isSwap tracks whether or not any swaps were made on this iteration
		isSwap := false
		// to len(arr)-1-i, since i position in the end already sorted
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
	return arr
}
