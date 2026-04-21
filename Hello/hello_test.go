package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"

		assertCorrectMessage(t, got , want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want:= "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("This is a failing test sample", func(t *testing.T){
		got := Hello("Asamoah","Twi" )
		want:= "Akwaaba, Asamoah"
		assertCorrectMessage(t, got, want)

	})
}

//This function is to reduce my failing error checks
//for helper functions, it's a good idea too accept a `testing.TB` which is an interface
//that *testing.T and *testing.B both satisfy
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() //This is needed to tell the test suite that this method is a helper
	if got != want {
		t.Errorf("got %q want %q", got , want )
	}
}