package main

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
