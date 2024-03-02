package main

import (
	"math"
	"sort"
)

const ALPHABET = "ABCDEFGHIKLMNOPQRSTUVWXYZ"

func GetIndex(c byte, table [][]byte) (x int, y int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == c {
				return i, j
			}
		}
	}
	return
}

func PlayFairCipher(key, plainText string) (encrypted string) {
	const n = 5
	table := [][]byte{}
	visited := make(map[byte]bool)

	for i := 0; i < n; i++ {
		table = append(table, make([]byte, n))
	}

	for _, c := range ALPHABET {
		visited[byte(c)] = false
	}

	i, j := 0, 0

	// fill the table
	for _, c := range key {
		if !visited[byte(c)] {
			visited[byte(c)] = true
			if c == 'I' || c == 'J' {
				visited['I'] = true
				visited['J'] = true
			}
			table[i][j] = byte(c)
			j++
			if j == n {
				j = 0
				i++
			}
		}
	}

	// fill with remaining letters
	for _, c := range ALPHABET {
		k := byte(c)
		if !visited[k] {
			table[i][j] = k
			j++
			if j == n {
				j = 0
				i++
			}
		}
	}

	var modifiedPlainText string
	for i := 0; i < len(plainText); i++ {
		if i+1 < len(plainText) && plainText[i] == plainText[i+1] {
			modifiedPlainText += string(plainText[i]) + "X"
		} else {
			modifiedPlainText += string(plainText[i])
		}
	}

	if len(modifiedPlainText)%2 != 0 {
		modifiedPlainText += "X"
	}

	// encryption
	for i := 0; i < len(modifiedPlainText); i += 2 {
		l1 := modifiedPlainText[i]
		l2 := modifiedPlainText[i+1]

		x1, y1 := GetIndex(l1, table)
		x2, y2 := GetIndex(l2, table)

		if x1 == x2 {
			encrypted += string(table[x1][(y1+1)%n])
			encrypted += string(table[x2][(y2+1)%n])
		} else if y1 == y2 {
			encrypted += string(table[(x1+1)%n][y1])
			encrypted += string(table[(x2+1)%n][y2])
		} else {
			encrypted += string(table[x1][y2]) + string(table[x2][y1])
		}
	}

	// prettyPrintGrid(table)

	return
}

func VigenereCipher(key, plainText string) (encrypted string) {
	if len(plainText) > len(key) {
		for len(plainText) > len(key) {
			key += key
		}
	}

	keyIndex := 0

	for _, c := range plainText {
		mi := int(c - 'A')
		ki := int(key[keyIndex] - 'A')

		encrypted += string(byte((mi+ki)%26) + 'A')
		keyIndex++
	}

	return
}

func ColumnarTransposition(key, plainText string) (encrypted string) {
	table := [][]byte{}
	keys := []byte(key)
	visited := make([]bool, len(keys))

	nRows := math.Ceil(float64(len(plainText)) / float64(len(key)))

	for i := 0; i < int(nRows); i++ {
		table = append(table, make([]byte, len(key)))
	}

	i, j := 0, 0
	for _, c := range plainText {
		table[i][j] = byte(c)
		j++
		if j == len(key) {
			j = 0
			i++
		}
	}

	// fill empty rows
	for i, row := range table {
		for j, val := range row {
			if val == 0 {
				table[i][j] = 'X'
			}
		}
	}

	// PrettyPrintGrid(table)

	// sort the keys
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		var col int

		for i, v := range key {
			if byte(v) == k && !visited[i] {
				col = i
				visited[i] = true
				break
			}
		}

		for i := 0; i < len(table); i++ {
			encrypted += string(table[i][col])
		}
	}

	return
}
