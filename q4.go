package main

import (
	"bufio"
	"math"
	"os"
	"sort"
)

func CrypticConnectionInput() (playfairKey, vignereKey, transpositionKey, cipherText string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	playfairKey = scanner.Text()

	scanner.Scan()
	vignereKey = scanner.Text()

	scanner.Scan()
	transpositionKey = scanner.Text()

	scanner.Scan()
	cipherText = scanner.Text()

	return
}

func PlayFairDecryption(key, cipherText string) (plainText string) {
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

	// PrettyPrintGrid(table)

	for i := 0; i < len(cipherText); i += 2 {
		l1 := cipherText[i]
		l2 := cipherText[i+1]

		x1, y1 := GetIndex(l1, table)
		x2, y2 := GetIndex(l2, table)

		if x1 == x2 {
			if y1-1 < 0 {
				y1 = n - 1
			} else {
				y1 = y1 - 1
			}

			if y2-1 < 0 {
				y2 = n - 1
			} else {
				y2 = y2 - 1
			}

			plainText += string(table[x1][(y1)%n]) + string(table[x2][(y2)%n])
		} else if y1 == y2 {
			if x1-1 < 0 {
				x1 = n - 1
			} else {
				x1 = x1 - 1
			}

			if x2-1 < 0 {
				x2 = n - 1
			} else {
				x2 = x2 - 1
			}

			plainText += string(table[(x1)%n][y1]) + string(table[(x2)%n][y2])
		} else {
			plainText += string(table[x1][y2]) + string(table[x2][y1])
		}

	}

	// PrettyPrintGrid(table)
	return
}

func VigenereDecryption(key, cipherText string) (plainText string) {
	keyIndex := 0
	for i := 0; i < len(cipherText); i++ {
		plainText += string((cipherText[i]-key[keyIndex]+26)%26 + 'A')
		keyIndex = (keyIndex + 1) % len(key)
	}
	return
}

func ColumnarDecryption(key, cipherText string) (plainText string) {
	table := [][]byte{}
	originalTable := [][]byte{}
	keys := []byte(key)
	visited := make([]bool, len(keys))
	origiVisited := make([]bool, len(keys))

	nRows := math.Ceil(float64(len(cipherText)) / float64(len(key)))

	for i := 0; i < int(nRows); i++ {
		table = append(table, make([]byte, len(key)))
		originalTable = append(originalTable, make([]byte, len(key)))
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	i, j := 0, 0
	for _, c := range cipherText {
		table[i][j] = byte(c)
		i++
		if i == int(nRows) {
			i = 0
			j++
		}
	}

	// PrettyPrintGrid(table)

	for l, k := range key {
		for i, sorted := range keys {
			if byte(k) == sorted && !visited[i] && !origiVisited[l] {
				visited[i] = true
				origiVisited[l] = true
				for j := 0; j < int(nRows); j++ {
					originalTable[j][l] = table[j][i]
				}
			}
		}
	}
	// PrettyPrintGrid(originalTable)
	for _, row := range originalTable {
		for _, val := range row {
			if val != 'X' {
				plainText += string(val)
			}
		}
	}
	return
}

func CrypticConnection(playfairKey, vignereKey, tranpositionKey, cipherText string) (decrypted string) {
	return PlayFairDecryption(playfairKey,
		VigenereDecryption(vignereKey,
			ColumnarDecryption(tranpositionKey, cipherText)))
}
