package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// START OMIT
func main() {
	curve := ecdh.P256()
	priv, _ := curve.GenerateKey(rand.Reader)
	pub, _ := curve.NewPublicKey([]byte("other-parties-public-key"))
	secret, _ := priv.ECDH(pub)
	fmt.Println(hex.EncodeToString(secret))
}

// END OMIT
