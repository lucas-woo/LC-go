pakcage main

func centeredSubarrays(nums []int) int {
	var visited map[int]bool = make(map[int]bool, 0)
	rN := 0
  for i := 0; i < len(nums); i++ {
		total := 0
		for j := i; j < len(nums); j++ {
			total += nums[j] 
			visited[nums[j]] = true 
			if _, ok := visited[total]; ok {
				rN++;
			}			
		}
		visited = make(map[int]bool, 0)
	}
	return rN
}