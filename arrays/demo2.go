package arrays

import "fmt"

func Demo2() {

	var sehirler [5]string
	sehirler[0] = "ankara"
	sehirler[1] = "istanbul"
	sehirler[2] = "muş"
	sehirler[3] = "antalya"
	sehirler[4] = "izmir"
	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}

}
