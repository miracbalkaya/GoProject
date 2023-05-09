package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)
	//fmt.Println(isimler)
	isimler[0] = "Mirac ,"
	isimler[1] = "aafaf ,"
	isimler[2] = "agaga ,"
	isimler = append(isimler, "ali")
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
