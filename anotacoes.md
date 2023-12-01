# Learning go - Aprenda Go Playlist
Playlist YT: https://www.youtube.com/watch?v=onK_nd--g1g&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=31

# Importante!!!!
Sempre rodar `go mod init github.com/youruser/yourrepo` antes de iniciar o projeto.
https://stackoverflow.com/questions/21001387/how-do-i-set-the-gopath-environment-variable-on-ubuntu-what-file-must-i-edit/53026674#53026674

## Função main 
A primeira linha da função main contem o escopo principal, é onde o interpretador vai começar. A primeira linha é onde vai começar e a ultima é onde vai encerrar. Porta de entrada pra qualquer programa em Go.
Notações: Sempre package ponto indentificador
Funções com tres pontos indicam uma funçao variadica, trabalha com qualquer numero de argumentos e consegue lidar com eles
Todos os tipos implementam uma interface vazia
Variaveis declaradas e nao utilizada, Go nao permite. Go vai retornar todas as variaveis declaradas no retorno, e caso voce nao queira utilizar um valor retornado, voce pode usar um _ para descartar a variavel
Variaveis é um objeto (uma posiçao na memoria) capaz de reter e representar valor ou expressao

## Variaveis, declaracoes e outros conceitos

#### Operador curto de declaraçao (:=)

Declara uma variavel. Tem tipagem automatica. O operador = é um operador de atribuição, nao de declaracao. Ela pode atribuir um valor a uma variavel já declarada, mas não pode ser usado sozinho e nem utilizar o := pra sobrescrever uma outra variavel. 
O declarador curto pode ser usado pra atribuir novos valores, mas apenas se voce delcarar pelo menos uma nova variavel em conjunto. 
Exemplo: 
```go
y := "lucas"
fmt.Printf("y: %v, %T \n", y, y)
y, z := "bob", "bom dia"
fmt.Printf("y: %v, %T \n", y, y)
fmt.Printf("z: %v, %T \n", z, z) 
```
O operador curto so vai funcionar no escopo de uma funçao ou code block, não é possivel fora dele. Para usar uma declaracao a nivel de package level scope (escopo a nivel global dentro do arquivo), voce pode usar o operador `var`.

Palavras reservadas em go: https://go.dev/ref/spec#Keywords
Pode atribuir expressões a variaveis tambem.

#### Palavra chave `var`

Variaveis que usam `var`, tem abragencia package level scope e precisa ser fora do code block. Voce pode inicializar a variavel e nao atribuir o valor, mas com isso, o valor so pode ser atribuido em um code block, nao em um package level scope. 

#### Tipos: 
São estaticos, nao podem ser alterados, pra sempre! Impossivel mudar o tipo depois de declarado. Isso funciona tanto com o operador curto quanto com var. Pra conferir tipos e printar, pode usar o `fmt.Printf`. Tipos primitivos: string, int, bool. Tipos compostos: slice, array, struct, map.

#### Valor zero:
Declaracao, inicializacao e atribuicao. Valor zero se encontrado em uma variavel antes dela ser utilizada pelo usuario. Usar sempre o operador curto, ou var em caso de package-level scope.
Os zeros de cada tipo:
    - ints: 0
    - floats: 0.0
    - booleans: false
    - strings: ""
    - pointers, functions, interfaces, slices, channels, maps: nil

#### Pacote fmt (funt)
String interpretada tem alguns operadores que podem ser interpretada automaticamente, como "\n", "\t". Se voce usar o caracter `` `` ele sera interpretado como raw literal.
 - Grupo #1: Print → standard out
        - func Print(a ...interface{}) (n int, err error) // nao adiciona linha
        - func Println(a ...interface{}) (n int, err error) // adiciona uma linha abaixo
        - func Printf(format string, a ...interface{}) (n int, err error) // pode adicionar a variavel e printar a variavel e o tipo dela
            - Format verbs. (%v %T)

 - Grupo #2: Print → string, pode ser usado como variável
        - func Sprint(a ...interface{}) string // vai concatenar os valores em uma unica string sem espaço
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string

 - Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

#### Criando o proprio tipo
Voce pode criar seu tipo e atribuir a uma `var`, como por exemplo: 
```go
type hotdog int

var b hotdog
```
Caso voce tenha uma variavel que seja do mesmo tipo, como o exemplo acima, e tente atribuir, vai dar erro. Por exemplo: 

