package main

import "fmt"

func main() {
	students := [4][3]int{
		{2, 2, 0},
		{2, 2, 2},
		{2, 1, 2},
		{2, 2, 2},
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if students[i][j] == 0 {
				fmt.Printf("%d行%d列学生旷课", i+1, j+1)
			}
		}
	}
}
