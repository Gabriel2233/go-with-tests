package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTmpFile(t testing.TB, initialData string) (*os.File, func()) {
	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("error while creating temporary file, %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[
			{"Name": "Gabriel", "Wins": 33},
			{"Name": "Rafael", "Wins": 22}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()
		want := []Player{
			{"Gabriel", 33},
			{"Rafael", 22},
		}
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[
			{"Name": "Gabriel", "Wins": 33},
			{"Name": "Rafael", "Wins": 22}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Gabriel")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("record wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[
			{"Name": "Gabriel", "Wins": 33},
			{"Name": "Rafael", "Wins": 22}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Gabriel")

		got := store.GetPlayerScore("Gabriel")
		want := 34

		assertScoreEquals(t, got, want)
	})

	t.Run("record wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, `[
			{"Name": "Gabriel", "Wins": 33},
			{"Name": "Rafael", "Wins": 22}
		]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pedro")

		got := store.GetPlayerScore("Pedro")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
