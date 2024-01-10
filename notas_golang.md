# Golang

## Tipos 
- atomic data types

  - string
  - int
  - int32
  - int64
  - uint
  - uint32
  - uint64
  - float32
  - float64
  - bool

- unsafe
  - Pointers

- abstract data types 
  - map[]<datatype>
  - struct{}
  - interface{}


## Como declarar variáveis em Go

```
package main

import "fmt"

func main() {
	var m0 string
	m0 = "hello"
	m1 := "world!"

	var (
		m3 = " Now i'm coding"
		m4 = "using go!"
	)

	fmt.Println(m0 + " " + m1 + m3 + m4)
}

```

## Arrays

- sintaxe
```
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred
```
- examplos
```
	var arr = [2]int{1, 2}
	arr2 := []int{1, 3, 4, 5, 6, 7, 7}

	var arr3 = [...]string{"a", "b", "c"}

	fmt.Println(arr, arr2, arr3)
```

## Ponteiros

```
	var arr = [2]int{1, 2}

	arr2 := &arr

	//endereco de memoria
	fmt.Println(&arr2)
	// valor armazenado no endereco
	fmt.Println(*arr2)
```

## Struct

```
func main() {
	std := Student{
		Name:  "Bruno",
		Age:   29,
		Class: "C.comp",
	}

	fmt.Println(std)

}

type Student struct {
	Name  string
	Age   int
	Class string
}

```
                          
