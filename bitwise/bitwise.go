package bitwise

func BinPow(x int) int {
	return 1 << x
}

func Pow(x int) int {
	res := 2

	for i := 1; i < x; i++ {
		res *= 2
	}

	return res
}
