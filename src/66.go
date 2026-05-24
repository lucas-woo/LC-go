package main
// import "fmt"

func plusOne(digits []int) []int {
	var nine bool = true

	for i := len(digits) -1; i >= 0; i-- {
		if digits[i] == 9 && nine {
			digits[i] = 0;
		} else if nine {
			digits[i]++;
			nine = false
			break;
		}
	}
	if !nine {
		return digits
	} 
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}