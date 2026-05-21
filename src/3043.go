package main

func longestCommonPrefix(arr1 []int, arr2 []int) int {

		m1 := make(map[int]map[int]bool, 0);
		m2 := make(map[int]map[int]bool, 0);

		for i := 0; i < 9; i++ {
			m1[i] = make(map[int]bool, 0)
			m2[i] = make(map[int]bool, 0)
		}

		for i := 0; i < len(arr1); i++ {
			for arr1[i] > 0 {
				if arr1[i] >= 100000000 {
					m1[8][arr1[i]] = true
					arr1[i] /= 10;
				} else if arr1[i] >= 10000000 {
					m1[7][arr1[i]] = true
					arr1[i] /= 10;				
				} else if arr1[i] >= 1000000 {
					m1[6][arr1[i]] = true
					arr1[i] /= 10;
				} else if arr1[i] >= 100000 {
					m1[5][arr1[i]] = true
					arr1[i] /= 10;
				} else if arr1[i] >= 10000 {
					m1[4][arr1[i]] = true
					arr1[i] /= 10;
				} else if arr1[i] >= 1000 {
					m1[3][arr1[i]] = true
					arr1[i] /= 10;
				} else if arr1[i] >= 100 {
					m1[2][arr1[i]] = true
					arr1[i] /= 10;
				} else if arr1[i] >= 10 {
					m1[1][arr1[i]] = true
					arr1[i] /= 10;
				} else {
					m1[0][arr1[i]] = true
					arr1[i] /= 10;
				}
			}
		}

		for i := 0; i < len(arr2); i++ {
			for arr2[i] > 0 {
				if arr2[i] >= 100000000 {
					m2[8][arr2[i]] = true
					arr2[i] /= 10;
				} else if arr2[i] >= 10000000 {
					m2[7][arr2[i]] = true
					arr2[i] /= 10;				
				} else if arr2[i] >= 1000000 {
					m2[6][arr2[i]] = true
					arr2[i] /= 10;
				} else if arr2[i] >= 100000 {
					m2[5][arr2[i]] = true
					arr2[i] /= 10;
				} else if arr2[i] >= 10000 {
					m2[4][arr2[i]] = true
					arr2[i] /= 10;
				} else if arr2[i] >= 1000 {
					m2[3][arr2[i]] = true
					arr2[i] /= 10;
				} else if arr2[i] >= 100 {
					m2[2][arr2[i]] = true
					arr2[i] /= 10;
				} else if arr2[i] >= 10 {
					m2[1][arr2[i]] = true
					arr2[i] /= 10;
				} else {
					m2[0][arr2[i]] = true
					arr2[i] /= 10;
				}
			}			
		}

		for i := 8;	i >= 0; i-- {

			for v := range m1[i] {
				if _, ok := m2[i][v]; ok {
					return i + 1
				}
			}

		}
		return 0
}