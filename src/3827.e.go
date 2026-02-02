package main

func countMonobit(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}

	totalCount := 2;

	twos := 2

	for (twos * 2) - 1 <= n {
		twos *= 2;
		totalCount++;
	}
	return totalCount
}