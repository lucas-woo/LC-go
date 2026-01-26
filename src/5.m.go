package main
import "strings"
func longestPalindrome(s string) string {
  
	var longest string
	for i := 0; i < len(s); i++ {

		var sb strings.Builder

		left := i
		isEven := true
		lastSeen := -1
		for right := len(s) - 1; right >= left; right-- {
			if right == left {
				isEven = false
			}
			if s[right] == s[i] && lastSeen == -1 {
				lastSeen = right
			}			
			if s[left] == s[right] {
				sb.WriteByte(s[left])
				left++;
				continue
			}
			if lastSeen != -1 {
				right = lastSeen
				lastSeen = -1;
			}
			sb.Reset()
			left = i
			if s[left] == s[right] {
				sb.WriteByte(s[left])
				left++;
				continue
			}
		}

		toBeConverted := sb.String()
		var converted string
		if isEven {
			for j := len(toBeConverted) - 1; j >= 0; j-- {
				sb.WriteByte(toBeConverted[j])
			}
			converted = sb.String()
		} else {
			for j := len(toBeConverted) - 2; j >= 0; j-- {
				sb.WriteByte(toBeConverted[j])
			}
			converted = sb.String()			
		}
		if len(converted) > len(longest) {
			longest = converted
		}
	}
	return longest
}