```go
type hotdog int

var b hotdog = 10

func main() {
	x := 40
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", x)

	b = x // erro!
}
```

#### Conversão de tipos

Voce pode converter tipos (nao é coerção como JS ou TS). Exemplo 
```go
type hotdog int

var b hotdog = 101

func main() {
	x := 40
	fmt.Printf("%v\n", x)

	x = int(b)

	fmt.Printf("%v\n", x)
}
```

Na parte de cima, ocorreu um erro. Mas para não dar erro, voce pode fazer o seguinte: 
```go
type hotdog int

var b hotdog = 10

func main() {
	x := 40
	fmt.Printf("%v\n", x)

	b = hotdog(x)

	fmt.Printf("%v\n", b)
}
```

## Tipos e outras coisas mais

#### Como os computadores funcionam
Esquema de codificaçao, 2n. Exemplo das lampadas: 1 lampada = 2 mensagens. Valores binarios condificam mensagens
ASCII é sistema de codificacao de uma forma que os computadores podem gravar letras, simbolos e etc. UTF-8 é um sistema super set de ASCII.

#### Tipos numéricos
ints vs float: Inteiros e floats sao numeros com fraçao.
uint vs int: uint é unsigned e o int pode conter o sinal (-)
Todos os tipos numericos sao distintos exceto byte e rune. Byte = uint8 e rune=int32.
Como regra geral, sempre usar int e deixar o computador escolher qual é mais eficiente pra voce
Como regra geral pra float, sempre usar float64.
Tipo int e float64 são tipos numericos, mas não sao compativeis

#### Overflow
Um uint16, por exemplo, vai de 0 a 65535.
Que acontece se a gente tentar usar 65536? Voce vai ter um cenario de overflow da linguagem
Ou se a gente estiver em 65535 e tentar adicionar mais 1? Ele vai zerar e começar de 0 em diante

#### Tipo string
Strings são sequencias de bytes e imutaveis. Uma string é um "slice of bytes". https://go.dev/play/p/dt2x1ies5b
Caso voce queira uma formatacao estranha, voce pode usar o sinal de string literal do JS (backtick) pra manter a formatacao e isso é uma raw string literals.
Jeitos de ver caracteres e bytes que estão presentes em uma string: https://go.dev/play/p/TU8k7r5rG4L
Fazer um range em uma string, te da caracter por caracter, nao byte por bytes. Se voce usar o for loop, ele vai dar por bytes.

#### Const
São valores imutaveis, pode ser tipadas ou nao. Só terao valores atribuidos quando forem usadas.
Quando usada const, ela só tem o tipo atribuido no uso, já variaveis com var tem um tipo definido quando definida.
Exemplo: https://go.dev/play/p/DEjDGb2fcaj
Formas de declarar: 
const x = y
const ( x = y, z = true, w = "ola" )

#### Iota
São numeros interios que necessariamente ainda nao estao tipados ainda. Serve para situacoes em que a constante nao seja igual uma a outra, so quer que ela seja diferente das outras, e ela se repeta na declaracao. Da pra descartar valores usando _ . Voce nao precisa declarar varias iota's, mas pode declarar apenas uma vez e o valor vai se repetindo, exemplo: 
```go
const (
	x = iota
	y
	z
)

func main() {
	fmt.Println(x, y, z)
}

// print 0 1 2
```

#### Deslocamento de bits
Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita. Operador de deslocamento de bits é <<. Exemplo: 
```go
func main() {
	x := 1
	z := x << 1
	fmt.Printf("%b \n", x)
	fmt.Printf("%b \n", z)
}

Print 
1 
10 
```

## Fluxo de controle

### For loop padrão 
```go
	for x := 0; x <= 10; x++ {
		fmt.Println(x)
	}

```
Não existe while !!!
Pra aprender mais: https://gobyexample.com/

### Nested loop
Cada ciclo do loop interno vai ocorrer pra uma interacao do loop externa. Exemplo do relogio!
```go
func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Horas ", horas)

		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print(" ", minutos)
		}

		fmt.Println("")
	}
}
```

