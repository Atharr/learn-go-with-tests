package main

import "fmt"

type Msg struct {
	prefix      string
	defaultName string
	suffix      string
}

var msg = map[string]Msg{
	"English": Msg{
		prefix:      "Hello, ",
		defaultName: "World",
		suffix:      "!",
	},
	"Portuguese": Msg{
		prefix:      "Olá, ",
		defaultName: "Mundo",
		suffix:      "!",
	},
	"Spanish": Msg{
		prefix:      "¡Hola, ",
		defaultName: "Mundo",
		suffix:      "!",
	},
}

func Hello(name, language string) string {
	_, pk := msg[language]
	if language == "" || !pk {
		language = "English"
	}
	if name == "" {
		name = msg[language].defaultName
	}
	return msg[language].prefix + name + msg[language].suffix
}

func main() {
	fmt.Println(Hello("Marcello", "English"))
}
