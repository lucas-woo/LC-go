package main

import "strconv"

func evalRPN(tokens []string) int {

	var stack []int = make([]int, len(tokens));
	var p int = -1;
	var operators map[string]bool = make(map[string]bool, 4);

	operators["*"] = true;
	operators["-"] = true;
	operators["+"] = true;
	operators["/"] = true;

	for _, v := range tokens {
		_, ok := operators[v];
		if !ok {
			temp, _ := strconv.Atoi(v)
			p++;
			stack[p] = temp 
			continue;
		}
		if v == "+" {
			v2 := stack[p];
			p--;
			v1 := stack[p];
			stack[p] = v1 + v2
		} else if v == "-" {
			v2 := stack[p];
			p--;
			v1 := stack[p];
			stack[p] = v1 - v2
		} else if v == "*" {
			v2 := stack[p];
			p--;
			v1 := stack[p];
			stack[p] = v1 * v2
		} else {
			v2 := stack[p];
			p--;
			v1 := stack[p];
			stack[p] = v1 / v2
		}
	}
	return stack[0]
}