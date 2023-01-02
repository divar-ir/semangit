package utils

import "github.com/sirupsen/logrus"

func GetResultOrPanic[TResult interface{}](result TResult, error error) TResult {
	if error != nil {
		logrus.Panic(error) // calls panic() after logging
	}
	return result
}

func PanicError(error error) {
	if error != nil {
		logrus.Panic(error) // calls panic() after logging
	}
}
