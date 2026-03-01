package main

func countStudents(students []int, sandwiches []int) int {
	ate := 0;
	j := 0;
	for true {
		before := ate
		for i := 0; i < len(students); i++ {
			if students[i] == -1 {
				continue;
			}
			if students[i] == sandwiches[j] {
				j++;
				ate++;
				students[i] = -1
			}
		}
		if before == ate {
			return len(sandwiches) - ate
		}
	}
	return 0
}