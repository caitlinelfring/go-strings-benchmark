package main

import (
	"fmt"
	"strings"
)

func main() {}

func Builder() string {
	s := &strings.Builder{}
	s.WriteString("The ")
	s.WriteString("quick ")
	s.WriteString("brown ")
	s.WriteString("fox ")
	s.WriteString("jumps ")
	s.WriteString("over ")
	s.WriteString("the ")
	s.WriteString("lazy ")
	s.WriteString("dog")
	return s.String()
}

func BuilderWithCap() string {
	s := &strings.Builder{}
	s.Grow(43)
	s.WriteString("The ")
	s.WriteString("quick ")
	s.WriteString("brown ")
	s.WriteString("fox ")
	s.WriteString("jumps ")
	s.WriteString("over ")
	s.WriteString("the ")
	s.WriteString("lazy ")
	s.WriteString("dog")
	return s.String()
}

func PlusEqual() string {
	s := ""
	s += "The "
	s += "quick "
	s += "brown "
	s += "fox "
	s += "jumps "
	s += "over "
	s += "the "
	s += "lazy "
	s += "dog"
	return s
}

func FmtSprint() string {
	return fmt.Sprintf("%s %s %s %s %s %s %s %s %s",
		"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog")
}

func Join() string {
	words := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	return strings.Join(words, " ")
}
