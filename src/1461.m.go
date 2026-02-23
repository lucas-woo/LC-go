package main

import "math"
func hasAllCodes(s string, k int) bool {
  if len(s) < k {
		return false
	}
	kMap := make(map[string]bool, 0)
	for i := 0; i <= len(s) - k; i++ {
		end := k + i
		if _, ok := kMap[s[i:end]]; !ok {
			kMap[s[i:end]] = true
		}
	}
	return len(kMap) == int(math.Pow(2, float64(k)))
}