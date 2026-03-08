package main


import "strconv"
func findDifferentBinaryString(nums []string) string {
	exists := make(map[int64]bool)
	for _, v := range nums {
		num, _ := strconv.ParseInt(v, 2, 64)
		exists[num] = true 
	}
	found := -1
	for i := 0; i < len(nums) + 1; i++ {
		if _, ok := exists[int64(i)]; !ok {
			found = i
			break;
		}
	}
	str := strconv.FormatInt(int64(found), 2)
	for len(str) < len(nums[0]) {
		str = "0" + str
	}
	return str
}