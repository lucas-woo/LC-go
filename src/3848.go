package main

func isDigitorialPermutation(n int) bool {

	numMap := make(map[int]int, 0)
	copy := n
	sumFactorial := 0
	for copy > 0 {
		cur := copy % 10;
		sumFactorial += findFactorial(cur)
		copy /= 10
		if _, ok := numMap[cur]; !ok {
			numMap[cur] = 1
			continue
		}
		numMap[cur]++;
	}

	numMapTwo := make(map[int]int, 0)
	another := sumFactorial
	for another > 0 {
		cur := another % 10;
		another /= 10
		if _, ok := numMapTwo[cur]; !ok {
			numMapTwo[cur] = 1
			continue
		}
		numMapTwo[cur]++;
	}	
	
	if len(numMap) != len(numMapTwo) {
		return false
	}
	for k, v := range numMap {
		val, ok := numMapTwo[k]
		if !ok || v != val {
			return false
		}
	}
	return true
}

func findFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * findFactorial(n - 1)
}