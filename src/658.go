package main

func findClosestElements(arr []int, k int, x int) []int {

	finalArr := make([]int, 0)
	if x <= arr[0] {
		for i := 0; i < k; i++ {
			finalArr = append(finalArr, arr[i])
		}
		return finalArr
	}

	if x >= arr[len(arr)-1] {
		for i := len(arr) - k; i < len(arr); i++ {
			finalArr = append(finalArr, arr[i])
		}
		return finalArr
	}



	found := -1
	prev := arr[0]
	notFound := true
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			found = i;
			notFound = false
			break
		}
		if x > prev && x < arr[i] {
			found = i
			break;
		}
		prev = arr[i]
	}

	if notFound {
		left := found - 1
		right := found

		for i := 0; i < k; i++ {
			if left <= -1 {
				right++;
				continue
			} 
			if right >= len(arr) {
				left--;
				continue
			}
			if x - arr[left] <= arr[right] - x {
				left--;
			} else {
				right++;
			}
		}

		left++;

		for i := 0; i < k; i++ {
			finalArr = append(finalArr, arr[left])
			left++;
		}
		return finalArr		
	}

	left := found - 1;
	right := found + 1;

	for i := 0; i < k - 1; i++ {
		if left <= -1 {
			right++;
			continue
		} 
		if right >= len(arr) {

			left--;
			continue
		}
		if (x - arr[left]) <= (arr[right] - x) {
			left--;
		} else {
			right++;
		}
	}
	left++
	for i := 0; i < k; i++ {
		finalArr = append(finalArr, arr[left])
		left++;
	}
	return finalArr
}