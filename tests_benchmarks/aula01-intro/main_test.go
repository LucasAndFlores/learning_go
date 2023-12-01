package main

import ("testing")

// estrutura que armazena os dados que serão testados em tabela
type test struct {
    data []int
    answer int
}

func TestSoma(t *testing.T ) {
    x := soma(3,2,1)
    result := 6
    if x != result {
        t.Error("Expected: ", result, "Got: ", x)
    }
}

func TestSomaEmTabela(t *testing.T) {
    tests := []test {
        {data: []int{1,2,3}, answer: 6},
        {data: []int{10, 10,10}, answer: 30},
        {data: []int{20, 20, 20}, answer: 60},
    } // cria os casos

    for _, v := range tests {
        r := soma(v.data...) // testa cada um com a funcao

        if r != v.answer {
            t.Error("Expected: ", v.answer, "Got: ", r)
        }
    }
}

func BenchmarkSoma(b *testing.B) { //benchmark test
    tests := []test {
        {data: []int{1,2,3}, answer: 6},
        {data: []int{10, 10,10}, answer: 30},
        {data: []int{20, 20, 20}, answer: 60},
    } // cria os casos

    for i := 0; i < b.N ; i++ {

    for _, v := range tests {
        soma(v.data...) 
    }
    }
}
