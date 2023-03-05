package main

import (
	"fmt"

	go_math "github.com/RianIhsan/go-math"
	go_say_name "github.com/RianIhsan/go-say-name"
)

func main() {
ATAS:
	var menu int
	fmt.Println(go_say_name.SayHello())
	fmt.Println("Kalkulator APP")
	fmt.Println("==========================================")
	fmt.Println("1. Pertambahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Print("Pilih Operasi 1/2/3: ")
	fmt.Scan(&menu)
	fmt.Println("==========================================")
	switch menu {
	case 1:
		var (
			bil1, bil2 int
		)
		fmt.Print("Masukan Bilangan Pertama: ")
		fmt.Scan(&bil1)
		fmt.Print("Masukan Bilangan Kedua: ")
		fmt.Scan(&bil2)
		fmt.Println("==========================================")
		go_math.Pertambahan(bil1, bil2)
	case 2:
		var (
			bil1, bil2 int
		)
		fmt.Print("Masukan Bilangan Pertama: ")
		fmt.Scan(&bil1)
		fmt.Print("Masukan Bilangan Kedua: ")
		fmt.Scan(&bil2)
		fmt.Println("==========================================")
		go_math.Pengurangan(bil1, bil2)
	case 3:
		var (
			bil1, bil2 int
		)
		fmt.Print("Masukan Bilangan Pertama: ")
		fmt.Scan(&bil1)
		fmt.Print("Masukan Bilangan Kedua: ")
		fmt.Scan(&bil2)
		fmt.Println("==========================================")
		go_math.Perkalian(bil1, bil2)
	default:
		var again string
		fmt.Println("Operasi salah!, Mau Mengulang?:(y/n)")
		fmt.Scan(&again)
		switch again {
		case "y":
			goto ATAS
		case "n":
			fmt.Println("Terima kasih!")
		}

	}
}