### Outros usos for
Um for em go sempre vai seguir a regra de ter a inicializacao, a condicao e o pos. Para tentar simular um while (ou seja, enquanto isso for verdadeiro, faca isso), precisamos separar os elementos e interar dentro do for. Por exemplo: 
```go
func main() {
	i := 0

	for i <= 10 {
		fmt.Printf("Indice menor que 10: %v \n", i)
		i++
	}
}

```

Um loop pode rodar "pra sempre" em go, apenas retirar a interacao e pode ate retirar a condicao. Exemplo: 

```go
func main() {
	i := 0

	for i <= 10 {
		fmt.Printf("Roda eternamente")
	}


	for {
		fmt.Printf("Roda eternamente")
	}
}

```
Voce pode ter loops "eternos", mas que contenham uma condicao e que possam usar break, e loop vai ser pausado internamente.
Exemplo: 
```go
func main() {
	i := 0

	for {
		if i < 10 {
			fmt.Printf("Indice menor que 10: %v \n", i)
			i++
		} else {
			fmt.Println("chega, cansei!")
			break
		}

	}
}
```

Tipos de for pela doc: https://go.dev/ref/spec#For_statements

Continue quebra aquela unica condicao e continua e o break quebra o loop como um todo.

### If, else if e else

O ! é um operador de negaçao. Ele inverte o resultado de comparacao (assim como no JS).

### Switch
O Foco é ter casos diferentes.
Exemplo: 
```go
func main() {
	x := 2
	switch {
	case x < 5:
		fmt.Println("Menor que 5")
	case x == 5:
		fmt.Println("igual 5")
	case x > 5:
		fmt.Println("Maior que 5")
	}
}
```

É possivel ter true ou uma variavel atribuida na frente do switch pra ele verificar de acordo com a variavel.

O fall-through por padrao, caso ele esteja certo, executa a mesma linha e assume que a proxima condicao tambem está correta. Exemplo: 

```go
	quemTaNoEscritorioHoje := "Marquinhos"
	switch quemTaNoEscritorioHoje {
	case "Marquinhos":
		fmt.Println("Marquinhos")
		fallthrough
	case "Maria":
		fmt.Println("Maria")
	case "Bob":
		fmt.Println("Bob")
	}
    // aqui ele vai print marquinhos (que atende a condicao) e tambem Maria
```
Voce pode ter varias expressoes juntas, separadas por virgula, por exemplo, ou até mesmo expressoes como variavel == ou !==.

Pode usar o switch para verificar o tipo da variavel, se necessario. Exemplo: 
```go
var x interface{} // tipo generico, mesmo usado no fmt Println

func main() {
	x = 45

	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float")
	default:
		fmt.Println("existe isso?")
	}
}
```

### Agrupamento de dados

### Arrays
Arrays, slices, maps e structs. Array nao é usado no dia a dia, mas o slice é usado no dia, só que construido em cima dos arrays.

Arrays tem tamanhos fixos. Len é o usado pra pegar o tamanho, usando `len(array)`. Mesmo que nao tenha dados la dentro, o tamanho dele sempre é fixo pelo que voce definiu anteriormente.

Pra declarar arrays, a sintaxe é 
```go
x := [5]int{1, 2, 3, 4, 5}
```

### Slices - Literal composta
É um tipo de dado composto. Array e slices são bem similares, mas o que diferencia é que um array voce precisa declarar o tamanho especifico dele, já o slice nao. Exemplo: 
```go
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
```

Os slices são mutaveis e acessados pelo indices para fazer modificaçoes, mas parece ter alguns metodos nativos tambem.

Caso voce tente atribuir um element em um indice que ainda nao existe no array, vai dar erro, exemplo: 

```go
    x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice[11] = 45 // panic: runtime error: index out of range [11] with length 5

    Outro Exemplo

    slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice))
	slice[5] = 45
```
Ou seja, voce só pode atribuir num indice que ja existe.
Slice sempre usam arrays subjacentes debaixo do pano, portanto é criado com um numero fixo de elementos, 

### Slice - for range
Com o range, é possivel percorrer o slice e ver o que tem la. Exemplo: 
```go
	slice := []string{"lucas", "larine", "bob", "nana", "milica"}

	for index, valor := range slice {
		fmt.Println("o valor do indice:", index, "é:", valor)
	}
```

