package main

func nextGreatestLetter(letters []byte, target byte) byte {
	smallest := letters[0]
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return smallest
}