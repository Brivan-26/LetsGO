package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"first word": "Definition of the first word"}
	t.Run("Known word", func(t *testing.T) {
		assertWordExists(t, dictionary, "first word", "Definition of the first word")
	})

	t.Run("Unkown word", func(t *testing.T) {
		_, got := dictionary.Search("unkown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		newKey := "first word"
		newValue := "first definition"

		err := dictionary.Add(newKey, newValue)

		assertError(t, err, nil)

		assertWordExists(t, dictionary, newKey, newValue)
	})
	t.Run("Word exists", func(t *testing.T) {
		dictionary := Dictionary{"bananas": "bananas definition"}
		err := dictionary.Add("bananas", "What ever...")
		
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word to be updated", func(t *testing.T) {
		dictionary := Dictionary{"first word": "first definition"}
		err := dictionary.Update("first word", "updated first definition")

		assertError(t, err, nil)
		assertWordExists(t, dictionary, "first word", "updated first definition")
	})
	t.Run("Unexisting word to be updated", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("first word", "updated")

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"first word": "first definition"}
	dictionary.Delete("first word")
	_, got := dictionary.Search("first word")
	assertError(t, got, ErrNotFound)
}

func assertWordExists(t testing.TB, dic Dictionary, key, want string) {
	t.Helper()
	got, _ := dic.Search(key)
	if got != want {
		t.Errorf("Got: %q, Expected: %q", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got error: %q, Exptected: %q", got, want)
	}
}