Slice sempre usam arrays subjacentes debaixo do pano, portanto é criado com um numero fixo de elementos e voce nao consegue adicionar elementos por indice. Mas com a funcao append, é possivel. Exemplo: 
```go
slice := []string{"lucas", "larine", "bob", "nana", "milica"}

	for index, valor := range slice {
		fmt.Println("o valor do indice:", index, "é:", valor)
	}

	slice = append(slice, "Saudades chicao")

	for index, valor := range slice {
		fmt.Println("o valor do indice depos do append é:", index, "é:", valor)
	}
```

é possivel ignorar o indice ou elementos usando o operador `_`.

### Slice - Fatiando ou deletando
Quando estamos falando de slice um slice (cortar uma parte da estrutura), voce vai marcar a onde começar e onde voce deseja parar, mas parar nao significa que voce vai pegar esse elemento. 
Exemplo: 
```go
	sabores := []string{"peperoni", "mussarela", "frango", "quatro queijos", "formagio"}

	fatia := sabores[0:2]

	fmt.Println(fatia) // vai printar peperoni e mussarela
```

Caso voce queira pegar de um determinado numero até o final, voce pode usar o `len()` pra auxiliar. 

```go
	sabores := []string{"peperoni", "mussarela", "frango", "quatro queijos", "formagio"}

	fatia := sabores[0:len(sabores)]
    // uma alternativa, caso voce vai começar do indice 0, é deixar em branco, por exemplo:

	fatia := sabores[:len(sabores)]

	fmt.Println(fatia) // vai printar peperoni e mussarela
```

Caso voce queira excluir um elmento do slice, voce vai tem que selecionar do começo ate o item, e do item até o final. Exemplo: 
Por exemplo, se a gente gostaria de excluir o sabor abacaxi: 
```go
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	sabores = append(sabores[:2], sabores[3:]...)

	fmt.Println(sabores)
```

### Slice - append
Anexa itens em um slice. Voce sempre vai tem que reatribuir um valor a variavel e o append vai ter dois argumentos, o primeiro é o slice que voce quer adicionar os elementos, e o segundo é o que voce quer adicionar. Exemplo: 
```go
	umaslice := []int{1, 2, 3}
	//outraslice := []int{10, 20, 30}

	umaslice = append(umaslice, 4, 5, 6)
```

caso voce queira meergear dois arrays, voce pode utilizar o operador unfurl (rest, no caso de JS), enumeration na documentação oficial, que são tres pontos depois do slice. Exemplo:

```go
    umaslice := []int{1, 2, 3}
	outraslice := []int{10, 20, 30}

	umaslice = append(umaslice, outraslice...)
```

### Slice - Make 

O make é uma forma de criar slices (que são arrays que não tem um tamanho fixo), mas ele coloca um terceiro parametro, que seria o `cap`. O cap é util pelo que ter apenas um slice sem cap (que vem de capacity), ajuda a otimizar o espaço desse array ocupado na memoria. Ele vai ter 3 parametros, o tipo do slice, o length inicial e o cap. Voce pode adicionar elementos até atingir o length, mas caso voce adicione um elemento que ultrapasse o length ou o cap (que vem de capacity), voce necessariamente precisa usar o `append`, porque ele vai criar um novo array (como o anterior) e vai reorganizar o length e a capacity. 

Esse tipo de abordagem é bem utilizado quando voce precisa otimizar os recursos ao maximo, como memoria e espaço utilizado. 

Go playground: https://go.dev/play/p/GOfinrXmXrF

### Maps - Intro
Maps são estruturas de chave e valor. Não são ordenados. Exemplo: 
```go
	amigosBr := map[string]int{
		"bruno": 456456,
		"nano":  568844,
	}
	fmt.Println(amigosBr)
	fmt.Println(amigosBr["bruno"]) // acessando chave e valor

	amigosBr["te amo JS"] = 11111 // assim voce adiciona valores!

	fmt.Println(amigosBr["te amo JS"])
```

Quando criaram a linguagem Go, eles criaram uma sintaxe "comma ok" para muitas coisas. No caso de maps, voce pode usar essa sintaxe, para verificar se uma chave existe no map. Exemplo: 


