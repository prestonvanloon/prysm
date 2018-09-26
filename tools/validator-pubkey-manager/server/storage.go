package main

import "context"

type storage interface {
	PubkeyMap(ctx context.Context) (map[string][]byte, error)
	SetPubkey(ctx context.Context, pod string, pkey []byte) error
	RemovePod(ctx context.Context, pod string) error
}
