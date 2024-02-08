package main

import "fmt"

func toplama(a int, b int) int {
	return a + b
}

func carpma(a int, b int) int {
	return a * b
}

func cikarma(a int, b int) int {
	return a - b
}

func bolme(a int, b int) int {
	return a / b
}

func main() {

	var yapilacakIslem string
	var sayi1 int
	var sayi2 int

	fmt.Println("Hesap Makinesine Hoşgeldiniz! Lütfen işlem seçiniz. (+, -, *, /)")
	fmt.Scanln(&yapilacakIslem)
	fmt.Println("Lütfen ilk sayıyı giriniz:")
	fmt.Scanln(&sayi1)
	fmt.Println("Lütfen ikinci sayıyı giriniz:")
	fmt.Scanln(&sayi2)

	if yapilacakIslem == "+" {
		fmt.Println(toplama(sayi1, sayi2))
	} else if yapilacakIslem == "-" {
		fmt.Println(cikarma(sayi1, sayi2))
	} else if yapilacakIslem == "*" {
		fmt.Println(carpma(sayi1, sayi2))
	} else if yapilacakIslem == "/" {
		fmt.Println(bolme(sayi1, sayi2))
	} else {
		fmt.Println("Lütfen geçerli bir işlem giriniz.")
	}

}