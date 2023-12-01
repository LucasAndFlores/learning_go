package main

import (
	"testing"
)

func TestStudentSuccess (t *testing.T) {
    stu1 := student{
        name: "Loki",
        age: 155,
    }

    stu1.DefineCompiled()

    if stu1.compiled == "" {
        t.Errorf("Something happen, error from test: %v", stu1)
    }
}

func TestStudentFailure (t *testing.T) {
    stuError := student{
        name: "heroi",
    } 

    stuError.DefineCompiled()

    if stuError.compiled != "" {
        t.Errorf("An error happened with the test, value: %v", stuError)
    }
}
