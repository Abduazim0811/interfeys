package main

import (
	"fmt"
	"os"
	sin "Product/internal/Users/SIGNIN"
	sup "Product/internal/Users/SIGNUP"
	ad   "Product/Internal/admin"
)

func main() {
	fmt.Println("------------------------")
	fmt.Println("Qaysi xizmatni tanlaysiz")
	
	fmt.Println("[1] Users")
	fmt.Println("[2] Admin")
	fmt.Println("[3] Exit")
	fmt.Println("------------------------")
	var son int
	var son2 int
	fmt.Scanln(&son)
	if son == 1 {
		fmt.Println("[1] Sign up")
		fmt.Println("[2] Sign in")
		fmt.Scanln(&son2)
		if son2==1{
			sup.SIgn_Up()
		}else if son2==2{
			sin.Sign_in()
		}	
	} else if son == 2 {
		// ad.Admin()
		fmt.Println("Admin panel mavjud emas tez orada qo'shiladi")
		// pr.PRoduct()
	} else {
		os.Exit(0)
	}
}
