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

//https://stackoverflow.com/questions/73243943/how-to-write-a-generic-max-function
func MultiMax[T constraints.Ordered](args ...T) T {
	if len(args) == 0 {
		return *new(T) // zero value of T
	}

	if isNan(args[0]) {
		return args[0]
	}

	max := args[0]
	for _, arg := range args[1:] {

		if isNan(arg) {
			return arg
		}

		if arg > max {
			max = arg
		}
	}
	return max
}

func isNan[T constraints.Ordered](arg T) bool {
	return arg != arg
}
