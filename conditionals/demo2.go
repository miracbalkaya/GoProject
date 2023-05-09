package conditionals

import "fmt"

func Demo2() {

	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Println("Limit yetersiz")
	} else if cekilmekIstenen == hesap {
		fmt.Println("İşlem tamamlandı,bakiye tükendi")
	} else {
		fmt.Println("İşlem tamamlandı")
	}

}
