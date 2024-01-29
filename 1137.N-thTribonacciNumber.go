func tribonacci(n int) int {
    if n < 2 {
        return n
    }
    a := make([]int, n+1)
    a[0] = 0
    a[1] = 1
    a[2] = 1
    for i := 3; i <= n; i++ {
        a[i] = a[i-3] + a[i-2] + a[i-1]
    } 
    return a[n]
}