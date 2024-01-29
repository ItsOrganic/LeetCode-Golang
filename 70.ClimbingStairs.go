//simple dynamic programming logic
func climbStairs(n int) int {
    if n < 3{
        return n
    }
    a := make([]int, n)
    a[0] = 1
    a[1] = 2
    for i := 2; i < n; i++ {
        a[i] = a[i-1] + a[i-2]
    }
    return a[n-1]
}