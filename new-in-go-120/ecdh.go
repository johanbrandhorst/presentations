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
	priv1, _ := curve.GenerateKey(rand.Reader)
	priv2, _ := curve.GenerateKey(rand.Reader)
	pub1, _ := curve.NewPublicKey(priv2.PublicKey().Bytes())
	secret, _ := priv1.ECDH(pub1)
	fmt.Println(hex.EncodeToString(secret))
}

// END OMIT
