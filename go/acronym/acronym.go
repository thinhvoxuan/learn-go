package acronym

import "strings"
import "unicode"

const testVersion = 2

// IsPunctIsSpace return char is Punc or Space
func IsPunctIsSpace(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r)
}

//Abbreviate convert string to string
func Abbreviate(str string) string {
	if pos := strings.Index(str, ":"); pos != -1 {
		return str[:pos]
	}
	result := ""
	arrayString := strings.FieldsFunc(str, IsPunctIsSpace)
	for _, str := range arrayString {
		result = result + strings.ToUpper(string(str[0]))
		arr2 := strings.FieldsFunc(str[1:], unicode.IsLower)
		for _, upperCase := range arr2 {
			result = result + string(upperCase)
		}
	}
	return result
}
