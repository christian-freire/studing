package main

import (
	"fmt"
	"unsafe"
)

var values struct {
	idade int32 // 4 Bytes começando na posição 0 da primeira palavra
	// Padding de 4 Bytes e relação a palavra que alocou 8 Bytes
	nota int64 // 8 Bytes começando na posição 0 da segunda palavra
	// Padding 0
	nerd bool // 1 Bytes começando na posição 0 da terceira palavra
	// Padding de 7 Bytes e relação a palavra que alocou 8 Bytes
} // Total ocupado em memória = 24 Bytes com 11 Bytes de padding.

var values2 struct {
	nota int64 // 8 Bytes começando na posição 0 da segunda palavra
	// Padding 0
	idade int32 // 4 Bytes começando na posição 0 da primeira palavra
	nerd  bool  // 1 Bytes começando na posição 0 da terceira palavra
	// Padding de 3 Bytes com relação a palavra que alocou 8 Bytes
} // Total ocupado em memória = 16 Bytes em memória.

func main() {
	// Variável do tipo estrutura contendo um atributo idade do tipo int32, que vai ocupar
	// 4 bytes de memória alocada
	var value struct {
		idade int32
	}

	// Variável intener do tipo integer, que não tem o tamanho de bits especificado, por esse motivo
	// irá ocupar o mesmo valor máximo da palavra computacional do seu sistema (que quase sempre é o mesmo
	// da arquitetura do seu processador), ou seja, no meu caso x64, 64 bits, que equivale a 8 bytes.
	var integer int
	var integer32 int32

	fmt.Printf("%v\n", unsafe.Sizeof(integer))
	fmt.Printf("%v\n", unsafe.Sizeof(integer32))

	fmt.Printf("%v\n", unsafe.Sizeof(value))

	padding()
}

func padding() {
	fmt.Println("Padding var")
	fmt.Printf("%v\n", unsafe.Sizeof(values))
	fmt.Printf("%v\n", unsafe.Sizeof(values2))
}
