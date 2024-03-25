package utilities

func FindInSlice[X any, T []X](slice T, condition func(object X) bool) (*X, *int) {

	var returnObject X
	var index int

	success := false

	for i, obj := range slice {
		if condition(obj) {
			returnObject = obj
			success = true
			index = i
			break
		}
	}

	if success {
		return &returnObject, &index
	} else {
		return nil, nil
	}
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
