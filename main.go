package main

import (
	"flag"
	"fmt"
	"os"

	helper "github.com/caochanhduong/valid-password/passwordHelper"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please specify your password")
		os.Exit(1)
	}
	password := args[0]
	isValid, message, err := helper.CheckValidPassword(password)
	if err != nil {
		fmt.Println("Internal service error.")
	} else if isValid {
		fmt.Println("OK! Your password is valid.")
	} else {
		fmt.Println("Sorry! Your password is invalid.")
		fmt.Println(message)
	}
}
