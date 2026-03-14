package main

import "sort"

func getHappyString(n int, k int) string {
	arr := make([]string, 3)
	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"

	for len(arr[0]) < n {
		newArr := make([]string, 0)
		for i := 0; i < len(arr); i++ {
			if arr[i][len(arr[i]) - 1] == 'a' {
				newArr = append(newArr, arr[i] + "b")
				newArr = append(newArr, arr[i] + "c")
			} else if arr[i][len(arr[i]) - 1] == 'b' {
				newArr = append(newArr, arr[i] + "a")
				newArr = append(newArr, arr[i] + "c")
			} else {
				newArr = append(newArr, arr[i] + "a")
				newArr = append(newArr, arr[i] + "b")
			}
		}
		arr = newArr
	}
	if k > len(arr) {
		return ""
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[k - 1]
}