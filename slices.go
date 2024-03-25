package utilities

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

func MapSlice[X any, R any](array []X, mapFunction func(object X, index int) R) []R {
	newArray := make([]R, 0)

	for i, v := range array {
		newArray = append(newArray, mapFunction(v, i))
	}

	return newArray
}

func ReduceSlice[R any, X any](array []X, resultObject *R, reduceFunction func(totalObject R, currentObject X)) {

	for _, v := range array {
		reduceFunction(*resultObject, v)
	}
}

func FilterSlice[X any](array []X, filterFunction func(object X, index int) bool) []X {
	newArray := make([]X, 0)

	for i, val := range array {
		if filterFunction(val, i) {
			newArray = append(newArray, val)
		}
	}

	return newArray
}
