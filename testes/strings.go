package main

import (
	"fmt"
	"strings"
)

func main() {
	//Primeira parte:
	// str1 := `Walker` //sequência imutável de bytes
	// fmt.Println(str1)
	// str1 = "João"
	// fmt.Println(str1)
	// str2 := "Olá"
	// fmt.Println(len(str2)) //4 bytes
	// fmt.Println(str2[1])

	// for i, r := range str2 {
	// 	fmt.Printf("Índice %d: %c\n", i, r)
	// }

	// runes := []rune(str2)
	// fmt.Println("Número de caracteres:", len(runes))

	//Segunda parte:
	// s := "Olá, Go!"

	// // Converter para slice de bytes
	// b := []byte(s)
	// fmt.Printf("Bytes (hex): % x\n", b)

	// // Converter para slice de runes
	// r := []rune(s)
	// fmt.Println("Número de runes:", len(r))
	// fmt.Print("Runes: ")
	// for _, runeValue := range r {
	// 	fmt.Printf("%c ", runeValue)
	// }
	// fmt.Println()

	// // Contando runes usando a função do pacote utf8
	// count := utf8.RuneCountInString(s)
	// fmt.Println("Contagem de runes (utf8):", count)

	//Terceira parte:
	texto := "  Go é incrível! O Go é rápido e simples.  "

	// Trim: remove espaços em branco do início e fim
	textoTrim := strings.TrimSpace(texto)
	fmt.Println("Trim:", textoTrim)

	// ToLower e ToUpper: convertem para minúsculas e maiúsculas
	fmt.Println("ToLower:", strings.ToLower(textoTrim))
	fmt.Println("ToUpper:", strings.ToUpper(textoTrim))

	// Contains: verifica se contém uma substring
	contem := strings.Contains(textoTrim, "rápido")
	fmt.Println("Contém 'rápido'? ", contem)

	// Index: posição da primeira ocorrência de "Go"
	pos := strings.Index(textoTrim, "Go")
	fmt.Println("Índice da primeira ocorrência de 'Go':", pos)

	// Split: divide a string com base em um separador (neste caso, espaço)
	partes := strings.Split(textoTrim, " ")
	fmt.Println("Split:", partes)

	// Join: junta os elementos do slice em uma única string, usando hífen como separador
	junta := strings.Join(partes, "-")
	fmt.Println("Join:", junta)

	// Replace: substitui todas as ocorrências de "Go" por "Golang"
	novoTexto := strings.Replace(textoTrim, "Go", "Golang", -1)
	fmt.Println("Replace:", novoTexto)
}
