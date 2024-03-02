package main

import (
	"bufio"
	"os"
)

func modifiedAdfgvxInput() (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	keyword := scanner.Text()

	scanner.Scan()
	permutation := scanner.Text()

	scanner.Scan()
	plainText := scanner.Text()

	return keyword, permutation, plainText
}

func modifiedAdfgvxCipher(key1, key2, permutation, plainText, transpositionKey string) (encrypted string) {
	// cols
	m := len(key1)
	// rows
	n := len(key2)
	table := [][]byte{}
	for i := 0; i < n; i++ {
		table = append(table, make([]byte, m))
	}

	i, j := 0, 0
	for _, c := range permutation {
		table[i][j] = byte(c)
		j++
		if j == m {
			j = 0
			i++
		}
	}

	for _, c := range plainText {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if table[i][j] == byte(c) {
					encrypted += string(key2[i]) + string(key1[j])
				}
			}
		}
	}

	// PrettyPrintGrid(table)

	encrypted = ColumnarTransposition(transpositionKey, encrypted)
	return
}
