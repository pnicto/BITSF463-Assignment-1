package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Q3Input() (straddlingKey string, num1 int, num2 int, cipherText string, rails int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	straddlingKey = scanner.Text()

	scanner.Scan()
	nums := scanner.Text()
	parts := strings.Split(nums, " ")
	num1, _ = strconv.Atoi(parts[0])
	num2, _ = strconv.Atoi(parts[1])

	scanner.Scan()
	cipherText = scanner.Text()

	scanner.Scan()
	rails, _ = strconv.Atoi(scanner.Text())

	return
}

func RailFenceDecryption(rails int, cipherText string) (plainText string) {
	table := make([][]byte, rails)
	for i := range table {
		table[i] = make([]byte, len(cipherText))
	}

	i, j := 0, 0
	diff := 1
	for _, c := range cipherText {
		table[i][j] = byte(c)
		i += diff
		j += 1

		if i == 0 || i == rails-1 {
			diff = -diff
		}
	}

	cipherTextIndex := 0

	for i := 0; i < rails; i++ {
		for j := 0; j < len(cipherText); j++ {
			if table[i][j] != 0 {
				table[i][j] = cipherText[cipherTextIndex]
				cipherTextIndex++
			}
		}
	}

	i, j = 0, 0
	diff = 1
	for range cipherText {
		plainText += string(table[i][j])
		i += diff
		j += 1
		if i == 0 || i == rails-1 {
			diff = -diff
		}
	}

	return
}

func StraddlingCheckerboardDecryption(straddlingKey string, cipherText string, num1, num2 int) (decoded string) {
	table := make([][]byte, 3)

	for i := range table {
		table[i] = make([]byte, 10)
	}

	k := 0

	for i, row := range table {
		for j := range row {
			if k >= len(straddlingKey) {
				break
			}
			if i == 0 && (j == num1 || j == num2) {
				continue
			}
			table[i][j] = straddlingKey[k]
			k++
		}
	}

	// decryption
	i := 0
	for i < len(cipherText) {
		x := cipherText[i] - '0'
		if x == byte(num1) {
			y := cipherText[i+1] - '0'
			decoded += string(table[1][y])
			i += 2
			continue
		} else if x == byte(num2) {
			y := cipherText[i+1] - '0'
			decoded += string(table[2][y])
			i += 2
			continue
		} else {
			decoded += string(table[0][x])
			i++
		}
	}

	// PrettyPrintGrid(table)
	return
}

func Q3(straddlingKey string, num1 int, num2 int, cipherText string, rails int) (decoded string) {
	middle := RailFenceDecryption(rails, cipherText)
	return StraddlingCheckerboardDecryption(straddlingKey, middle, num1, num2)
}
