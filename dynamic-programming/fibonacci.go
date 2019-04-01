// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence,
// such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
// F(0) = 0,  F(1) = 1
// F(N) = F(N - 1) + F(N - 2), for N > 1.

package dynamic

func Fibonacci(N int) int {
	var m = make(map[int]int)
	m[0], m[1], m[2] = 0, 1, 1
	return fibonacci(N, m)
}

func fibonacci(n int, m map[int]int) int {
	if n < 0 {
		return 0
	}
	if v, ok := m[n]; ok {
		return v
	}
	fib := fibonacci(n-1, m) + fibonacci(n-2, m)
	m[n] = fib
	return fib
}
