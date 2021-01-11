package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	mach string

	ErrInvalidNumber   = errors.New("invalid number error")
	ErrInvalidOperator = errors.New("invalid operator error")
	ErrDivideZero      = errors.New("divide on zero error")
)

const (
	plus  string = "+"
	minus string = "-"
	mult  string = "*"
	div   string = "/"
)

func callcul() (float64, float64, float64, error) {
	var err error
	arg1, err := inparg(1)
	if err != nil {
		return 0, 0, 0, err
	}
	mach = inpmachen()

	arg2, err := inparg(2)
	if err != nil {
		return 0, 0, 0, err
	}
	result, err := operation(arg1, arg2)
	return result, arg1, arg2, err
}

func inparg(num int) (float64, error) {
	var (
		res       float64
		intString string
	)

	if num == 1 {
		fmt.Println("Введите первое число")
	} else {
		fmt.Println("Введите второе число")
	}
	fmt.Scanf(
		"%s\n",
		&intString,
	)
	res, err := strconv.ParseFloat(intString, 64)
	if err != nil {
		return 0, ErrInvalidNumber
	}
	return res, nil
}

func inpmachen() string {
	var err int
	err = 1
	for err > 0 {
		fmt.Println("-----------------------------")
		fmt.Println("Выберите действие:")
		fmt.Println(plus, " сложить")
		fmt.Println(minus, " вычесть")
		fmt.Println(mult, " умножить")
		fmt.Println(div, " разделить")
		fmt.Println("-----------------------------")
		fmt.Scanf(
			"%s/n",
			&mach,
		)
		if mach == plus || mach == minus || mach == mult || mach == div {
			err = 0
			switch mach {
			case plus:
				fmt.Println(" сложить")
			case minus:
				fmt.Println(" вычесть")
			case mult:
				fmt.Println(" умножить")
			case div:
				fmt.Println(" разделить")
			}
		} else {
			fmt.Println("*** Введите правильный символ!!! ***")
		}
	}
	fmt.Println("-----------------------------")
	return mach
}

func operation(arg1, arg2 float64) (float64, error) {
	var res float64

	switch mach {
	case plus:
		res = arg1 + arg2
	case minus:
		res = arg1 - arg2
	case mult:
		res = arg1 * arg2
	case div:
		{
			if arg2 == 0 && mach == div {
				return 0, ErrDivideZero
			} else {
				res = arg1 / arg2 // деление возможно
			}
		}
	}
	return res, nil
}

func printresult(result, arg1, arg2 float64, err error) {

	switch err {
	case nil:
		{
			fmt.Println("-----------------------------")
			fmt.Print("Результат:", "\n")
			fmt.Printf("%.5f", arg1)
			fmt.Print(" ", mach, " ")
			fmt.Printf("%.5f", arg2)
			fmt.Print(" = ")
			fmt.Printf("%.5f", result)
			fmt.Print("\n")
		}
	case ErrInvalidNumber:
		{
			fmt.Println("\n", "*** К сожалению, вы ввели не число!!")
			fmt.Println("*** Операция прекращена ***")
		}
	case ErrInvalidOperator:
		{
			fmt.Println("\n", "*** К сожалению, вы ввели неверный символ ***")
			fmt.Println("*** Операция прекращена ***")
		}
	case ErrDivideZero:
		{
			fmt.Println("-----------------------------")
			fmt.Println("*** Деление на ноль невозможно ***")
			fmt.Println("*** Операция прекращена ***")
		}
	}
	fmt.Print("-----------------------------", "\n", "\n")
}

func main() {
	var command string
	mach = "?"

	command = "Y"
	for strings.ToLower(command) == "y" || strings.ToLower(command) == "н" {
		fmt.Println("-----------------------------")
		fmt.Println("|       Калькулятор          |")
		fmt.Println("|  Считать, не пересчитать!  |")
		fmt.Println("|                            |")
		fmt.Println("|   (c) jiliaevyp@gmail.com  |")
		fmt.Println("-----------------------------")
		//result, arg1, arg2, err :=
		printresult(callcul())
		fmt.Println("Продолжить? (Y)")
		fmt.Println("Закончить? (Enter)")
		command = ""
		fmt.Scanf(
			"%s\n",
			&command,
		)

	}
	fmt.Println("Рад был с Вами пработать!")
	fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
}
