package main

import (
	"math/rand"
	"time"
	"unicode"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomAdjective() string {
	return Adjectives[rand.Intn(len(Adjectives)-1)]
}

func randomNoun() string {
	return Nouns[rand.Intn(len(Nouns)-1)]
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func uppercaseFirstLetter(word string) string {
	a := []rune(word)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

func lowercaseFirstLetter(word string) string {
	a := []rune(word)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
func generateName() string {
	rand.Seed(time.Now().UnixNano())
	Noun1 := uppercaseFirstLetter(randomNoun())
	Noun2 := uppercaseFirstLetter(randomNoun())
	Adjective := lowercaseFirstLetter(randomAdjective())
	return Noun1 + Adjective + " " + Noun2
}

func main() {

}
