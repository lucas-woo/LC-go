package main

func numOfStrings(patterns []string, word string) int {
    
	cnt := 0
  for _, pat := range patterns {
    for i := 0; i < len(word); i++ {
        if pat[0] == word[i] && numOfStringsHelper(pat, word, i){
            cnt++
            break;
        }
    }
	}
	return cnt
}

func numOfStringsHelper(pat, word string, i int) bool {
  k := 0
  for j := i; j < len(word) && k < len(pat); j++ {
    if pat[k] != word[j] {
        return false
    }
    k++
    if k == len(pat) {
        return true
    }
	}
  return false
}