package main

func twoEditWords(queries []string, dictionary []string) []string {
	arr := make([]string, 0)

	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(dictionary); j++ {
			invalid := 0
			for k := 0; k < len(dictionary[j]); k++ {
				if invalid > 2 {
					break;
				}
				if queries[i][k] != dictionary[j][k] {
					invalid++
				}
			}
			if invalid <= 2 {
				arr = append(arr, queries[i])
				break;
			}
		}
	}
	return arr
}