```go
	amigosBr := map[string]int{
		"bruno": 456456,
		"nano":  568844,
	}

    tala, existe := amigosBr["fantasma"]

	fmt.Println(tala, existe) // output: 0 false
```

### Maps - Range e deletando valores
Range é uma forma de acessar os maps, chave e valor. Exemplo: 
```go
	qualquercoisa := map[int]string{
		123: "muito massa",
		456: "ta na hora de molhar o biscoito",
		987: "POWPOWPOW",
		852: "yeah man",
	}

	for key, value := range qualquercoisa {
		fmt.Println(key, value)

	}
```

Para deletar, voce vai usar a palavra `delete(map, chave)`.

## Structs 
Structs são tipos que podem conter estruturas mais "complexas" e diferentes tipos dentro dele. Voce pode declarar cada campo de uma maneira explicita ou de forma mais simples. Exemplo: 

```go
import "fmt"

type cliente_restaurante struct {
	nome    string
	fumante bool
	gasto   float64
}

var testing []cliente_restaurante

func main() {
	cliente1 := cliente_restaurante{
		nome:    "Xuxinha",
		fumante: false,
		gasto:   145.58,
	}

	testing = append(testing, cliente1)

	cliente2 := cliente_restaurante{
		"Zezinho",
		true,
		250.56,
	}

	testing = append(testing, cliente2)

	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			testing = append(testing, cliente_restaurante{
				nome:    "lucas",
				fumante: true,
				gasto:   0.45,
			})
		} else {
			testing = append(testing, cliente_restaurante{
				nome:    "fabiana",
				fumante: false,
				gasto:   25.45,
			})
		}

	}

	fmt.Println(testing, len(testing))
}

```

### Structs embutidos
Structs podem ser embutidos (um dentro do outro). Exemplo:
```go
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario float64
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "xuxinha",
			idade: 35,
		},
		titulo:  "pizzaiolo",
		salario: 455.89,
	}

	fmt.Println(pessoa1, pessoa2.salario)
}

```
Voce pode acessar todo os campos usando `pessoa2.titulo`, e quando voce tem structs embutidos, se os campos usados no structs interno nao ter conflito com o externo, voce tambem tem a possibilidade de acessar diretamente. Exemplo: 

```go
package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario float64
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "xuxinha",
			idade: 35,
		},
		titulo:  "pizzaiolo",
		salario: 455.89,
	}

    // pesso2.pessoa.nome = pessoa2.nome, funciona!

	fmt.Println(pessoa1, pessoa2.salario)
}

```

### Structs anonimos 
Forma de se declarar struct anonimo: 
```go
	a := struct {
		nome  string
		idade int
	}{
		nome:  "mario",
		idade: 45,
	}
	fmt.Println(a)
```

## Funcoes
Estrutura padrao de funcao:

```go
func (receiver) identifier(parameters) (returns) {code}

```
Tudo em go é 'pass-by-value'. Basicamente, é como se a gente tivesse passando sempre o result final, e nao necessariamente apontando pra uma variavel.

Funcao basica e multiplos retornos (quase que retorno um objeto?): 
```go
//retorno simples
func soma(x int, y int) int {
	return x + y
}

// retorno multiplos e parametro variadico: 
// o operador variadico vai aceitar varios ints e dentro da funcao vai ser ttratado como um slice
func soma(x ...int) (int, int) {
	soma := 0
	for _, value := range x {
		soma += value
	}

	return soma, len(x)
}
```
Caso voce tenha um parametro diferente do variadico, o variadico precisa sempre ser o ultimo

### Desenrolar (enumerar) uma slice 
Caso voce tenha um array e vai usar uma funcao com o parametro variadico, existe uma forma de voce passar esse array usando o proprio operador variadico. Exemplo: 
```go
func main() {
	si := []int{10, 10, 40, 40}
	total := soma(si...)
	fmt.Println(total)
}

func soma(x ...int) int {
	soma := 0
	for _, value := range x {
		soma += value
	}

	return soma
}
```

Quando se usa um parametro variadico, o valor tambem pode ser zero ou infinito.

