package main
import "strings"
func addBinary(a string, b string) string {

	var stringBuilder strings.Builder;

	ai := len(a) -1;
	bi := len(b) -1;

	aCur := false
	bCur := false

	carryBit := false;

	for {
		if ai < 0 && bi < 0 {
			break
		}

		if ai >= 0 && a[ai] == '1'{
			aCur = true
		} 
		if bi >= 0 && b[bi] == '1' {
			bCur = true
		}

		if aCur && bCur {
			if carryBit {
				stringBuilder.WriteByte('1')				
			} else {
				carryBit = true
				stringBuilder.WriteByte('0')
			}
		} else if aCur || bCur {
			if carryBit {
				stringBuilder.WriteByte('0')
			} else {
				carryBit = false
				stringBuilder.WriteByte('1')
			}
		} else {
			if carryBit {
				carryBit = false
				stringBuilder.WriteByte('1')
			} else {
				stringBuilder.WriteByte('0')
			}				
		}
		aCur = false
		bCur = false
		ai--;
		bi--;
	}
	if carryBit {
		stringBuilder.WriteByte('1');
	}

	toRev := stringBuilder.String()
	runes := []rune(toRev)
	i := 0
	j := len(toRev) - 1;
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++;
		j--;
	}
	return string(runes)
}