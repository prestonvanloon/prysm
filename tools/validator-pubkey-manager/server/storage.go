package main

type storage interface {
	GetConfigMap() (map[string][]byte, error)
	SetPubkey(pod string, pkey []byte) error
	RemovePod(pod string) error
}