### Defer
Defer é uma expressao que pode ser usada no codigo para executar uma instrução por ultimo. Por exemplo: 
```go
func main() {
	defer fmt.Println("vem primeiro")
	fmt.Println("vem depois")

}

// output: 
vem depois
vem primeiro
```
Caso voce use dois deffers, a ordem será quase que um stack. Exemplo:
```go
func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

}

output: 
3
4
2
1
```
Caso voce tenha uma coisa que queira fazer e precise de algo depois, o defer pode ser usado. Pra fazer cleanup.

### Metodos
Metodo é uma funcao anexada a um tipo! Voce pode anexar funcoes (que sao metodos) de um tipo, e ela pode ser apenas chamada naquele tipo em especifico. Para anexar essa funcao ao tipo, voce vai usar o receiver e referenciaro type. Exemplo: 
```go
type pessoa struct {
	nome  string
	idade int
}

// func (receiver) identifier (parameters)(returns) {code}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia a todos")
}

func main() {
	bambino := pessoa{"Bambino", 8}
	bambino.oibomdia()
	fmt.Println("Hello, 世界")
}
```

### Interfaces e poliformismo
Em go valores, podem ter mais que um tipo
A interface permite que um valor tenha mais que um tipo
Declaração: keyword identifier type → type x interface
Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
Esse tipo será o seu tipo e também o tipo da interface.

Em go, voce pode ter interfaces, que vao encapsular metodos iguais, mas que vao ser implementados em diferentes structs.
Caso voce queira ter acesso a cada um dos metodos diferentes, voce consegue ter esse poliformismo atraves de uma funcao.
Exemplo feito na aula: https://go.dev/play/p/_17aJnzFeQg
https://gobyexample.com/interfaces

### Funcoes anonima
Pra fazer funcoes anonimas, elas precisam ser auto executadas
Usadas na go routines ou quando precisa so uma vez
Exemplo:
x := 15

func(x int) {
fmt.Println(x, "vezes 10 e:")
fmt.Println(x * 10)
}(x)


### Funcoes como expressao
E possivel atribuir uma funcao a uma variavel e usar essa funcao chamando pela variavel (assim como JS), por exemplo:
x := 15

multiply := func(x int) {
fmt.Println(x, "vezes 10 e:")
fmt.Println(x * 10)
}

multiply(x)

### Retornar uma funcao
Funcoes podem retornar funcoes, por exemplo:
https://go.dev/play/p/6cg8xZbp9SL

### Callbacks
https://go.dev/play/p/QRK7liUukIw

### Closure
E capturar um contexto
https://go.dev/play/p/MFOeIVXSTK5

### Recursao
https://go.dev/play/p/VjMbFCpnnWH

## Ponteiros

### O que sao Ponteiros
Todos os valores armazenados no computador, ficam armazenados na memoria.Todo ponteiro se refere a esse endereço na memoria. Voce usa os ponteiros atraves do operador `&` e tambem pode de-reference usando `*`. Exemplo:

```go
package main

import "fmt"

func main() {

	x := 15

	y := &x

	fmt.Println(x, y) // vai exibir o valor de x e o endereço de referencia da variavel de x (que no caso, é a variavel y)
	fmt.Println(*y)   // aqui testamos fazendo o de-reference do valor de y, ou seja, podemos acessar 10

	fmt.Printf("tipo de y %T, tipo de x: %T \n", y, x) // Tipo de y é um ponteiro para um int e x é um int.

	*y++ // essa é uma forma de alterar o valor do que esta armazenado no endereço da memoria

	fmt.Println(x, y)

}
```

### Quando usar ponteiros
Em go, tudo é pass-by-value, entao os ponteiros precisam ser usados quando queremos compartilhar esse valor e nao ficar criando valores distintos e ocupando mais a memoria. As ocasioes que os ponteiros podem ser uteis sao: quando nao queremos passar grandes volumes de dados pra la e pra ca e quando queremos mudar um valor em sua localização original. Exemplo:

```go
package main

import "fmt"

func main() {

	x := 0

	estarecebeovalor(x)
	estarecebeoponteiro(&x)
	fmt.Println(x)

}

func estarecebeovalor(x int) { // essa funcao faz uma copia da variavel x e altera o valor dela, mas nao muda x
	x++

}

func estarecebeoponteiro(x *int) { // essa funcao faz recebe o ponteiro da variavel x e altera o valor que esta na memoria. X é incrementado
	*x++
	fmt.Println("Valor de x dentro da funcao: ", *x)
}
```

