package main

import (
    "fmt"
)

const englishGreet = "Hello, "

func Hello(name string) string{
    greet := englishGreet + name
    if name == "" {
        greet = "Hello World"
    }
    return greet
}

func main(){
    fmt.Println(Hello("Kyle"))
}
