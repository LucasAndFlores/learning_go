package main

import (
	"errors"
	"testing"
)

func TestSucessDouble(T *testing.T) {
    expected := 36
    v, _ := double(18)
    if v != float64(expected) {
        T.Error("This should not happen, value: ", v)
    }
}

func TestFailureDouble (T *testing.T) {
    exp := errors.New("Can't execute the function with this value: -15")

    _, err := double(-15)
    if err.Error() != exp.Error() {
        T.Errorf("Wrong error! : %v", err)
    }
}
