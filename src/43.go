package main

func permute(nums []int) [][]int {
  finalArr := make([][]int, 0)
	permuteArr([]int{}, nums, &finalArr)
	return finalArr
}

func permuteArr(cur []int, available []int, arr *[][]int) {
	if len(available) == 0 {
		*arr = append(*arr, cur)
		return
	}

	for i := 0; i < len(available); i++ {
		copy := make([]int, 0)
		copyAvailable := make([]int, 0)
		for j := 0; j < len(cur); j++ {
			copy = append(copy, cur[j])
		}
		copy = append(copy, available[i])
		for j := 0; j < len(available); j++ {
			if available[j] == available[i] {
				continue
			}
			copyAvailable = append(copyAvailable, available[j])
		}
		permuteArr(copy, copyAvailable, arr)
	}
}