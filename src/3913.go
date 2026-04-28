package main

import "sort"
import "strings"
func sortVowels(s string) string {
    vowels := make(map[byte]int, 0)
    vowels['a'] = 0;
    vowels['e'] = 0
    vowels['i'] = 0
    vowels['o'] = 0
    vowels['u'] = 0
    vowelsOccur := make(map[byte]int, 0)
    vowelsOccur['a'] = -1;
    vowelsOccur['e'] = -1
    vowelsOccur['i'] = -1
    vowelsOccur['o'] = -1
    vowelsOccur['u'] = -1    
    for i := 0; i < len(s); i++ {
        
        if _, ok := vowels[s[i]]; ok {
            vowels[s[i]]++
            if vowelsOccur[s[i]] == -1 {

                vowelsOccur[s[i]] = i
            }
        }
    }
    arr := make([][]int, 0)
    
    for k, v := range vowels {
        if v <= 0 {
            continue;
        }
        arr = append(arr, []int{int(k),v, vowelsOccur[k]})
    }
    sort.Slice(arr, func(i,j int)bool {
        if arr[i][1] == arr[j][1] {
            return arr[i][2] < arr[j][2]
        }
        return arr[i][1] > arr[j][1]
    })

    var sb strings.Builder;
    ptr := 0
    for i := 0; i < len(s); i++ {
        if _, ok := vowels[s[i]]; ok {
            if arr[ptr][1] <= 0 {
                ptr++
            }
            sb.WriteByte(byte(arr[ptr][0]))
            arr[ptr][1]--;
        } else {
            sb.WriteByte(s[i])
        }
    }
    return sb.String()
}