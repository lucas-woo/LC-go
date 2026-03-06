package main

func checkOnesSegment(s string) bool {
    seenZero := false
    for i := 0; i < len(s); i++ {
        if s[i] == '1' && seenZero {
            return false
        } 
        if s[i] == '0' {
            seenZero = true
        }
    }
    return true
}