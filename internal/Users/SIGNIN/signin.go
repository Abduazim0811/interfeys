package Signin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SignIn struct {
	Email    string
	Password string
	Price    int
}

func SIGNIN(s_in *SignIn) {
	file, err := os.Open("/home/abduazim/Projects/Golang/NT_Homeworks/interfeys/users.txt")
	if err != nil {
		fmt.Println("Faylni ochishda xatolik:", err)
		return
	}
	defer file.Close()

	fmt.Printf("Email:\t")
	fmt.Scanln(&s_in.Email)
	fmt.Printf("Password:\t")
	fmt.Scanln(&s_in.Password)
	fmt.Printf("Price:\t")
	fmt.Scanln(&s_in.Price)

	lampochka := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, s_in.Email) && strings.Contains(line, s_in.Password) {
			fmt.Println("Siz muvaffaqiyatli kirdingiz 🥳🥳🥳")
			lampochka = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scannerda xatolik:", err)
		return
	}

	if !lampochka {
		fmt.Println("Kirishda xatolik: email yoki parol noto'g'ri")
		return
	} else {
		Menu()
	}
}

// Menu
func Menu() {
	var num string
	var slc []string
	file3, err := os.Open("/home/abduazim/Projects/Golang/NT_Homeworks/interfeys/warehouse.txt")
	if err != nil {
		log.Fatalf("Faylni ochishda xatolik: %v", err)
	}
	defer file3.Close()

	scanner2 := bufio.NewScanner(file3)
	var arr []string
	for scanner2.Scan() {
		line := scanner2.Text()
		arr = append(arr, line)
		fmt.Println(line)
	}
	fmt.Println("")
	fmt.Println("Yuqoridagilardan birini tanlang👆🏻👆🏻👆🏻")
	fmt.Scanln(&num)

	cnt, cnt2 := 0, 0

	for _, work := range arr {
		natija := strings.Split(work, " ")
		// filee, err := os.OpenFile("/home/abduazim/Projects/Golang/NT_Homeworks/interfeys/savatcha.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		// if err != nil {
		// 	fmt.Println("Fayl ochilmadi: ", err)
		// 	return
		// }
		// defer filee.Close()
		if num == natija[0] {
			fmt.Println("Nechta olmoqchisiz: ")
			fmt.Scanln(&cnt)
			n, err := strconv.Atoi(natija[2])
			if err != nil {
				fmt.Println("Bu yerda xatolik bor!!!   ", err)
				return
			}
			cnt++
			sentence := strconv.Itoa(cnt2) + " " + string(natija[1]) + " " + strconv.Itoa(cnt*n)
			slc=append(slc, sentence)
			// yozuvchi := bufio.NewWriter(filee)
			// fmt.Fprintln(yozuvchi, sentence)
			// yozuvchi.Flush()

			fmt.Println("Yana nimadur olishni hohlaysizmi?")
			fmt.Println("[1]HA\t[0]YOQ")
			son:=0
			fmt.Scanln(&son)
			if son==0{
				Savatcha(slc)
				break
			}else{
				Menu()
			}
		}

	}
	// Savatcha()
}

func Sign_in() {
	var s_in SignIn
	SIGNIN(&s_in)
}

func Savatcha(slc []string){
	fmt.Println("")
	var resault []string

	fmt.Println("Savatdagi ma'lumotlar")
	for _,char:=range slc{
		fmt.Println(char)
	}

	fmt.Println("Savatdagi maxsulotlarni o'chirishni hohlaysizmi?")
	fmt.Println("[1]HA\t[0]YOQ")
	son:=0
	fmt.Scanln(&son)

	if son==1{
		fmt.Println("Qaysi malumotlarni o'chirishni hohlaysiz?")
		for _,char:=range slc{
			fmt.Println(char)
		}
 
	}
}
