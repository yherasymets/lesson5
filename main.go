package main

import "fmt"

const (
	applePrice = 5.99
	pearPrice  = 7
	aplleCount = 9
	pearCount  = 8
)

var myMoney float64 = 23

func main() {
	buyApple := aplleCount * applePrice
	buyPear := pearCount * pearPrice
	sum := buyApple + float64(buyPear)
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", sum)

	howMuchPear := myMoney / pearPrice
	howMuchApple := float64(myMoney) / applePrice
	fmt.Println("Скільки груш ми можемо купити?", howMuchPear)
	fmt.Println("Скільки яблук ми можемо купити?", howMuchApple)

	twoAndtwo := (2*applePrice + 2*pearPrice) < float64(myMoney)
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?", twoAndtwo)
}
