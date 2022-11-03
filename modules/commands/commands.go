package commands

import (
	"fmt"
	"strconv"
)

func ShowAllAvailableCommands() int {
	var command string

	fmt.Println("=================================")
	fmt.Println("List of available commands :")

	fmt.Print("\n")
	fmt.Print(" [1] 🔍 Search a Book by Name \n")
	fmt.Print(" [2] 📚 Borrow a Book \n")
	fmt.Print(" [3] 📘 Return Borrowed Book \n")
	fmt.Print(" [4] 📙 Show All Borrowed Book \n")
	fmt.Print(" [5] 📕 Show All Overdue Book \n")
	fmt.Print("\n")
	fmt.Println("=================================")

	fmt.Print("\n", "$ ")
	fmt.Scanln(&command)

	value, err := strconv.Atoi(command)
	if err != nil {
		panic(err)
	}

	return value
}

func ChooseBookToBorrow() string {
	return InputTemplate("Input book ISBN to borrow $ ")
}

func ChooseBookToReturn() string {
	return InputTemplate("Input book ISBN to return $ ")
}

func SearchForBook() string {
	return InputTemplate("Input book name $ ")
}
