package main
func rangeAddQueries(n int, queries [][]int) [][]int {
	var rA [][]int = make([][]int, n)
	for i := range rA {
		rA[i] = make([]int, n)
	}
	for _, v := range queries {
		row1 := v[0];
		col1 := v[1];
		row2 := v[2];
		col2 := v[3];
		for i := row1; i <= row2; i++ {
			for j := col1; j <= col2; j++ {
				rA[i][j] += 1;
			}
		}
	}
  return rA;
}

//row col row col