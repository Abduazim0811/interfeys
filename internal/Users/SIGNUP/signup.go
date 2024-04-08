package Signup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"strconv"
)

type Signup struct {
	First_name string
	Last_name  string
	Email      string
	Password   string
	Price      int
}

func SIGNUP(s_up Signup) {
	var sentence string
	var nm []string
	fmt.Printf("First_Name:\t")
	fmt.Scanln(&s_up.First_name)
	fmt.Printf("Last_Name:\t")
	fmt.Scanln(&s_up.Last_name)
	fmt.Printf("Email:\t")
	fmt.Scanln(&s_up.Email)
	fmt.Printf("Password:\t")
	fmt.Scanln(&s_up.Password)

	if len(s_up.Password) < 8 {
		fmt.Println("Parol kamida 8 ta belgidan kam bo'lishi mumkin emas!!!")
		fmt.Println("Parolni qayta kiriting")
		fmt.Printf("Password:\t")
		fmt.Scanln(&s_up.Password)
	}
	fmt.Printf("Price:\t")
	fmt.Scanln(&s_up.Price)

	nm = append(nm, s_up.First_name, s_up.Last_name, s_up.Email, s_up.Password, fmt.Sprintf("%d", s_up.Price))

	fileName := "/home/abduazim/Projects/Golang/NT_Homeworks/interfeys/users.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Faylni ochishda xatolik:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, s_up.Email) {
			fmt.Println("Bu email allaqachon ro'yxatdan o'tgan:", line)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	filee, err := os.OpenFile("/home/abduazim/Projects/Golang/NT_Homeworks/interfeys/users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Fayl ochilmadi:", err)
		return
	}
	defer filee.Close()

	sentence = strings.Join(nm, " ")
	yozuvchi := bufio.NewWriter(filee)
	fmt.Fprintln(yozuvchi, sentence)
	yozuvchi.Flush()

	fmt.Println("Siz muvaffaqiyatli kirdingiz🥳🥳🥳")
	fmt.Println("")
	Menu()
}
var slc []string
func Menu() {
	var num string
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
				SIgn_Up()
				break
			}else{
				Menu()
			}
		}

	}
	// Savatcha()
}


func SIgn_Up() {
	var s_up Signup
	SIGNUP(s_up)
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
		son2:=0
		fmt.Scanln(&son2)

		for i,char:=range slc{
			if i+1!=son2{
				resault=append(resault, char)
			}
		}
		filee, err := os.OpenFile("/home/abduazim/Projects/Golang/NT_Homeworks/interfeys/savatcha.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Fayl ochilmadi: ", err)
			return
		}
		defer filee.Close()
		
		for _ ,char:=range resault{
			yozuvchi := bufio.NewWriter(filee)
			fmt.Fprintln(yozuvchi, char)
			yozuvchi.Flush()
		}
		sum:=0

		for _ ,char:= range resault{
			natija:=strings.Split(char," ")
			numm,err:=strconv.Atoi(natija[2])
			if err!=nil{
				fmt.Println("Xatolik bor: ", err)
			}
			sum+=numm
		}

		if sum>=s_up.Price{
			fmt.Println("Sotildi")
		}else{
			fmt.Println("Xisobizda mag'lag' yetarli emas!!!")
		}
	}
}



