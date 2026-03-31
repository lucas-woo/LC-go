package main

func checkStrings(s1 string, s2 string) bool {
	evenS1 := make(map[byte]int, 0)
	evenS2 := make(map[byte]int, 0)
	oddS1 := make(map[byte]int, 0)
	oddS2 := make(map[byte]int, 0)

	for i := 0; i < len(s1); i++ {
		if i & 1 == 0 {
			if _, ok := evenS1[s1[i]]; !ok {
				evenS1[s1[i]] = 1
			} else {
				evenS1[s1[i]]++
			}
		} else {
			if _, ok := oddS1[s1[i]]; !ok {
				oddS1[s1[i]] = 1
			} else {
				oddS1[s1[i]]++
			}
		}
	}
	for i := 0; i < len(s2); i++ {
		if i & 1 == 0 {
			if _, ok := evenS2[s2[i]]; !ok {
				evenS2[s2[i]] = 1
			} else {
				evenS2[s2[i]]++
			}
		} else {
			if _, ok := oddS2[s2[i]]; !ok {
				oddS2[s2[i]] = 1
			} else {
				oddS2[s2[i]]++
			}
		}
	}

	for k, v := range evenS1 {
		 val, ok := evenS2[k];
		if !ok {
			return false
		}
		if val != v {
			return false
		}
	}
	for k, v := range oddS1 {
		 val, ok := oddS2[k];
		if !ok {
			return false
		}
		if val != v {
			return false
		}
	}	
	return true
}