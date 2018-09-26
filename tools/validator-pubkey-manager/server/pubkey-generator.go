package main

import "crypto/rand"

// RandomPubkey Generates a random pubkey
func RandomPubkey() []byte {
	pkey := make([]byte, 32)
	rand.Read(pkey)

	return pkey
}
