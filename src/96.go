package main


func numTrees(n int) int {
	var someMap map[int]int = make(map[int]int, 0)
	someMap[0] = 1;
	someMap[1] = 1;
	someMap[2] = 2

	if n <= 2 {
		return someMap[n]
	}

	for i := 3; i <= n; i++ {
		createTree(i, someMap)
	}
	return someMap[n]
}

func createTree(n int, someMap map[int]int) {
	for i := 1; i <= n; i++ {
		left := max(i - 1, 1)
		right := max(1, n - i)
		someMap[n] += someMap[left] * someMap[right]
	}
}