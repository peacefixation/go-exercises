package math

// Abs return the absolute value of an integer
func Abs(n int64) int64 {
	if n < 0 {
		return n * -1
	}
	return n
}

// Signum return the signum of a number (-1, 0, 1)
func Signum(n int) int {
	if n < 0 {
		return -1
	} else if n > 0 {
		return 1
	} else {
		return 0
	}
}
