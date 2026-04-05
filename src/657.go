package main

func judgeCircle(moves string) bool {
	ve := 0;
	h := 0;
	for _, v := range moves {
		if v == 'L' {
			h--;
		} else if v == 'R' {
			h++
		} else if v == 'U' {
			ve++
		} else {
			ve--;
		}
	}
	return ve == 0 && h == 0
}