## JSON
Tudo que vai precisar ser exportado em go, é necessario ter a primeira letra maiscula. Isso vale pra struct, ou seja, todos os campos que vao ser exportados para um json, precisamos ter a primeira letra maiscula. 
Playground com marshal: https://go.dev/play/p/J2ph8XbN1ko
Caso voce tenha um campo no json e um outro campo internamente no codigo, voce pode usar notacoes json. 
Playground com unmarshal: https://go.dev/play/p/TuN8USWmEJ1
Temos a opcao de fazer com o encoder, ao inves do marshal. O encoder é "parecido" quando fazemos um parse de um json de uma requisicao no node js, ja que a request.data é uma stream. No encoder, voce precisa dar um "pipe" nos resultados. 
Exemplo do encode: https://go.dev/play/p/u2JvgaWscbP

### Sort customizado - Sobrepondo um metodo go
Tipos de sort: https://go.dev/play/p/SzsVzjnwe_K
Sort customizado: https://go.dev/play/p/oFJuajwrJ29

### Bcrypt
https://go.dev/play/p/x8QGZUq1vMl

## Concorrencia

### Concorrencia e paralelismo
A concorrencia, por design, faz com que varias coisas sejam executadas de forma independentes. Varias coisas acontecem ao mesmo tempo e uma coisa nao depende da outra. A concorrencia é um padrao, uma forma de pensar o seu programa em si. Nao importa se voce vai ter 1 ou 1000 processadores, ele vai rodar da mesma forma. Se voce so tem um processador, nada acontece de maneira paralela, so acontece de maneira sequencial, mas o sistema operacional, implementando threads, permite que isso seja executada de forma concorrente. DOS (um SO que so permitia rodar um programa por vez) e quando o windows surge, já é possivel rodar mais de um programa pro vez. O paralelismo acontece quando a gente tem processos concorrentes sendo executados em um processador de multiplos nucleos ou mais de um processador. Por padrao, o Go executa o codigo usando o maximo de processadores possiveis.

### Goroutines e waitGroups
Nao é threads, mas é proximo disso. Duas Goroutines é como se fossem programas separados e independentes, mas que podem se comunicar entre si. Toda vez que a gente coloca a palavra `go` antes de `func`, ela automaticamente se torna uma goroutine. A partir do momento que essa funçao for executada, voce nao tem mais controle sobre ela.
https://go.dev/play/p/-AJUiPHXF_r
Sempre precisamos "esperar" pelo nossa gorountine acabar para poder encerrar a funcao main. Usamos o `sync.WaitGroup` pra fazer essa sincronizacao.
Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
    - func Add: "Quantas goroutines?"
    - func Done: "Deu!"
    - func Wait: "Espera todo mundo terminar."
https://go.dev/play/p/00sUihSBAjY

### Race condition
Programar de forma concorrente pode ser dificil por conta da sutilezas necessarias para acessar as variaveis compartilhadas. Por exemplo: ter varias goroutines acessando a mesma variavel, vai dar ruim!. Em Go, os valores que a gente quer compartilhar sao passados atraves de canais, e nao em threads separadas da execucao. E esse valor nunca esta no mesmo lugar ao mesmo tempo.
- Race condition: 
        Função 1       var     Função 2
         Lendo: 0   →   0
         Yield          0   →   Lendo: 0
         var++: 1               Yield
         Grava: 1   →   1       var++: 1
                        1   ←   Grava: 1
         Lendo: 1   ←   1
         Yield          1   →   Lendo: 1
         var++: 2               Yield
         Grava: 2   →   2       var++: 2
                        2   ←   Grava: 2

Mutex é quando voce "trava" uma variavel ate ela ser modificada.

Testar local depois: https://go.dev/play/p/trnnwi0-dTA

### Mutex
Ele "trava" uma variavel e deixa so uma goroutine mexer com aqueles valores num dado momento. Todas as outras threads que querem usar aquilo precisam esperar o fim da execucao dessa goroutine. Uma trava de exclusao mutua

