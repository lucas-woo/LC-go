package main

func minDeletionSize(strs []string) int {
	total := 0
	for i := 0; i < len(strs[0]); i++ {
		var cur byte;
		for j := 0; j < len(strs); j++ {
			if strs[j][i] < cur {
				total++;
				break;
			}
			cur = strs[j][i]
		}
	}
	return total
}