package functions

import "fmt"

func Topla(sayi1 int, sayi2 int, sayi3 int) int {
	toplam := sayi1 + sayi2 + sayi3

	return toplam
}

func Selam(kullanici string) {
	fmt.Println("merhaba", kullanici)

}
