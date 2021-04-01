package main

// TODO: UTF-8 size for code points U+0000 to U+007F
const MAX_SIZE = 26

type Trie interface {
	Insert()
	Search() bool
	Delete()
}

type Registry struct {
	table   [MAX_SIZE]*Node
	isEmpty bool
}

type Node struct {
	registry *Registry
	isWord   bool
}

type SingleTrie struct {
	root *Node
}

func Init() *SingleTrie {
	new := &SingleTrie{root: New()}
	return new
}

func New() *Node {
	return &Node{registry: &Registry{isEmpty: true}}
}

func (t *SingleTrie) Insert(word string) {
	current := t.root
	for _, n := range word {
		index := n - 'a'
		if current.registry.table[index] == nil {
			current.registry.isEmpty = false
			current.registry.table[index] = New()
		}
		current = current.registry.table[index]
	}
	current.isWord = true
}

func (t *SingleTrie) Search(word string) bool {
	current := t.root
	for _, n := range word {
		index := n - 'a'
		if current.registry.table[index] == nil {
			return false
		}
		current = current.registry.table[index]
	}
	return current.isWord
}

func (t *SingleTrie) Delete(word string) {
	current := t.root
	for _, n := range word {
		index := n - 'a'
		if current.registry.table[index] != nil {
			current = current.registry.table[index]
		}
	}
	current.isWord = false
}
