// 1-masala
// package main

// import (
// 	"fmt"
// )

// type Animal interface {
// 	MakeSound()
// }

// type Dog struct{}

// func (c Dog) MakeSound() {
// 	fmt.Println("Woooow")
// }

// type Cat struct{}

// func (c Cat) MakeSound() {
// 	fmt.Println("Meow")
// }

// type Bird struct{}

// func (c Bird) MakeSound() {
// 	fmt.Println("Chirp")
// }

// func main() {
// 	animals := []Animal{Dog{}, Cat{}, Bird{}}

// 	for _, animal := range animals {
// 		animal.MakeSound()
// 	}
// }

// 2-masala
package main

import (
	"fmt"
)

type Employee interface {
	CalculateSalary() float64
	DisplayDetails()
}

type FullTimeEmployee struct {
	Ism           string
	Familiya      string
	OylikMaoshi   float64
	IshlaganSoati float64
}

func (fte FullTimeEmployee) CalculateSalary() float64 {
	return fte.OylikMaoshi * fte.IshlaganSoati
}

func (fte FullTimeEmployee) DisplayDetails() {
	fmt.Println("Ism:", fte.Ism)
	fmt.Println("Familiya:", fte.Familiya)
	fmt.Println("Oylik maoshi:", fte.OylikMaoshi)
	fmt.Println("Ishlagan soati:", fte.IshlaganSoati)
	fmt.Println("Jami maoshi:", fte.CalculateSalary())
}

type ContractEmployee struct{
	Ismi string
	Familiyasi string
	SoatlikMaoshi float64
	IshlaganSoati float64
}

func (ce ContractEmployee) CalculateSalary() float64 {
	return ce.SoatlikMaoshi * 160
}

func (ce ContractEmployee) DisplayDetails() {
	fmt.Println("Ismi:", ce.Ismi)
	fmt.Println("Familiyasi:", ce.Familiyasi)
	fmt.Println("Soatlik maoshi:", ce.SoatlikMaoshi)
	fmt.Println("Ishlagan soati:", ce.IshlaganSoati)
	fmt.Println("Jami maoshi:", ce.CalculateSalary())

}
func main() {
	var employees []Employee
	employees = append(employees, FullTimeEmployee{Ism: "Abduazim", Familiya: "Yusufov", OylikMaoshi: 5000, IshlaganSoati: 160})
	employees = append(employees, FullTimeEmployee{Ism: "Abdulaziz", Familiya: "Xoliqulov", OylikMaoshi: 2000, IshlaganSoati: 160})
	employees = append(employees, ContractEmployee{Ismi: "Shaxboz", Familiyasi: "Vahobov", SoatlikMaoshi: 1000, IshlaganSoati: 160})
	employees = append(employees, ContractEmployee{Ismi: "Doston", Familiyasi: "Soliyev", SoatlikMaoshi: 2000, IshlaganSoati: 160})
	fmt.Println("To'liq kunlik xodimlar:")
	for _, empemploye:=range employees{
		if _, ok := empemploye.(FullTimeEmployee); ok {
			empemploye.DisplayDetails()
			fmt.Println("")
		}else if _, ok := empemploye.(ContractEmployee); ok {
			empemploye.DisplayDetails()
			fmt.Println("")
		}
	}

}0

// leetcode
package main

import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {
	cnt := 0
	res:=""
	for _, char := range s {
		if char == '(' {
			cnt++
		} else if char == ')' {
			cnt--
		}
		if cnt<0{
			continue
		}
		res+=string(char)
	}
	return res
}

func main() {
	s := "))(("
	resault := minRemoveToMakeValid(s)
	fmt.Println(resault)
}
