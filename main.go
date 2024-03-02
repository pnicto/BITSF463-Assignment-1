package main

import "fmt"

func main() {
	keyword, permutation, plainText := modifiedAdfgvxInput()
	fmt.Println(modifiedAdfgvxCipher("GRAPHY", "CRYPTO", permutation, plainText, keyword))
}
