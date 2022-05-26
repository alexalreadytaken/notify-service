package utils

import "regexp"

var numberRegex = regexp.MustCompile(`\d+`)

func IsDigit(str string) bool {
	return numberRegex.MatchString(str)
}
