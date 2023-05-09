package main

import "golesson/error_handling"

func main() {

	//veriables.Demo1()
	//fmt.Println()
	//loops.Demo5()
	//functions.Selam("Miraç")
	//var sonuc = functions.Topla(2,3,4)
	//fmt.Println(sonuc * 10)
	//sonuc1,sonuc2,sonuc3,_ := functions.DortIslem(8,5)
	//fmt.Println("toplam :", sonuc1)
	//fmt.Println("çarpma :", sonuc2)
	//fmt.Println("bölme :", sonuc3)
	//fmt.Println("çıkarma :", sonuc4)

	//fmt.Println(functions.ToplaVariadic(1,2,3,4,5,))
	//fmt.Println(functions.ToplaVariadic(1,2,3))
	//sayilar := []int {4,6,3}
	//fmt.Println(functions.ToplaVariadic(sayilar...))
	//maps.Demo1()
	//sayi := 20
	//pointers.Demo1(&sayi)
	//fmt.Println("maindeki sayı : ",sayi)

	//sayilar := []int{1,2,3}
	//pointers.Demo2(sayilar)
	//fmt.Println("maindeki sayı : ",sayilar[0])

	//go goroutines.CiftSayilar()
	//go goroutines.TekSayilar()
	//time.Sleep(5 * time.Second)
	// fmt.Println("Main bitti")
	//ciftSayiCn := make(chan int)
	//tekSayiCn := make(chan int)
	//go channels.CiftSayilar(ciftSayiCn)
	//go channels.TekSayilar(tekSayiCn)

	//ciftSayiToplam, tekSayiToplam := <- ciftSayiCn, <- tekSayiCn

	//carpim := ciftSayiToplam * tekSayiToplam
	//fmt.Println("Çarpım : ",carpim)

	error_handling.Demo2()

}
