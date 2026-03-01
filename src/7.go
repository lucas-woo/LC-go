package main

func reverse(x int) int {

	reversedValue := 0
	powTen := 0;
	isNegative := false

	overflow := false
	overflowed := false
	if x < 0 {
		isNegative = true;
		x *= -1;
	}

	for x > 0 {

		curVal := x % 10;

		reversedValue *= 10;
		reversedValue += curVal;

		if !overflowed {
		if powTen == 0 {
			if curVal == 2 {
				overflow = true
			}
			if curVal > 2 {
				overflowed = true
			}
		} 
		if powTen == 1 && overflow {
			if curVal < 1 {
				overflow = false
			} 
			if curVal > 1 {
				overflowed = true
			}			
		}
		if powTen == 2 && overflow {
			if curVal < 4 {
				overflow = false
			}
			if curVal > 4 {
				overflowed = true
			}						
		}
		if powTen == 3 && overflow {
			if curVal < 7 {
				overflow = false
			}
			if curVal > 7 {
				overflowed = true
			}						
		}
		if powTen == 4 && overflow {
			if curVal < 4 {
				overflow = false
			}
			if curVal > 4 {
				overflowed = true
			}						
		}
		if powTen == 5 && overflow {
			if curVal < 8 {
				overflow = false
			}
			if curVal > 8 {
				overflowed = true
			}						
		}
		if powTen == 6 && overflow {
			if curVal < 3 {
				overflow = false
			}
			if curVal > 3 {
				overflowed = true
			}						
		}
		if powTen == 7 && overflow {
			if curVal < 6 {
				overflow = false
			}
			if curVal > 6 {
				overflowed = true
			}						
		}
		if powTen == 8 && overflow {
			if curVal < 4 {
				overflow = false
			}
			if curVal > 4 {
				overflowed = true
			}						
		}
		if powTen == 9 && overflow && isNegative {
			if curVal <= 8 {
				overflow = false
			}
		}												
		if powTen == 9 && overflow && !isNegative {
			if curVal <= 7 {
				overflow = false
			}
		}			
		}
												
		powTen++;
		x /= 10
	}

	if powTen > 10 {
		return 0
	}
	if powTen == 10 && (overflowed || overflow) {
		return 0;
	}
	if isNegative {
		return -reversedValue
	}
	return reversedValue
}

// -2,147,483,648 to 2,147,483,647 