package main

import (
	"fmt"
)

func testZifry() {
	var mass [100]int
	fmt.Println("Введите размер массива")
	fmt.Scanf("%d\n",&num)
	fmt.Println("")
	fmt.Println("=================================================")
	fmt.Println("Первый запуск сложнейшей программы на Go! ))))))))")
	fmt.Println("==================================================", "\n")
	fmt.Println("")
	fmt.Println("Ну теперь держись интернет!")
    summ := 0
	fmt.Println("kol=", num)
	for i := 1; i <= num; i++ {
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
		//fmt.Print( mass)
		fmt.Print("сумма = ", summ, "\n")
	}

}

func callcul() {
	var arg1, arg2, result int64
	var mach string
	const plus string = "+"
	const minus string = "-"
	const mult string = "*"
	const div string = "/"
	fmt.Println("Это калькулятор")
	fmt.Print("-------------------", "\n")
	arg1, arg2, mach = inparg()
	switch mach {
	case plus: result = arg1 + arg2
	case minus: result = arg1 + arg2
	case mult: result = arg1 + arg2
	case div: result = arg1 + arg2
	}
	fmt.Print(arg1, " ", mach, " ", arg2, " = ", result,"\n")
	fmt.Print("-------------------")
}

func inparg() (int64, int64, string) {
	var arg1, arg2 int64
	var oper_err, divzero_err int
	var mach string
	const div string = "/"
	fmt.Println("Введите первое число")
	fmt.Scanf("%d\n",&arg1)
	fmt.Println("Введите второе число")
	fmt.Scanf("%d\n",&arg2)
	mach = machen()
	if mach != "?" || mach == div && arg2 != 0 {
		return arg1, arg2, mach
	}

	if mach == div && arg2 == 0 {
		fmt.Print("Делить на ноль нельзя", "\n")
		return arg1, arg2, mach
		}
		return arg1, arg2, mach
	} else {
		fmt.Print("Введите правильную операцию", "\n")
		return 0, 0, "+"
	}


	//return()
}
func machen() (string) {
	var mach string
	const plus string = "+"
	const minus string = "-"
	const mult string = "*"
	const div string = "/"
	fmt.Println("Что с ними сделать?")
	fmt.Println(plus, " сложить,", minus, " вычесть,", mult, " умножить,", div, " разделить")
	fmt.Scanf("%s\n",&mach)
	if mach == plus || mach == minus || mach == mult || mach == div {
		return mach
	} else {
		return "?"
	}
}


func main() {
	//testZifry()
	callcul()

}
