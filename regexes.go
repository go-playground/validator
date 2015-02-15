package validator

import "regexp"

const (
	alphaRegexString        string = "^[a-zA-Z]+$"
	alphaNumericRegexString string = "^[a-zA-Z0-9]+$"
)

var (
	alphaRegex        = regexp.MustCompile(alphaRegexString)
	alphaNumericRegex = regexp.MustCompile(alphaNumericRegexString)
)
