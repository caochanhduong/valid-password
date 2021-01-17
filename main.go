package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

const (
	UpperCasePattern = "[A-Z]+"
	LowerCasePattern = "[a-z]+"
	NumberPattern    = "[0-9]+"
)

func checkValidPassword(pw string) bool {
	matchUpper, _ := regexp.MatchString(UpperCasePattern, pw)
	matchLower, _ := regexp.MatchString(LowerCasePattern, pw)
	matchNumber, _ := regexp.MatchString(NumberPattern, pw)
	return matchNumber && matchUpper && matchLower
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please specify your password")
		os.Exit(1)
	}
	password := args[0]
	if checkValidPassword(password) {
		fmt.Println("OK! Your password is valid.")
	} else {
		fmt.Println("Sorry! Your password is invalid.")
	}
}
