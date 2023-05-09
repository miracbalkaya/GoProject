package loops

import "fmt"

func Demo3() {

	sayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 1

	fmt.Println("Bir sayı tahmin et")
	fmt.Scanln(&tahminEdilenSayi)

	for sayi != tahminEdilenSayi {
		if tahminEdilenSayi < sayi {
			fmt.Println("Daha büyük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
		if tahminEdilenSayi > sayi {
			fmt.Println("Daha küçük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}
	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena değil"
	}

	if tahminEdilenSayi == sayi {
		fmt.Printf("Doğru tahmin ettiniz. %v tahmin : %v", tahminSayisi, basariDurumu)
	}

}
