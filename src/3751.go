package main

func totalWaviness(num1 int, num2 int) int {
	total := 0
	for i := max(101, num1); i <= num2; i++ {
		n := i;
		right := n % 10
		n /= 10
		mid := n % 10;
		n /= 10
		for n > 0 {
			left := n % 10;
			n /= 10;
			if (mid < left && mid < right) || (mid > right && mid > left) {
				total++;
			}
			right = mid
			mid = left
		}
	}

	return total
}