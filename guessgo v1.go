package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100) + 1

	fmt.Println("Sayı Tahmin Oyununa hoşgeldiniz. Lütfen 1 - 100 arasında bir sayı tahmin ediniz. Sayıyı doğru bilmek için altı hakkınız var")
	fmt.Println(randomNumber)

	var tahminHakki = 6

	if tahminHakki > 0 {
		for i := 0; i < tahminHakki; {

			var tahmin int

			fmt.Scanln(&tahmin)

			if tahmin == randomNumber {
				fmt.Println("Tebrikler doğru bildiniz! Cevap: ", randomNumber)

				break

			} else if tahmin > randomNumber {
				fmt.Println("Tahmin ettiğiniz sayı cevaptan daha büyük. Tekrar deneyin.")

				tahminHakki = tahminHakki - 1
			} else if tahmin < randomNumber {
				fmt.Println("Tahmin ettiğiniz sayı cevaptan daha küçük. Tekrar deneyin")

				tahminHakki = tahminHakki - 1
			} else {
				fmt.Println("Lütfen geçerli bir değer giriniz.")
			}

		}

	}

	if tahminHakki == 0 {
		fmt.Println("Elendiniz...")
	}

}