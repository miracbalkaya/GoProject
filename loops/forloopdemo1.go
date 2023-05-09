package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba dünya"
	i := 1
	//infinite loop
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}

}
