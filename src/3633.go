package main


func earliestFinishTime(lst, ld, wst, wd []int) int {
    minland := 1 << 31
    for i := range lst {
        minland = min(minland, lst[i] + ld[i])
    } 
    minwat := 1 << 31
    for i := range wst {
        minwat = min(minwat, wst[i] + wd[i])
    } 
    mindur := 1 << 31
    for i := range lst {
        mindur = min(mindur, max(minwat, lst[i]) + ld[i])
    }
    for i := range wst {
        mindur = min(mindur, max(minland, wst[i]) + wd[i])
    } 
    return mindur
}