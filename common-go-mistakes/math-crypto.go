package main

import (
	"crypto/ed25519"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// 😰
	pub, priv, err := ed25519.GenerateKey(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = pub
	_ = priv
}
