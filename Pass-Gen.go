package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		length int
		number int
	)
	fmt.Println("Enter Length Of The Password: ")
	fmt.Scanln(&length)
	fmt.Println("Enter Number Of Password To Generate:  ")
	fmt.Scanln(&number)
	for i := 0; i < number; i++ {
		fmt.Println(generatePassword(length))
	}
}

func generatePassword(length int) string {
	var (
		alphabets = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
		numbers   = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
		special   = []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "_", "-", "+", "=", "~", "`", "|", "}", "{", "[", "]", ":", ";", ",", ".", "?", "/"}
		password  = ""
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		switch rand.Intn(3) {
		case 0:
			password += alphabets[rand.Intn(len(alphabets))]
		case 1:
			password += numbers[rand.Intn(len(numbers))]
		case 2:
			password += special[rand.Intn(len(special))]
		}
	}
	return password
}
