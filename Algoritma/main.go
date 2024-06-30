package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("===================SOAL 1======================")
	reverseAlphabet()
	fmt.Println("")
	fmt.Println("===================SOAL 2======================")
	longestSentence()
	fmt.Println("")
	fmt.Println("===================SOAL 3======================")
	matchQuery()
	fmt.Println("")
	fmt.Println("===================SOAL 4======================")
	matrix := [][]int32{
		{1, 2, 0},
		{4, 5, 6},
		{7, 8, 9},
	}

	diagonalDifference(matrix)

}

func reverseAlphabet() {
	input := "NEGIE1"
	var letters string
	var digits string
	fmt.Println("Input: ", input)
	// Memisahkan huruf dan angka
	for _, char := range input {
		if unicode.IsDigit(char) {
			digits += string(char)
		} else {
			letters += string(char)
		}
	}

	// Membalikkan urutan huruf
	var reversed string
	for i := len(letters) - 1; i >= 0; i-- {
		reversed += string(letters[i])
	}

	// Menggabungkan huruf yang sudah dibalik dengan angka
	result := reversed + digits

	fmt.Println("Reverse Alphabet: ", result)

}

func longestSentence() {
	input := "Saya sangat senang mengerjakan soal algoritma"
	fmt.Println("Input:", input)

	// Memisahkan kalimat menjadi kata-kata
	words := strings.Split(input, " ")

	longest := words[0]
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}

	fmt.Println(longest, ": ", len(longest))

}

func matchQuery() {
	input := [...]string{"xc", "dz", "bbb", "dz"}
	query := [...]string{"bbb", "ac", "dz"}
	fmt.Println("Input:", input)
	fmt.Println("Query:", query)

	matchCounts := make([]int, len(query))

	for i, q := range query {
		count := 0
		for _, word := range input {
			if word == q {
				count++
			}
		}
		matchCounts[i] = count
	}

	fmt.Println("Output:", matchCounts)
}

func diagonalDifference(arr [][]int32) int32 {

	var diagonal1 int32
	var diagonal2 int32
	n := len(arr)

	for i := 0; i < n; i++ {
		diagonal1 += arr[i][i]
		diagonal2 += arr[i][n-i-1]
	}

	fmt.Printf("diagonal pertama = %d + %d + %d = %d\n", arr[0][0], arr[1][1], arr[2][2], diagonal1)
	fmt.Printf("diagonal kedua = %d + %d + %d = %d\n", arr[0][2], arr[1][1], arr[2][0], diagonal2)

	difference := diagonal1 - diagonal2
	if difference < 0 {
		difference = -difference
	}

	fmt.Printf("maka hasilnya adalah %d - %d = %d\n", diagonal1, diagonal2, difference)

	return difference
}
