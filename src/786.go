package main

import "strings"
func rotateString(s string, goal string) bool {
    var sb strings.Builder;

    for i := 0; i < len(s); i++ {
        sb.WriteByte(s[i])
        var sb2 strings.Builder
        for j := i + 1; j < len(s); j++ {
            sb2.WriteByte(s[j])
        }
        sb2.WriteString(sb.String())
        if sb2.String() == goal {
            return true
        }
    }
    return false
}