package utilities

func FindInSlice[X any, T []X](slice T, condition func(object X) bool) *X {

	var returnObject X

	success := false

	for _, obj := range slice {
		if condition(obj) {
			returnObject = obj
			success = true
			break
		}
	}

	if success {
		return &returnObject
	} else {
		return nil
	}
}

func MapSlice[X any, R any](array []X, mapFunction func(object X) R) []R {
	newArray := make([]R, 0)

	for _, v := range array {
		newArray = append(newArray, mapFunction(v))
	}

	return newArray
}

func ReduceSlice[R any, X any](array []X, resultObject *R, reduceFunction func(totalObject R, currentObject X)) {

	for _, v := range array {
		reduceFunction(*resultObject, v)
	}
}
