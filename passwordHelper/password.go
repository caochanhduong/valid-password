package passwordHelper

import (
	"regexp"
)

func CheckValidPassword(pw string) (isValid bool, warningMessage string, err error) {
	matchUpper, err := regexp.MatchString(UpperCasePattern, pw)
	if !matchUpper {
		warningMessage += "Password must have at least one uppercase character.\n"
	}
	matchLower, err := regexp.MatchString(LowerCasePattern, pw)
	if !matchLower {
		warningMessage += "Password must have at least one lowercase character.\n"
	}
	matchNumber, err := regexp.MatchString(NumberPattern, pw)
	if !matchNumber {
		warningMessage += "Password must have at least one number.\n"
	}
	matchSpecial, err := regexp.MatchString(SpecialPattern, pw)
	if matchSpecial {
		warningMessage += "Password must have at least one special character."
	}
	isValid = matchNumber && matchUpper && matchLower && !matchSpecial
	return
}
