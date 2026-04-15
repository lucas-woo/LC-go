package main

func closestTarget(words []string, target string, startIndex int) int {
	found := false
	for i := 0; i < len(words); i++ {
		if words[i] == target {
			found = true
		}
	}
	if !found {
		return -1
	}
	i := startIndex
	smallest := -1
	for {
		if words[i%len(words)] == target {
			smallest = i - startIndex
			break
		}
		i++
	}
	i = startIndex
	for {
		if words[(i+len(words)) % len(words)] == target {
			smallest = min(startIndex - i, smallest)
			return smallest
		}
		i--
	}
	return smallest
}