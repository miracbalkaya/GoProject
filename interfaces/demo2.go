package interfaces

import "fmt"

type Mortgage struct {
	creditPaymetTotal float64
	adress            string
	rate              float64
}

type Car struct {
	creditPaymetTotal float64
	carInfo           string
	rate              float64
}
type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymetTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymetTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymetTotal: 1000, adress: "istanbul"}
	credit2 := Mortgage{rate: 12, creditPaymetTotal: 5000, adress: "ankara"}
	credit3 := Car{rate: 15, creditPaymetTotal: 60000, carInfo: "BMW"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Toplam ödeme : ", total)

}
