package main

func minJumps(nums []int) int {
    if len(nums) <= 2 {
        return max(0, len(nums)-1)
    }
    maxN := 0
    for _, n := range nums {
        maxN = max(n, maxN)
    }
    minprimefactor := make([]int, maxN+1)
    minprimefactor[0], minprimefactor[1] = 0, 1
    for i := int64(2); i <= int64(maxN); i++ {
        if minprimefactor[i] == 0 {
            minprimefactor[i] = int(i)
        }
        for j := i*i; j <= int64(maxN); j += i {
            if minprimefactor[j] == 0 {
                minprimefactor[j] = minprimefactor[i]
            }
        }
    }
    isPrime := func(x int) bool {
        return x >= 2 && minprimefactor[x] == x
    }
    primeFactors := func(x int) []int {
        var res []int
        for x > 1 {
            p := minprimefactor[x]
            res = append(res, p)
            for x % p == 0 {
                x /= p
            }
        }
        return res
    }

    primeToPos := make(map[int][]int)
    for i, n := range nums {
        pfs := primeFactors(n)
        for _, p := range pfs {
            primeToPos[p] = append(primeToPos[p], i)
        }
    }
    q := [][2]int{[2]int{0, 0}}
    queued := make([]bool, len(nums))
    queued[0] = true
    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        pos, dist := curr[0], curr[1]
        if pos == len(nums)-1 {
            return dist
        }
        num := nums[pos]
        if isPrime(num) {
            primejumps := primeToPos[num]
            for i := len(primejumps)-1; i >= 0; i-- {
                nextpos := primejumps[i]
                if queued[nextpos] {
                    continue
                }
                queued[nextpos] = true
                q = append(q, [2]int{nextpos, dist+1})
            }
            delete(primeToPos, num)
        }
        if pos-1 >= 0 && !queued[pos-1] {
            queued[pos] = true
            q = append(q, [2]int{pos-1, dist+1})
        }
        if pos+1 <= len(nums)-1 && !queued[pos+1] {
            queued[pos+1] = true
            q = append(q, [2]int{pos+1, dist+1})
        }     
    }
    return -1
}