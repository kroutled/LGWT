package main

import(
    "testing"
)

func TestHello(t *testing.T){
    t.Run("Greeting person by name", func(t *testing.T){
        got := Hello("Kyle", "")
        want := "Hello, Kyle"
        assertCorrectMessage(t, got, want)
    })
    t.Run("Testing greeting with no name", func(t *testing.T){
        got := Hello("", "")
        want := "Hello World"
        assertCorrectMessage(t, got, want)
    })
    t.Run("Testing Spanish input", func(t *testing.T){
        got := Hello("Kail", "Spanish")
        want := "Hola, Kail"
        assertCorrectMessage(t, got, want)
    })
    t.Run("Testing French Input", func(t *testing.T){
        got := Hello("Kyle", "French")
        want := "Bonjour, Kyle"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string){
    t.Helper()
    if got != want {
        t.Errorf("got: %q want: %q", got, want)
    }
}
