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

<hr>

# PROJETO: MICRO REATOR
### IDEIA
**usar diferentes tipos de entradas (uint8/16/32æ64) para controlar um reator**
*o projeto precisa manter ele instavel, ele vai ter seus valores alterados constantemente de acordo com um array*
*preciso: digitar numeros especificos de cada tipo quando o terminal pedir*
<strong>EXEMPLO</strong>
```bash
reator em 34% estabilize com um uint64 valido:
resposta: algum numero valido no valor de 64bits nao sinalizado
reator: verificando tipos...
reator: reator ok!
```
<hr>