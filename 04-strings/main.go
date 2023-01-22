package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	word := "A word"
	replacer := strings.NewReplacer("A", "Another")
	finallyWord := replacer.Replace(word)
	pl(finallyWord)
	pl("Length", len(finallyWord))
	pl("Contains Another", strings.Contains(finallyWord, "Another"))
	pl("o index", strings.Index(finallyWord, "o"))
	pl("Split", strings.Split(finallyWord, ""))
	pl("Lower", strings.ToLower(finallyWord))
	pl("StartsWith ", strings.HasPrefix(finallyWord, "A"))
}
