package main

import "strconv"
func letterCombinations(digits string) []string {
	digitsMap := [8]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	returnArr := make([]string, 0)
	loopArr := make([]string, len(digits))

	for i := 0; i < len(digits); i++ {
		d, _ := strconv.Atoi(string(digits[i]))
		loopArr[i] = digitsMap[d-2]
	}

	for i := 0; i < len(loopArr[0]); i++ {

		one := loopArr[0][i]
		for j := 0; len(loopArr) > 1 && j < len(loopArr[1]); j++ {
			two := loopArr[1][j]
			for k := 0; len(loopArr) > 2 && k < len(loopArr[2]); k++ {
				three := loopArr[2][k]
				for l := 0; len(loopArr) > 3 && l < len(loopArr[3]); l++ {
					four := loopArr[3][l]
					returnArr = append(returnArr, string([]byte{one, two, three, four}))
				}
				if len(digits) <= 3 {
					returnArr = append(returnArr, string([]byte{one, two, three}))
				}
			}
			if len(digits) <= 2 {
				returnArr = append(returnArr, string([]byte{one, two}))
			}
		}
		if len(digits) <= 1 {
			returnArr = append(returnArr, string([]byte{one}))
		}
	}
	return returnArr
}