package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"unique"
)

// wordGen returns a generator of random words of length wordLen
// from a set of nDistinct unique words.
func wordGen(nDistinct, wordLen int) func() string {
	vocab := make([]string, nDistinct)
	for i := range nDistinct {
		word := randomString(wordLen)
		vocab[i] = word
	}
	return func() string {
		word := vocab[rand.Intn(nDistinct)]
		return strings.Clone(word)
	}
}

// randomString returns a random string of length n.
func randomString(n int) string {
	b := make([]byte, n)
	_, _ = cryptorand.Read(b)
	return string(b)
}

const nDistinct = 100
const wordLen = 40

// START OMIT
const nWords = 10000

var words []string
var internedWords []unique.Handle[string]

func generateWords(nWords int, genFn func() string) []string {
	words = make([]string, nWords)
	for i := range nWords {
		words[i] = genFn()
	}
	return words
}

func generateWordsInterned(nWords int, genFn func() string) []string {
	internedWords = make([]unique.Handle[string], nWords)
	for i := range nWords {
		internedWords[i] = unique.Make(genFn())
	}
	return words
}

// END OMIT

func main() {
	generate := wordGen(nDistinct, wordLen)
	memBeforeWithout := &runtime.MemStats{}
	runtime.ReadMemStats(memBeforeWithout)

	// store words without interning
	generateWords(nWords, generate)
	memAfterWithout := &runtime.MemStats{}
	runtime.ReadMemStats(memAfterWithout)
	memUsedWithout := memAfterWithout.Alloc - memBeforeWithout.Alloc
	fmt.Printf("Memory used when NOT interning: %d KB\n", memUsedWithout/1024)

	words = nil
	runtime.GC()

	memBeforeWith := &runtime.MemStats{}
	runtime.ReadMemStats(memBeforeWith)
	// store words with interning
	generateWordsInterned(nWords, generate)

	memAfterWith := &runtime.MemStats{}
	runtime.ReadMemStats(memAfterWith)
	memUsedWith := memAfterWith.Alloc - memBeforeWith.Alloc
	fmt.Printf("Memory used when interning: %d KB\n", memUsedWith/1024)
}
