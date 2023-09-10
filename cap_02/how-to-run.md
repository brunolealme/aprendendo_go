- go mod init main.go
- go run /{path}/main.go

blank operator = _
utilizado quando não queremos utilizar o valor de retorno de uma funcao

```
 bytezz,errors := fmt.Println("algo a ser printado")
 ---------------- o cenário acima causaria erro se uma das variáveis não fossem utilizadas
 _,errors := fmt.Println("algo a ser printado")
 
```