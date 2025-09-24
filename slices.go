package utilities

import "strings"

func RemoveFromSlice[X any](slice []X, index int) []X {
	// Check if the index is valid
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Create a new slice by concatenating the elements before and after the index

	return append(slice[:index], slice[index+1:]...)
}

func FindInSlice[X any](slice []X, condition func(object X) bool) int {
	for i, value := range slice {
		if condition(value) {
			return i
		}
	}
	return -1
}

func Map[T any, R any](input []T, f func(v T, i int) R) []R {
	out := make([]R, len(input))
	for i, v := range input {
		out[i] = f(v, i)
	}
	return out
}

func Filter[T any](input []T, f func(v T, i int) bool) []T {
	out := make([]T, 0)
	for i, v := range input {
		if f(v, i) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T any, R any](input []T, init R, reducer func(R, T) R) R {
	acc := init
	for _, v := range input {
		acc = reducer(acc, v)
	}
	return acc
}

func WalkMap(m map[string]any, path []string, cb func(path []string, key string, value any)) {
	for k, v := range m {
		cb(path, k, v)

		// Recurse into nested maps only
		if nested, ok := v.(map[string]any); ok {
			WalkMap(nested, append(path, k), cb)
		}
	}
}

func RemoveLastOccurrence(s, substr string) string {
	idx := strings.LastIndex(s, substr)
	if idx == -1 {
		return s
	}

	return s[:idx] + s[idx+len(substr):]
}
