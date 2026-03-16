package main

func minDays(bloomDay []int, m int, k int) int {
    left := 0;
    right := 0;
    for i := 0; i < len(bloomDay); i++ {
        right = max(right, bloomDay[i])
    }
    least := -1
    for left <= right {
        mid := (left + right) / 2
        ok := validate(bloomDay, m, k, mid)
        if ok {
            least = mid;
            right = mid - 1;
        } else {
            left = mid + 1;
        }
    }
    return least
}

func validate(bloomDay []int, m, k, day int) bool {
    count := 0;
    cur := 0
    for i := 0; i < len(bloomDay); i++ {
        if day >= bloomDay[i] {
            cur++;
        } else {
            cur = 0
        }
        if cur == k {
            cur = 0;
            count++;
        }
        if count >= m {
            return true
        }
    }
    return false
}