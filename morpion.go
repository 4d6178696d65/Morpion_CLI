package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	size     = 9
	actionP1 = "X"
	actionP2 = "0"
)

var (
	tab = [size]string{
		"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9"}
	p1 = true //player 1 starts
)

func main() {
	run()
}

func run() {
	var choice int
	for true {
		display()
		choice = input()
		inputPlayer(choice)
		if win() {
			display()
			fmt.Println(playerName(), "a gagné la partie")
			os.Exit(0) //exit the game
		} else if loose() {
			fmt.Println("Match nul !")
			os.Exit(0)
		}
		p1 = !p1 //change player
	}
}

func display() {
	for i := 0; i < size; i++ {
		fmt.Print(" ", tab[i])
		if (i+1)%3 == 0 { // every 3 numbers go to new line
			fmt.Println()
		}
	}
}

func playerName() string {

	if p1 {
		return "Joueur 1"
	}
	return "Joueur 2"

}

func input() int {
	var (
		inputOK     = false
		inputNumber = 0
		err         error
		scanner     = bufio.NewScanner(os.Stdin)
	)

	for inputOK == false {
		fmt.Print(playerName(), " entrez une valeur entre 1 et ", size, ": ")
		scanner.Scan()                                  // analyze input data
		inputNumber, err = strconv.Atoi(scanner.Text()) // keep the data (int) in memory
		if err != nil {
			fmt.Println("Vous devez entrer uniquement un nombre")
		} else if inputNumber < 1 || inputNumber > size {
			fmt.Println("Vous devez entrer un nombre entre 1 et", size)
		} else if tab[inputNumber-1] == actionP1 || tab[inputNumber-1] == actionP2 {
			fmt.Println("Cet emplacement est déjà occupé")
		} else {
			inputOK = true
		}

	}

	return inputNumber - 1 // index is from 0
}

func inputPlayer(inputNumber int) { // assign player action to input number
	if p1 {
		tab[inputNumber] = actionP1
	} else {
		tab[inputNumber] = actionP2

	}
}

func win() bool {
	resultTab := [][size]bool{
		{
			true, true, true,
			false, false, false,
			false, false, false},
		{
			false, false, false,
			true, true, true,
			false, false, false},
		{
			false, false, true,
			false, false, true,
			false, false, true},
		{
			false, true, false,
			false, true, false,
			false, true, false},
		{
			true, false, false,
			false, true, false,
			false, false, true},
		{
			false, false, true,
			false, true, false,
			true, false, false},
		{
			false, false, false,
			false, false, false,
			true, true, true},
		{
			true, false, false,
			true, false, false,
			true, false, false}}

	var copyTab [size]bool // copy of the original tab with bool data inside
	for index, valeur := range tab {
		if p1 && valeur == actionP1 {
			copyTab[index] = true
		} else if !p1 && valeur == actionP2 {
			copyTab[index] = true
		}
	}

	count := 0
	for _, resultTab := range resultTab {
		for i := 0; i < len(copyTab); i++ {
			if copyTab[i] == true && copyTab[i] == resultTab[i] { // analyze copy tab with result tab
				count++
				if count == 3 {
					return true
				}
			}
		}
		count = 0
	}
	return false
}

func loose() bool {
	o := 0
	for _, valeur := range tab {
		if valeur == actionP1 || valeur == actionP2 { // count the number of players actions
			o++
		}
	}

	return o == len(tab) // until tab lenght
}
