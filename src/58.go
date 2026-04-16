package main
func lengthOfLastWord(s string) int {
  var last int;
	var space bool = true;
	for _, v := range s {
		if v == rune(' ') {
			if !space {
				space = true
			}
			continue;
		}
		if space {
			last = 0;
			space = false;
		}
		last ++;
	}
	return last
}