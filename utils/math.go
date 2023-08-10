package utils

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Clamp[T constraints.Ordered](v, low, high T) T {
	return Min(Max(v, low), high)
}

func ClampCycle[T constraints.Ordered](v, low, high T) T {
	if v > high {
		return low
	}
	if v < low {
		return high
	}
	return v
}
