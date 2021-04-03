package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		want *SingleTrie
	}{
		{"test_init", &SingleTrie{root: New()}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Node
	}{
		{"new_empty_node", &Node{registry: &Registry{isEmpty: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleTrie_Insert(t *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		word string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{"single_insert", fields{root: New()}, args{word: "single"}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SingleTrie{
				root: tt.fields.root,
			}
			last := tr.Insert(tt.args.word)
			if !last.isWord {
				t.Errorf("SingleTrie.Insert() = %v, want %v", last.isWord, true)
			}
			if r := tr.Search(tt.args.word); !r {
				t.Errorf("SingleTrie.Insert() = %v, want %v", r, true)
			}
		})
	}
}

func TestSingleTrie_Search(t *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		word string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"word_is_present_after_insertion", fields{root: New()}, args{word: "single"}, true},
		{"word_not_present", fields{root: New()}, args{word: "complex"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SingleTrie{
				root: tt.fields.root,
			}

			if strings.Compare(tt.name, "word_is_present_after_insertion") == 0 {
				tr.Insert(tt.args.word)
			}

			if got := tr.Search(tt.args.word); got != tt.want {
				t.Errorf("SingleTrie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleTrie_Delete(t *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		word string
	}

	type tests []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}

	positiveTests := tests{
		{"inserted_word_is_not present_after_deletion", fields{root: New()}, args{word: "flawless"}, nil},
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SingleTrie{
				root: tt.fields.root,
			}

			last := tr.Insert(tt.args.word)

			if got := tr.Delete(tt.args.word); !reflect.DeepEqual(got, last) {
				t.Errorf("SingleTrie.Delete() = %v, want %v", got, last)
			}
		})
	}

	trieRoot := New()
	negativeTests := tests{
		{"missing_word_does not_fail_deletion", fields{root: trieRoot}, args{word: "excellent"}, trieRoot},
	}

	for _, tt := range negativeTests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SingleTrie{
				root: tt.fields.root,
			}

			if got := tr.Delete(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleTrie.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
