package main

func canReach(arr []int, start int) bool {

	
	visited := make(map[int]bool, 0)
	
	stack := make([]int, 0)
	stack = append(stack, start)
	for len(stack) > 0 {
		popped := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if _, ok := visited[popped]; ok {
			continue
		}
		if arr[popped] == 0 {
			return true
		}
		visited[popped] = true
		if popped + arr[popped] < len(arr) {
			stack = append(stack, popped + arr[popped])
		}
		if popped - arr[popped] >= 0 {
			stack = append(stack, popped - arr[popped])
		}
	}

	return false
}