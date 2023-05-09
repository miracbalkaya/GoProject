package arrays

import "fmt"

func Demo4() {

	var sayilar1 [2][3]int

	sayilar1[0][0] = 1
	sayilar1[0][1] = 3
	sayilar1[0][2] = 5
	sayilar1[1][0] = 2
	sayilar1[1][1] = 4
	sayilar1[1][2] = 6

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar1[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}
