package utils

import "fmt"

func GetResultOrPanicError[TResult interface{}](result TResult, error error) TResult {
	if error != nil {
		panic(fmt.Errorf("Error getting result.\nError: %v\nCurrent result: %v", error, result))
	}
	return result
}

func PanicError(error error) {
	if error != nil {
		panic(error)
	}
}
