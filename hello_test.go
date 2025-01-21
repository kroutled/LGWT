package main

import(
    "testing"
)

func TestHello(t *testing.T){
    t.Run("Greeting person by name", func(t *testing.T){
        got := Hello("Kyle")
        want := "Hello, Kyle"

        if got != want {
            t.Errorf("got: %q want: %q", got, want)
        }
    })
    t.Run("Testing greeting with no name", func(t *testing.T){
        got := Hello("")
        want := "Hello World"

        if got != want {
            t.Errorf("got: %q want: %q", got, want)
        }
    }) 
}
