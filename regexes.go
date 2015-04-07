package validator

import "regexp"

const (
	alphaRegexString        = "^[a-zA-Z]+$"
	alphaNumericRegexString = "^[a-zA-Z0-9]+$"
	numericRegexString      = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString       = "^[0-9]+$"
	hexadecimalRegexString  = "^[0-9a-fA-F]+$"
	hexcolorRegexString     = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	rgbRegexString          = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	rgbaRegexString         = "^rgba\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*((0.[1-9]*)|[01])\\s*\\)$"
	hslRegexString          = "^hsl\\(\\s*(0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*((0|[1-9]\\d?|100)%)\\s*,\\s*((0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegexString         = "^hsla\\(\\s*(0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*((0|[1-9]\\d?|100)%)\\s*,\\s*((0|[1-9]\\d?|100)%)\\s*,\\s*((0.[1-9]*)|[01])\\s*\\)$"
	emailRegexString        = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	// urlRegexString          = `^((ftp|http|https):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|((www\.)?)?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?)|localhost)(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`
	// urlRegexString = "^(?:(?:https?|ftp):\\/\\/)(?:\\S+(?::\\S*)?@)?(?:(?!10(?:\\.\\d{1,3}){3})(?!127(?:\\.\\d{1,3}){3})(?!169\\.254(?:\\.\\d{1,3}){2})(?!192\\.168(?:\\.\\d{1,3}){2})(?!172\\.(?:1[6-9]|2\\d|3[0-1])(?:\\.\\d{1,3}){2})(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)(?:\\.(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)*(?:\\.(?:[a-z\u00a1-\uffff]{2,})))(?::\\d{2,5})?(?:\\/[^\\s]*)?$"
)

var (
	alphaRegex        = regexp.MustCompile(alphaRegexString)
	alphaNumericRegex = regexp.MustCompile(alphaNumericRegexString)
	numericRegex      = regexp.MustCompile(numericRegexString)
	numberRegex       = regexp.MustCompile(numberRegexString)
	hexadecimalRegex  = regexp.MustCompile(hexadecimalRegexString)
	hexcolorRegex     = regexp.MustCompile(hexcolorRegexString)
	rgbRegex          = regexp.MustCompile(rgbRegexString)
	rgbaRegex         = regexp.MustCompile(rgbaRegexString)
	hslRegex          = regexp.MustCompile(hslRegexString)
	hslaRegex         = regexp.MustCompile(hslaRegexString)
	emailRegex        = regexp.MustCompile(emailRegexString)
)

func matchesRegex(regex *regexp.Regexp, field interface{}) bool {
	fieldAsString := field.(string) //this will panic, we do not explicit panic
	return regex.MatchString(fieldAsString)
}
