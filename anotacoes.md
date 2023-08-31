# Função main 
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
