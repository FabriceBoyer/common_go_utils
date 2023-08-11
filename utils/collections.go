package utils

import (
	"math/rand"
	"time"
)

// order is not important
func Remove[T interface{}](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// order is important
func RemoveKeepOrder[T interface{}](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

func InsertAt[T any](a []T, index int, value T) []T {
	n := len(a)
	if index < 0 {
		index = (index%n + n) % n
	}
	switch {
	case index == n: // nil or empty slice or after last element
		return append(a, value)

	case index < n: // index < len(a)
		a = append(a[:index+1], a[index:]...)
		a[index] = value
		return a

	case index < cap(a): // index > len(a)
		a = a[:index+1]
		var zero T
		for i := n; i < index; i++ {
			a[i] = zero
		}
		a[index] = value
		return a

	default:
		b := make([]T, index+1) // malloc
		if n > 0 {
			copy(b, a)
		}
		b[index] = value
		return b
	}
}

func GetRandomSample(s []string, n int) []string {
	rand.Seed(time.Now().UnixNano())
	if n > len(s) {
		n = len(s) - 1
	}
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s[:n]
}
