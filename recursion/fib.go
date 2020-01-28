package recursion

func Fib(n int) int {
	if n < 3 {
		return 1
	}

	var (
		r int

		n1, n2 = 1, 1
	)

	for i := 2; i < n; i++ {
		r = n1 + n2
		n2 = n1
		n1 = r
	}

	return r
}

func FibRec(n int) int {
	if n < 3 {
		return 1
	}

	return FibRec(n-2) + FibRec(n-1)
}
