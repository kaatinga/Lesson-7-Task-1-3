package main

import (
	"fmt"
	"log"
	"time"
)

// Переменная для хранения времени когда нужно остановиться для выполнения задачи 1 домашнего задания
var whenToStop time.Time

// Переменная для хранения результата ввода для выполнения задачи 3 домашнего задания
var input string

func main() {

	whenToStop = time.Now().Add(+time.Second * 15)
	
	fmt.Println("Type 'exit' and press 'Enter' in order to stop revolving.")
	go spinner(250 * time.Millisecond)
	for {
		fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("Enough as you typed 'Exit'!")
			break
		}
	}
}

func spinner(delay time.Duration) {
	fmt.Println("I am revolving...")
	for r := 0; ; r++ {
		fmt.Print(r,"... ")
		time.Sleep(delay)
		if time.Now().After(whenToStop) {
			log.Fatalln("Enough as time is depleted!")
		}
	}
}
