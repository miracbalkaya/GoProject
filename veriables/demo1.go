package veriables

import "fmt"

func Demo1() {
	var metin string = "miraç"
	fmt.Println(metin)
	fmt.Println("hello")
	fmt.Println("hello")

	var kdv int = 10
	fmt.Println(kdv)

	kdv3 := 34
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)

	var durum bool

	var metin1 string = "ali"
	var metin2 string = "ahmet"

	durum = metin1 != metin2
	fmt.Println(durum)
}
