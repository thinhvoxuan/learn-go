package hello

const testVersion = 2
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Helloworld function
func Helloworld(name string, lang string) string {
	if len(name) == 0 {
		name = "World"
	}
	return greetingPrefix(lang) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		return frenchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
}
