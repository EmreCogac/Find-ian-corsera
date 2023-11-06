package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Println("give me a word :")
	fmt.Scan(&word)
	Findian(word)
}

func Findian(word string) {
	var UpWord = strings.ToUpper(word)
	wordString := strings.Split(UpWord, "")
	var haveI = 0
	var haveA = 0
	var haveN = 0
	for i := 0; i < len(wordString); i++ {
		if wordString[i] == "I" {
			haveI++
		} else if wordString[i] == "A" {
			haveA++
		} else if wordString[i] == "N" {
			haveN++
		}
	}
	if haveI != 0 && haveA != 0 && haveN != 0 {
		fmt.Println("have i,a,n")
	} else {
		fmt.Println("i , a and n not founded")
	}
}
