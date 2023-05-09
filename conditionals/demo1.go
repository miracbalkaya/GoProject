package conditionals

import "fmt"

func Demo1() {

	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Limit yetersiz")
	}
	if cekilmekIstenen <= hesap {
		fmt.Println("İşlem tamamlandı ")
		hesap = hesap - cekilmekIstenen
	}
	fmt.Println("Güncel bakiye : " + fmt.Sprintf("%v", hesap))

}
