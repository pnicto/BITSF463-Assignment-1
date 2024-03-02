package main

import "fmt"

func main() {
	straddlingKey, num1, num2, cipherText, rails := Q3Input()
	fmt.Println(Q3(straddlingKey, num1, num2, cipherText, rails))
}
