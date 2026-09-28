# NOTAS SOBRE OS APRENDIZADOS EM GO
 - vetores (slice)
 - array
 - tipos (uint, int, int64/32/16/8...)

# codigo base de arquivo .go
<hr>

 ```go
 package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

 func main () {
    var nome string

    fmt.Println("rodando")
    fmt.Print("qual seu nome?: ")
    _, err := fmt.Scanln(&nome)
    if err != nil {
        fmt.Println("erro ao ler:", err)
        return
    }

    fmt.Println("opa: ", nome)
 }
 ```
scanln é simples, pra ler uma linha inteira com espaços, use: bufio
## IDEIA DE PROJETO GOLANG PRA MAIOR ESTRUTURA NO LONGO PRAZO:

