package main

import "strings"

func findKthBit(n int, k int) byte {
	var sb strings.Builder

	sb.WriteByte('0')

	for sb.Len() < k {
		temp := reverseInvert(sb.String())
		sb.WriteByte('1')
		sb.WriteString(temp)
	}
	return sb.String()[k - 1]
}

func reverseInvert(str string) string {
	arr := []byte(str)
	i := 0;
	j := len(arr) - 1;
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++;
		j--;
	}
	for k := 0; k < len(arr); k++ {
		if arr[k] == '1' {
			arr[k] = '0'
		} else {
			arr[k] = '1'
		}
	}
	return string(arr)
}