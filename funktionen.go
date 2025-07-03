package main

func istGerade(n int) bool {
	return n%2 == 0
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fakultaet(n int) int {
	if n == 0 {
		return 1
	}
	return n * fakultaet(n-1)
}
