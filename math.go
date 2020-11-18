package main

// Min returns the smaller integer
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger integer
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
