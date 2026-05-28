package tries

import "testing"

func TestPrefixTreeExampleFlow(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Fatal(`Search("apple") = false, want true`)
	}
	if trie.Search("app") {
		t.Fatal(`Search("app") = true, want false before inserting app`)
	}
	if !trie.StartsWith("app") {
		t.Fatal(`StartsWith("app") = false, want true`)
	}

	trie.Insert("app")

	if !trie.Search("app") {
		t.Fatal(`Search("app") = false, want true after inserting app`)
	}
}

func TestPrefixTreeSharedPrefixes(t *testing.T) {
	trie := Constructor()

	for _, word := range []string{"car", "card", "care", "dog"} {
		trie.Insert(word)
	}

	for _, word := range []string{"car", "card", "care", "dog"} {
		if !trie.Search(word) {
			t.Fatalf("Search(%q) = false, want true", word)
		}
	}

	for _, prefix := range []string{"c", "ca", "car", "do"} {
		if !trie.StartsWith(prefix) {
			t.Fatalf("StartsWith(%q) = false, want true", prefix)
		}
	}
}

func TestPrefixTreeDoesNotTreatPrefixesAsWords(t *testing.T) {
	trie := Constructor()

	trie.Insert("there")
	trie.Insert("their")

	for _, word := range []string{"t", "th", "the", "ther", "thei"} {
		if trie.Search(word) {
			t.Fatalf("Search(%q) = true, want false because prefix was not inserted", word)
		}
	}

	for _, prefix := range []string{"t", "th", "the", "ther", "thei"} {
		if !trie.StartsWith(prefix) {
			t.Fatalf("StartsWith(%q) = false, want true", prefix)
		}
	}
}

func TestPrefixTreeMissingBranches(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")
	trie.Insert("banana")

	for _, word := range []string{"apply", "ban", "band", "cat"} {
		if trie.Search(word) {
			t.Fatalf("Search(%q) = true, want false", word)
		}
	}

	for _, prefix := range []string{"apz", "band", "cat"} {
		if trie.StartsWith(prefix) {
			t.Fatalf("StartsWith(%q) = true, want false", prefix)
		}
	}
}

func TestPrefixTreeRepeatedInsert(t *testing.T) {
	trie := Constructor()

	trie.Insert("same")
	trie.Insert("same")

	if !trie.Search("same") {
		t.Fatal(`Search("same") = false, want true after repeated inserts`)
	}
	if !trie.StartsWith("sam") {
		t.Fatal(`StartsWith("sam") = false, want true`)
	}
}
