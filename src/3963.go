package main

func createGrid(m int, n int) []string {
 
	arr := make([]string, m)
	t := ""
	for i := 0; i < n; i++ {
		t = t + "."
	}
	arr[0] = t

	for i := 1; i < len(arr); i++ {
		
		t := ""
		for i := 0; i < n - 1; i++ {
			t = t + "#"
		}
		t = t + "."
		arr[i] = t
	}
	return arr
}