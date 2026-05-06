package main

func rotateTheBox(boxGrid [][]byte) [][]byte {
	box := make([][]byte, len(boxGrid[0]));
	for i := 0; i < len(box); i++ {
		box[i] = make([]byte, len(boxGrid))
	}

	for i := 0; i < len(box); i++ {
		for j := 0; j < len(box[0]); j++ {
			box[i][j] = '.'
		}
	}

	for i := 0; i < len(boxGrid); i++ {
		for j := 0; j < len(boxGrid[0]); j++ {
			box[j][len(boxGrid) - i - 1] = boxGrid[i][j]
		}
	}

	for col := 0; col < len(box[0]); col++ {


		for i := len(box) - 2; i >= 0; i-- {

			if box[i][col] != '#' {
				continue;
			}
			box[i][col] = '.'
			j := i;
			for j + 1 < len(box) && box[j + 1][col] == '.' {
				j++
			}
			box[j][col] = '#'
		}
	}

	return box
}