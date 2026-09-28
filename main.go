package main

import "fmt"


var a int = 10
var b int = 20
var c uint16 = 30
var d uint8 = 40
var e uint64 = 50

var f int64 = 60
var g int32 = 70
var h int16 = 80
var i int8 = 90

const name string = "joao"
var array [5]int = [5]int{1, 2, 3, 4, 5}

var slice []int = []int{1, 2, 3, 4, 5} //slice é um vetor dinamico
var portas []int = []int{80, 443, 8080}



var nome string
func main() {
	fmt.Println("rodano")
	if a > b {
		fmt.Println("a é maior que b")
	}

	fmt.Println(array[2])
	fmt.Println(slice[4])	
	slice = append(slice, 6)
	fmt.Println(slice[5])
	array[2] = 10
	fmt.Println(array[2])
	
	for i := 0; i < len(portas); i++ {
		fmt.Println(portas[i])
	}

    fmt.Println("rodando")
    fmt.Print("qual seu nome?: ")
    _, err := fmt.Scanln(&nome)
    if err != nil {
        fmt.Println("erro ao ler:", err)
        return
    }

    fmt.Println("opa: ", nome)
	

}