package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
)

const (
	englishPrefix = "Hello "
	spanishPrefix = "Hola "
	frenchPrefix  = "Bounjour "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Gabriel", ""))
}
