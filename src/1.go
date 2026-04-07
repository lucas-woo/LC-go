package main

func twoSum(nums []int, target int) []int {
    var hashmap map[int]int = make(map[int]int);
		var mySlice [] int = []int{};
    var i int = 0;
    for i < len(nums) {

			var temp int = target - nums[i]
			v, exists := hashmap[temp]
			if exists {
				mySlice = append(mySlice, v, i);
				break;
			}

			_, ok := hashmap[nums[i]]
      if ok {
				i++;
				continue;
			}			
			hashmap[nums[i]] = i;
      i++;
    }
	return mySlice
}
