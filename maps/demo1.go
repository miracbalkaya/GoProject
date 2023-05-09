package maps

import "fmt"

func Demo1() {

	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"
	fmt.Println(sozluk["table"])
	fmt.Println("eleman sayısı: ", len(sozluk))
	fmt.Println(sozluk)
	delete(sozluk, "book")
	fmt.Println("eleman sayısı: ", len(sozluk))
	fmt.Println(sozluk)

	deger, varMi := sozluk["book"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu : ", varMi)

	//TODO Slice yapısı
	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2["glass"])

}
