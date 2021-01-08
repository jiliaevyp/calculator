package main

import "fmt"

func main() {

	var mass [20]int
	fmt.Println("")
	fmt.Println("=================================================")
	fmt.Println("Первый запуск сложнейшей программы на Go! ))))))))")
	fmt.Println("==================================================")
	fmt.Println("")
	fmt.Println("Ну теперь держись интернет!")
    summ := 0
	for i := 1; i <= len(mass); i++ {
		mass[i-1] = i
		summ += mass[i-1]
		fmt.Print(i,"  ", mass[i-1])
		if i % 2 == 0 {
			fmt.Print(" четный")
		} else {
			fmt.Print(" нечетный")
		}
		switch i {
		case 0: fmt.Print(" ноль", "\n")
		case 1: fmt.Print(" один", "\n")
		case 2: fmt.Print(" два", "\n")
		case 3: fmt.Print(" три", "\n")
		case 4: fmt.Print(" четыре", "\n")
		case 5: fmt.Print(" пять", "\n")
		case 6: fmt.Print(" шесть", "\n")
		case 7: fmt.Print(" семь", "\n")
		case 8: fmt.Print(" восемь", "\n")
		case 9: fmt.Print(" девять", "\n")
		case 10: fmt.Print(" десять", "\n")
		}
		fmt.Print("-------------------", "\n")
		fmt.Print( mass)
		fmt.Print("сумма = ", summ, "\n")
	}

}
