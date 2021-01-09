package main

import (
	"fmt"
)

func callcul() {
	var arg1, arg2, result int64
	var mach string
	var err int
	const plus string = "+"
	const minus string = "-"
	const mult string = "*"
	const div string = "/"
	fmt.Println("Это калькулятор")
	fmt.Print("-------------------", "\n")
	arg1, arg2, err = inparg()
	if err > 0 {
		 fmt.Println("Числа не должны равняться нулю, операция прекращена")
			fmt.Println("До новых встреч!","\n","Ждем с нетерпением!")
			return
	}
	mach = inpmachen()
	switch mach {
		case plus: result = arg1 + arg2
		case minus: result = arg1 - arg2
		case mult: result = arg1 * arg2
		case div: {
			if arg2 == 0 && mach == div {
				fmt.Println("Деление на ноль невозможно")
				return
			} else {
				result = arg1 / arg2
			}
		}
		default:
			{ fmt.Println("Команда не распознана, операция прекращена")
			  fmt.Println("До новых встреч!","\n","Ждем с нетерпением!")
				return
			}
	}
	fmt.Print(arg1, " ", mach, " ", arg2, " = ", result,"\n")
	fmt.Print("-------------------","\n")
	return
}

func inparg() (int64, int64, int) {
	var arg1, arg2 int64
	var i int
	fmt.Println("Введите первое число")
	for i:= 0; i<5; i++ {
		fmt.Scanf("%d\n",&arg1)
		if arg1 == 0 {
			fmt.Println("Число не должно равно нулю")
		} else {
			break
		}
	}
	if i == 5 {
		return 0, 0, 1
	}
	fmt.Println(arg1)
	fmt.Println("Введите второе число")
	fmt.Scanf("%d\n",&arg2)
	for i:= 0; i<5; i++ {
		fmt.Scanf("%d\n",&arg1)
		if arg1 == 0 {
			fmt.Println("Число не должно равно нулю")
		} else {
			break
		}
	}
	if i == 5 {
		return 0, 0, 1
	}
	return arg1, arg2, 0
}

func inpmachen() (string) {
	var mach string
	const plus string = "+"
	const minus string = "-"
	const mult string = "*"
	const div string = "/"
	fmt.Println("Что с ними сделать?")
	fmt.Println(plus, " сложить,", minus, " вычесть,", mult, " умножить,", div, " разделить")
	for i:= 0; i < 5; i++ {
		fmt.Scanf("%s\n",&mach)
		if mach == plus || mach == minus || mach == mult || mach == div {
			return mach
		} else {
			fmt.Println("Команда не распознана, введите новую")
			fmt.Println(plus, " сложить,", minus, " вычесть,", mult, " умножить,", div, " разделить")
		}
	}
	fmt.Println("Попытки ввода команды исчерпаны", "\n")
	return "?"
}


func main() {

	callcul()
}
