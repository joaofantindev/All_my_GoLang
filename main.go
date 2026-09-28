package main

import "fmt"
//some variables with different types
	var age uint32 = 30
	var salary float64 = 4582.32
	var height float32 = 1.75
	var isActive bool = true
	var name string = "joao"
	
type Person struct {

}
//struct to remember

func number() {
	//function that processes a lot of numbers and manipulate them
}
func pointer(n *int) {
	//just for pointers
	*n *= 2
	fmt.Println(n)
	fmt.Println(*n)
}
func main() {
	fmt.Println("running")
	x := 10
	p := &x

	fmt.Println(p) //memory address of x
	fmt.Println(*p) //x value

	*p = 20 //change the value of x through the pointer
	fmt.Println(x) //x value after change
	pointer(&x) //call the pointer function with the address of x
	//deve ser apenas com o endereço de memoria
	fmt.Println(oi())
}
func oi() string {
	return "oi"
}
type Animal interface {
	Som() string
}
type Cachorro struct{}
type Gato struct{}

func (Cachorro) Som() string { return "AuAu!"}
func (Gato) Som() string { return "Miau!"}
/*
interface é um tipo que define um conjunto de métodos. 
Um tipo implementa uma interface se ele possui todos os métodos 
definidos na interface. No exemplo acima, a interface Animal define um método Som(), 
e as structs Cachorro e Gato implementam esse método.
Isso permite que você trate diferentes tipos de animais de forma uniforme, chamando o
método Som() em qualquer instância de Animal, independentemente de ser um Cachorro ou um Gato.
*/
func emitirSom(a Animal) {
	fmt.Println(a.Som())
}
// interface vazia (aceita qualquer tipo de dado)
func interfaces(v interface{}) {
	fmt.Println(v)
}

//goroutines, interfaces, concorrencia, channels