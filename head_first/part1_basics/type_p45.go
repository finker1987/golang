package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello")
	//Новая строка \n
	fmt.Print("hello\n")
	//Табуляция \t
	fmt.Println("hello\thello")
	//Двойная кавычка \"
	fmt.Println("hello\"")
	//Обратный слеш \\
	fmt.Println("hello\\hello")
	//Одинарные кавычки для типа byte
	var rawBinary byte = '\x27'
	fmt.Println(rawBinary)
	//Контактенация строк
	helloWorld := "Привет МИР!"
	andGoodmorning := helloWorld + "и доброе утро"
	fmt.Println(andGoodmorning)
	//функция подсчета длинны строки в байтах (len)
	lenStr := len(helloWorld)
	fmt.Println(lenStr)
	lenRune := utf8.RuneCountInString(helloWorld)
	fmt.Println(lenRune)

}
