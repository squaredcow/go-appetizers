package main

import "fmt"

func main() {
	trie := Init()
	buildTrie := []string{"oregon", "oreo", "ore", "organ", "aragorn"}
	for _, n := range buildTrie {
		trie.Insert(n)
	}
	removeTrie := []string{"oregon", "oreo", "ore", "organ", "aragorn"}
	for _, n := range removeTrie {

		fmt.Println("Is present?", trie.Search(n))
		fmt.Println("Deleting", n)
		trie.Delete(n)
		fmt.Println("Is present, now?", trie.Search(n))
	}
}
