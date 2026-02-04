package main

func findDisappearedNumbers(nums []int) []int {
  var numsMap map[int]bool = make(map[int]bool, len(nums))
	var returnArr []int;
	for _, v := range nums {
		numsMap[v] = true
	}
	for i := 1; i <= len(nums); i++ {
		_, ok := numsMap[i];
		if !ok {
			returnArr = append(returnArr, i)
		}
	}
	return returnArr
}