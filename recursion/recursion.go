package main

import "fmt"

func main() {
	cache := make(map[int64]int64)
	mExponent := memoized(exponent, cache)
	fmt.Println(mExponent(2, 32))
	fmt.Println(mExponent(2, 16))
}

func exponent(x, n int64) int64 {
	if n == 0 {
		return 1
	} else if n%2 == 1 {
		return x * exponent(x*x, n/2)
	} else {
		return exponent(x*x, n/2)
	}
}

var memoized = func(fn func(x, n int64) int64, cache map[int64]int64) func(x, n int64) int64 {
	return func(x, n int64) int64 {
		if val, ok := cache[n]; ok {
			return val
		}

		cache[n] = fn(x, n)
		return cache[n]
	}
}
