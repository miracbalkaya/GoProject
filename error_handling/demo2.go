package error_handling

import (
	"errors"
	"fmt"
)

func tahminEt(tahmin int) (string, error) {

	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı girniz", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı girniz", nil
	}
	return "Bildiniz", nil
}
func Demo2() {
	mesaj, hata := tahminEt(498)
	fmt.Println(mesaj)
	fmt.Println(hata)

}
