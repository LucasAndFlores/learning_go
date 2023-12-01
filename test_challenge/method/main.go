package main

import "fmt"

type student struct {
    name string 
    age int
    compiled string
}

func (s *student) DefineCompiled() {
    if s.name == "" {
        s.compiled = ""
        return
    }

    if s.age < 1 {
        s.compiled = ""
        return
    }

    s.compiled = fmt.Sprintf("My name is %v and I have %v years old", s.name, s.age)
}


func main () {
    s1 := student {
        name: "Lucas",
        age: 30,
    }

    s1.DefineCompiled()

    fmt.Println(s1.compiled == "")
}
