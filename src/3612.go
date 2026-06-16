package main

import (
	"strings"
)

func processStr(s string) string {

  arr := make([]byte, 1000000)

	last := 0;

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if last <= 0 {
				continue
			}
			last--;
		} else if s[i] == '%' {
			processStrReverse(arr, last - 1)
		} else if s[i] == '#' {
			processStrCopy(arr, last)
			last *= 2
		} else {
			arr[last] = s[i]
			last++
		}
	}

	var sb strings.Builder
	for i := 0; i < last; i++ {
		sb.WriteByte(arr[i])
	}
	return sb.String()
}

func processStrFindDup(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			c++
		}
	}
	return c
}

func processStrCopy(arr []byte, last int) {
	for i := 0; i < last; i++ {
		arr[last+i] = arr[i]
	}
}

func processStrReverse(arr []byte, last int) {
	i := 0;
	for i < last {
		arr[i],arr[last] = arr[last],arr[i]
		i++;
		last--;
	}
}