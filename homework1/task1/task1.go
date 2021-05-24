package main

import (
	"fmt"
	"runtime"
)

// Вывод паники в консольConsole-log fake Panic
func getPanic() {
	panic("::Panic in getPanic::")
}

func callMakePanic() {
	// Вызов отложенного выполнения обработчика паники
	defer panicHandler()
	defer fmt.Println("defer 1 makePanic")

	// Вложенный вызов получения паники
	getPanic()
	fmt.Println("after makePanic")
}

// Обработчик паники с выводом в консоль с указанием байтов
func panicHandler() {
	if v := recover(); v != nil {
		buff := make([]byte, 1024)
		runtime.Stack(buff, false)
		fmt.Printf("panic with value: %v, %s\n", v, buff)
	}
}

func main() {
	callMakePanic()

	// Подтверждение, что после вызванной паники в программе
	// выполнение не будет прервано
	fmt.Println("")
	fmt.Println("happy app filnalization")
}
