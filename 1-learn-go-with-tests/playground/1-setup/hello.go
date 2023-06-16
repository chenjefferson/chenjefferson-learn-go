package main

import (
	"fmt"
)

const (
	english  = "English"
	spanish  = "Spanish"
	japanese = "Japanese"
	chinese  = "Chinese"
)
const (
	EnglishHelloPrefix  = "Hello, "
	SpanishHelloPrefix  = "Hola, "
	JapaneseHelloPrefix = "Ohaiyoo, "
	ChineseHelloPrefix  = "Nihao, "
)

func Hello(name string, language string) string {
	prefix := EnglishHelloPrefix

	if len(name) == 0 {
		name = "world"
	}

	switch language {
	case english:
		prefix = EnglishHelloPrefix
	case spanish:
		prefix = SpanishHelloPrefix
	case japanese:
		prefix = JapaneseHelloPrefix
	case chinese:
		prefix = ChineseHelloPrefix
	default:
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
