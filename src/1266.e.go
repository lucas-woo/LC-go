package main

func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) == 1 {
		return 0
	}
	total := 0;
	for i := 1; i < len(points); i++ {
		total += max(max(points[i][0] - points[i-1][0], points[i-1][0] - points[i][0]), max(points[i][1] - points[i-1][1], points[i-1][1] - points[i][1]))
	}
	return total
}