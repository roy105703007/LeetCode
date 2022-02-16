func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    gi, si := 0, 0
    count := 0
    for gi < len(g) && si < len(s){
        if g[gi] <= s[si] {
            gi++
            si++
            count++
        } else {
            si++
        }
    }
    return count
}
