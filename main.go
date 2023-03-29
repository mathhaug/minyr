package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	yr "/minyr/yr/yr"
)

func main() {
	src, err := os.Open("/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	log.Println(src)

	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("Konverterer alle maalingene gitt i grader Celsius til grader Fahrenheit.")
		} else {
			fmt.Println("Vennligst velg convert, average eller exit")
		}
	}
}

