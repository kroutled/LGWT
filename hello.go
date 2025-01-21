package main

import (
    "fmt"
)

const englishGreet = "Hello, "
const spanishGreet = "Hola, "
const frenchGreet = "Bonjour, "
const sp = "Spanish"
const en = "English"
const fr = "French"

func Hello(name, language string) string{
    greet := englishGreet + name
    if name == "" {
        greet = "Hello World"
    }
    if language == sp {
        greet = spanishGreet + name
    }
    if language == fr {
        greet = frenchGreet + name
    }
    return greet
}

func main(){
    fmt.Println(Hello("Kyle", ""))
}
