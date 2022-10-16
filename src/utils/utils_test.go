package utils

import (
	"errors"
	"semangit/src/testUtils"
	"testing"
)

func TestPanicError(t *testing.T) {
	t.Run("Should panic if error is returned", func(t *testing.T) {
		returnError := func() error {
			return errors.New("error")
		}
		testUtils.AssertPanics(t, func() {
			PanicError(returnError())
		})
	})
	t.Run("Should not panic if error is nil", func(t *testing.T) {
		returnError := func() error {
			return nil
		}
		PanicError(returnError())
	})
}

func TestGetResultOrPanicError(t *testing.T) {
	t.Run("Should panic if error is returned", func(t *testing.T) {
		returnError := func() (interface{}, error) {
			return nil, errors.New("error")
		}
		testUtils.AssertPanics(t, func() {
			GetResultOrPanicError(returnError())
		})
	})
	t.Run("Should return result if no error occurs", func(t *testing.T) {
		returnResult := func() (string, error) {
			return "result", nil
		}
		GetResultOrPanicError(returnResult())
	})
}
