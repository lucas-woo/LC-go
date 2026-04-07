package main
import "sort";

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums) - 2; i++ {
		low := i + 1;
		high := len(nums) - 1;
		for low < high {
			curSum := nums[i] + nums[low] + nums[high]
			closestSum = closest(curSum, closestSum, target)
			if curSum < target {
				low++;
			} else {
				high--;
			}
		}
	}

	return closestSum
}


func closest(closestSum, currentSum, target int) int {
	var clSum, crSum int;
	if closestSum > target {
		clSum = closestSum - target
	} else {
		clSum = target - closestSum
	}
	if currentSum > target {
		crSum = currentSum - target
	} else {
		crSum = target - currentSum
	}
	if clSum < crSum {
		return closestSum
	} else {
		return currentSum
	}

}