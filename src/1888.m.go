package main

func minFlips(s string) int {

    ones := 0;
    twos := 0;
    flippedOnes := 0;
    flippedTwos := 0;

    for i := 0; i < len(s); i++ {
        if int(s[i] - '0') == (i % 2) {
            ones++;
            flippedOnes++
        } else {
            flippedTwos++;
            twos++
        }
        flippedOnes = min(flippedOnes, twos)
        flippedTwos = min(flippedTwos, ones)
    }

    if len(s) % 2 == 0 {
        return min(ones,twos);
    } 
    return min(min(ones,twos), min (flippedOnes,flippedTwos))

}