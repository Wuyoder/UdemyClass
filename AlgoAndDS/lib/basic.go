package lib

import (
	"reflect"
	"strings"
)

// func CallFunc() interface{} {
// 	// return charCount("Hello, World!")
// 	// return getASCII("azAZ09")
// 	// return decodeASCII([]int{97, 98, 99})
// 	// return same([]int{1, 2, 3}, []int{1, 4, 9, 9})
// 	// return anagram("asx", "xasa")
// 	// return sumZero([]int{-3, -2, -1, 0, 1, 2, 3, 4})
// 	// return countUniqueValues([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13})
// 	// return maxSubArraySum([]int{1, 2, 5, 2, 8, 1, 10}, 3)
// 	// return searchNum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
// 	// return recursivePower(2, 2)
// 	// return recursiveFactrial(7)
// 	// return recursiveProductOfArray([]int{1, 2, 3, 10, 5})
// 	// return recursiveRange(100)
// 	// return recursiveFib(28)
// 	// return recursiveReverse("awesome")
// 	// return isPalindrome("asdfghjhgfdsa")
// 	// return someRecursive([]int{1, 2, 30, 4}, func(i int) bool { return i > 10 })
// 	// return flatten([]interface{}{[]interface{}{1}, []interface{}{2}, []interface{}{3}})
// 	// obj1 := map[string]interface{}{
// 	// 	"outer": 2,
// 	// 	"obj": map[string]interface{}{
// 	// 		"inner": 2,
// 	// 		"otherObj": map[string]interface{}{
// 	// 			"superInner":     2,
// 	// 			"notANumber":     true,
// 	// 			"alsoNotANumber": "yup",
// 	// 		},
// 	// 	},
// 	// }
// 	// return nestedEvenSum(obj1)
// 	// return capitalizeWords([]string{"i", "am", "learning", "recursion"})
// }

func capitalizeWords(s []string) []string {
	if len(s) == 1 {
		return []string{strings.ToUpper(s[0])}
	}
	return append([]string{strings.ToUpper(s[0])}, capitalizeWords(s[1:])...)
}

func nestedEvenSum(obj map[string]interface{}) int {
	sum := 0
	for _, v := range obj {
		switch t := v.(type) {
		case map[string]interface{}:
			sum += nestedEvenSum(t)
		case int:
			if t%2 == 0 {
				sum += t
			}
		}
	}
	return sum
}

func flatten(arr interface{}) []interface{} {
	flatArr := []interface{}{}
	for _, value := range arr.([]interface{}) {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			flatArr = append(flatArr, flatten(value)...)
		default:
			flatArr = append(flatArr, value)
		}
	}
	return flatArr
}

func someRecursive(i []int, f func(int) bool) bool {
	if len(i) == 0 {
		return false
	}

	if f(i[0]) {
		return true
	}

	return someRecursive(i[1:], f)
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func recursiveReverse(s string) string {
	if len(s) == 1 {
		return s
	}
	return s[len(s)-1:] + recursiveReverse(s[:len(s)-1])
}

func recursiveFib(s int) int {
	if s <= 2 {
		return 1
	}
	return recursiveFib(s-1) + recursiveFib(s-2)
}

func recursiveRange(s int) int {
	if s == 1 {
		return 1
	}
	return s + recursiveRange(s-1)
}

func recursiveProductOfArray(i []int) int {
	if len(i) == 1 {
		return i[0]
	}
	return i[0] * recursiveProductOfArray(i[1:])
}

func recursiveFactrial(f int) int {
	if f < 0 {
		return 0
	}
	if f == 1 {
		return 1
	}
	return f * recursiveFactrial(f-1)
}

func recursivePower(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * recursivePower(base, exponent-1)
}

func searchNum(num []int, n int) int {
	//binary search
	low, high := 0, len(num)-1
	for low <= high {
		mid := (low + high) / 2
		if num[mid] == n {
			return mid
		} else if num[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func maxSubArraySum(num []int, n int) int {
	tempMax, Max := 0, 0
	for i := 0; i < n; i++ {
		tempMax += num[i]
	}

	Max = tempMax
	for j := 1; j < len(num)-n+1; j++ {
		tempMax += num[j+n-1] - num[j-1]
		if tempMax > Max {
			Max = tempMax
		}
	}

	return Max
}

func countUniqueValues2(num []int) int {
	// space complexity O(n)
	numMap := map[int]bool{}
	for _, n := range num {
		numMap[n] = true
	}

	return len(numMap)
}

func countUniqueValues(num []int) int {
	// space complexity O(1)
	if len(num) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(num); j++ {
		if num[i] != num[j] {
			i++
			num[i] = num[j]
		}
	}

	return i + 1
}

func sumZero(num []int) (result [][]int) {
	e := len(num) - 1
	s := 0

	if len(num) < 2 {
		return
	}

	result = [][]int(nil)
	for s < e {
		sum := num[s] + num[e]
		if sum == 0 {
			result = append(result, []int{num[s], num[e]})
			for s < e && num[s] == num[s+1] {
				s++
			}
			for s < e && num[e] == num[e-1] {
				e--
			}
		} else if sum > 0 {
			e--
		} else {
			s++
		}
	}

	return
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s2Map := map[rune]bool{}
	for _, s := range s2 {
		s2Map[s] = true
	}

	for _, c := range s1 {
		if !s2Map[c] {
			return false
		}
	}

	return true
}

func same(root, square []int) bool {
	if len(root) != len(square) {
		return false
	}
	sMap := map[int]bool{}
	for _, s := range square {
		sMap[s] = true
	}
	for _, r := range root {
		if !sMap[r*r] {
			return false
		}
	}
	return true
}

func charCount(s string) map[string]int {
	result := map[string]int{}
	for _, c := range s {
		key := strings.ToLower(string(c))
		// if key == " " {
		// 	continue
		// }
		result[key]++
	}
	return result
}

func getASCII(c string) map[string]int {
	result := map[string]int{}
	for _, s := range c {
		result[string(s)] = int(s)
	}
	return result
}

func decodeASCII(s []int) string {
	result := ""
	for _, i := range s {
		result += string(i)
	}
	return result
}
