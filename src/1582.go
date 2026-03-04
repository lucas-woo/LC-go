package main

func numSpecial(mat [][]int) int {
    visitedRow := make([]bool, len(mat))
    visitedCol := make([]bool, len(mat[0]))
    count := 0
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            if visitedRow[i] {
                break;
            }
            if visitedCol[j] {
                continue;
            }
            if mat[i][j] == 0 {
                continue
            }

            onesCount := 0
            for k := 0; k < len(mat[0]); k++ {
                if mat[i][k] == 1 {
                    onesCount++;
                }
            }
            anotherCount := 0
            for k := 0; k < len(mat); k++ {
                if mat[k][j] == 1 {
                    anotherCount++;
                }
            }

            if (onesCount + anotherCount) == 2 {
                count++;
            }
            visitedRow[i] = true
            visitedCol[j] = true
        }
    }
    return count
}