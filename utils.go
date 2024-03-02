package main

import "fmt"

func PrettyPrintMap(m map[byte][]byte) {
	for k, v := range m {
		fmt.Printf("%c: %s\n", k, v)
	}
	fmt.Println()
}

func PrettyPrintGrid(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		if len(grid[i]) == 0 {
			continue
		}
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%d %c ", grid[i][j], grid[i][j])
		}
		fmt.Println()
	}
}
