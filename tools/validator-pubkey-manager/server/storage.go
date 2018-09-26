package main

type storage interface {
	PubkeyMap() (map[string][]byte, error)
	SetPubkey(pod string, pkey []byte) error
	RemovePod(pod string) error
}
