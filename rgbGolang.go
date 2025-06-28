package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const colorRed = "\033[0;31m"
const colorNone = "\033[0m"

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func main() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	var Blue = "\033[34m"
	var Magenta = "\033[35m"
	var Cyan = "\033[36m"
	var Gray = "\033[37m"
	var White = "\033[97m"
	Use(Reset, Red, Green, Yellow, Blue, Magenta, Cyan, Gray, White) // чтобы не было назойливой ошибки что переменные не используются

	// rgb
	fmt.Fprintf(os.Stdout, "Red: \033[0;31m %s None: \033[0m %s", "red string", "colorless string")
	fmt.Fprintf(os.Stdout, "Red: %s %s None: %s %s", colorRed, "red string", colorNone, "colorless string")
	println()

	// brightness
	println("\033[33;1m This is Bright Yellow \033[0m")
	println("\033[33m This is Yellow \033[0m")
	println("\033[33;2m This is I dont know what Yellow \033[0m")
	println("\033[33;0m This is definately not Yellow \033[0m")

	var mode int
	fmt.Print("Which mode of rgb string do you prefer(possible answers are 1, 2, Enter for default(2)): ")
	_, err := fmt.Scanln(&mode)
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// }
	if mode == 0 {
		mode = 2
	}
	var word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Which word do you want to make an rgb(Enter for default Hello world): ")
	word, _ = reader.ReadString('\n')
	// _, err = fmt.Scanln(&word)
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// }
	if word == "\n" {
		word = "Hello world"
	}
	switch mode {
		case 1:
			for i := 0; ; i++ {
				colorCode := 30 + (i % 8)
				fmt.Printf("\033[%vm %s \033[0m\n", colorCode, word)
				time.Sleep(100000000)
			}
		case 2:
			for i := 0; ; i++ {
				for index, j := range word {
					colorCode := 30 + ((i+index) % 8)
					fmt.Printf("\033[%vm%c\033[0m", colorCode, j)
				}
				fmt.Println()
				// fmt.Printf("\033[%vm Hello world \033[0m\n", colorCode)
				time.Sleep(100000000)
			}
		default:
			fmt.Println("Invalid mode, you should have chosen 1 or 2, not ", mode, " You think you clever boii???")
	}
}
