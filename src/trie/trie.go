package main

import (
	"fmt"
	"sort"
	"strings"
)

//Trie with slice of strings
type Trie struct {
	Values []string
}

//Constructor to initialize Tire
func Constructor() Trie {
	v := make([]string, 0)

	return Trie{Values: v}
}

//Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
	t.Values = append(t.Values, word)
}

//Search Returns if the word is in the trie
func (t *Trie) Search(word string) bool {
	sort.Strings(t.Values)
	i := sort.SearchStrings(t.Values, word)

	if i < len(t.Values) && t.Values[i] == word {
		return true
	}

	return false
}

//StartsWith Returns if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {

	for _, v := range t.Values {
		if strings.HasPrefix(v, prefix) {
			return true
		}
	}
	return false
}

func main() {
	obj := Constructor()

	obj.Insert("apple")
	obj.Insert("ball")
	obj.Insert("cat")
	fmt.Println(obj.Values)
	r := obj.Search("cat")
	fmt.Println(r)
	r = obj.Search("dog")
	fmt.Println(r)
	r = obj.StartsWith("ca")
	fmt.Println(r)
}
