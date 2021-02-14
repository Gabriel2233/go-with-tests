package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test number one"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test number one"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		want := ErrNotFound

		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "test number one"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "test number one"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test number one")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "test number one"
		dictionary := Dictionary{word: definition}
		newDefinition := "updated test number one"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "test number one"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test number one"}

	dictionary.Delete("test")

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
