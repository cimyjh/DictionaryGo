package main

import (
	"fmt"

	"github.com/cimyjh/gojobs/DictionaryGo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Seccond")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
