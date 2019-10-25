package main

import "fmt"


var greetings = map[string]string {
	"English" : "Hello, ",
	"Spanish" : "Hola, ",
	"French"	: "Bonjour, ",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetings[language] + name
}

func main(){
	fmt.Println(Hello("Hattan","English"))
}