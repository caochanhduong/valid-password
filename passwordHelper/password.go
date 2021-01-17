package passwordHelper

import (
	"regexp"
)

func CheckValidPassword(pw string) (isValid bool, warningMessage string) {
	matchUpper, _ := regexp.MatchString(UpperCasePattern, pw)
	if !matchUpper {
		warningMessage += "Password must have at least one uppercase character.\n"
	}
	matchLower, _ := regexp.MatchString(LowerCasePattern, pw)
	if !matchLower {
		warningMessage += "Password must have at least one lowercase character.\n"
	}
	matchNumber, _ := regexp.MatchString(NumberPattern, pw)
	if !matchNumber {
		warningMessage += "Password must have at least one number.\n"
	}
	matchSpecial, _ := regexp.MatchString(SpecialPattern, pw)
	if matchSpecial {
		warningMessage += "Password must have at least one special character.\n"
	}
	isValid = matchNumber && matchUpper && matchLower && !matchSpecial
	return
}
