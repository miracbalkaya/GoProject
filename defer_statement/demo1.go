package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}

func B() {
	defer A()
	defer C()
	fmt.Println("b fonksiyonu çalıştı")
}

func C() {
	fmt.Println("c fonksiyonu çalıştı")
}
