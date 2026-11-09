package main

const (
	englishHelloPrefix = "Hello, "
	italian            = "IT"
	italianHelloPrefix = "Ciao, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	switch lang {
	case italian:
		return italianHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}

}

func main() {
}
