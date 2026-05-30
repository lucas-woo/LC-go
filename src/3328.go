package main

func maxOperations(s string) int {
	var pile int = 0
	var operations int = 0
	for i, v := range s {
		if v == '1' {
			pile++;
			continue;
			} 			
		if i == len(s) - 1 {
			if s[i] == '0' {
				operations += pile
			}
			break;
		}
		if s[i+1] == '1' {
			operations += pile
		}
	}
	return operations
}