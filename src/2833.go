package main


func furthestDistanceFromOrigin(moves string) int {
	d := 0
	something := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'R' {
			d++
		} else if moves[i] == 'L' {
			d--;
		} else {
			something++;
		}
	}
	if d < 0 {
		return (-d) + something
	}
	return d + something
}

