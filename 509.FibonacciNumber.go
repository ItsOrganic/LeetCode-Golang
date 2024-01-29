// using dynamic programming approach 
func fib(n int) int {
    if n < 2{
    return n
    }
    if n == 2 {
        return n-1
    }
    a := make([]int, n)
    a[0] = 1
    a[1] = 1
    for i := 2; i < n; i++ {
        a[i] = a[i-1] + a[i-2]
    }
    return a[n-1]
}