func minPartitions(n string) int {
    rN := 0;
    for _, v := range n {
        rN = max(rN, int(v-'0'))
        if rN == 0 {
            return 9
        }
    }
    return rN
}