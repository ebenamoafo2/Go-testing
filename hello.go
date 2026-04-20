package main

import (
	"fmt"
)

const (

 french = "French"
 spanish = "Spanish"
 twi = "Twi"
	
 englishHelloPrefix ="Hello, "
 spanishHelloPrefix = "Hola, "
 frenchHelloPrefix = "Bonjour, "
 twiHelloPrefix = "Akwaaba, "
)


func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french: 
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case twi: 
		prefix = twiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world",""))
}
