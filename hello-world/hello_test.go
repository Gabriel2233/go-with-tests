package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gabriel", "")
		want := "Hello Gabriel"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is passed as a parameter", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello' in Spanish", func(t *testing.T) {
		got := Hello("Gabriel", "Spanish")
		want := "Hola Gabriel"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello' in French", func(t *testing.T) {
		got := Hello("Gabriel", "French")
		want := "Bounjour Gabriel"

		assertCorrectMessage(t, got, want)
	})
}
