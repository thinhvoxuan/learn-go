package acronym

import "strings"
import "unicode"

const testVersion = 3

// IsPunctIsSpace return char is Punc or Space
func IsPunctIsSpace(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r)
}

//Abbreviate convert string to string
func Abbreviate(str string) string {
	result := ""
	arrayString := strings.FieldsFunc(str, IsPunctIsSpace)
	for _, str := range arrayString {
		result = result + strings.ToUpper(string(str[0]))
	}
	return result
}