## Channels
É a maneira de transferir dados entre go routines. Forma correta de fazer sincronizacao e codigo concorrente. Serve pra ordenar, sincronizar, orquestrar e um certo buffering. Canais sao como corredores de prova de revezamento. Eles precisam "passar o bastao" pro proximo, e caso nao tenha, isso vai dar um erro. Se voce tentar inserir uma informação, e printar, vai dar um erro. Exemplo: https://go.dev/play/p/yRuz-kejNTm. É possivel usar buffer dentro do canal, o que faz o canal armazenar os valores, mas que automaticamente, bloqueiam usando o tamanho. Exemplo: https://go.dev/play/p/e2jGlXtPpbq. Caso voce tente inserir qualquer item a mais, vai dar erro. Mas por via de regra, o buffer nao é muito usado.

### Direcionais e utilizando.
Canais podem ser direcionais ou bidirecionais. Os canais podem só receber informaçao ou so enviar informacao (direcionais). A seta sempre aponta pra o lado do tipo de canal. Exemplo da sintaxe:
```go
    - send chan ←
        - error: "invalid operation: ←cs (receive from send-only type chan ←  int)"
    - receive ← chan
        - error: "invalid operation: cr ← 42 (send to receive-only type ← chan int)"
```
Exemplos: https://go.dev/play/p/utFOO-RxYPn / Tipagem de canais: https://go.dev/play/p/H1uk4YGMBB
Voce pode ter o canal sendo bidirecional, mas voce tambem pode utilizar ele só como receiver ou só como sender, depende de como voce tipa ele.

### Range e close (canais)
Pratico: https://go.dev/play/p/MKPBJBv_JFg

### Select 
Funciona como switch, so que pra canais. Não é sequencial.
Exemplos praticos:
https://go.dev/play/p/RsSszu-YGKu
https://go.dev/play/p/olLK50ppTBO
https://go.dev/play/p/-cYNCEQEbJG
https://go.dev/play/p/BFMsCi6UYUC

### Tratamento de erros
https://go.dev/play/p/Ax8iIuHaYdc

### Recover
https://go.dev/play/p/Fd2FbXaEaan

## Erros em go

### Tratamento de errors
Em go, nah exist `try-catch-finally`.  Tipo erro representa Como os erros sao implementados. Se caso houver um erro, a interface retorna uma string, caso contrario, retorno nil. Criar o habito de usar sempre defer. Sempre lide com erro imediatamente.

### Erros com informacoes adicionais
Formas padroes de customizar o erros: https://go.dev/play/p/GcP4SioAgD5
type + interface: https://go.dev/play/p/FyVo_AYwjb8 

## Documentacao
O comando `godoc` gera a documentacao local via http. O comando `go doc` vai ler documentacao de acordo com o package no terminal. Voce pode rodar toda doc do go localmente com o comando `godoc -http:8080`, que tambem contem a sua documentacao

## Tests e benchmarks

### Intro
Tests devem: ficar num arquivo que termine com `_test.go`, que é a forma que o go usa para identificar e rodar arquivos de testes. Tambem devem ficar na mesma package que o codigo a ser testado. E as funções de testes seguem um padrão `func TestNome(t *testing.T)`, onde voce pode usar esse parametro `t` para falhas. Voce pode rodar os testes usando `go test` ou `go test -v`, esse vai te mostrar os resultados de cada test case.

### Testes em tabela
Teste em serie para testar diferentes variações.

### Ferramentas nativas
`gofmt` é a ferramenta para formatar o codigo. `go vet` vai procurar por possíveis erros e `golint` vai trazer sugestoes pro codigo.

### Benchmarks
Permitem testar a velocidade ou performance do codigo. Precisa usar o `_test.go` e a sintaxe é `func BenchmarkFunc (b *testing.B) { for i := 0; i < b.N; i++ }`. Esse codigo vai rodar seu codigo diversas vezes e em diversos cenários. Pra rodar todos os testes de benchmark, voce usa `go test -bench .` ou, para rodar apenas uma funcao, seria `go test -bench NomeFuncao`.

### Cobertura
Para exibir a cobertura dos testes. Voce precisa utilizar a flag `-coverprofile [arquivo]` para salvar a analise de cobertura em um arquivo. Comandos: 
```go
go test -cover
go test -coverprofile c.out
go test -cover -html=c.out >>>> abre no browser
```
