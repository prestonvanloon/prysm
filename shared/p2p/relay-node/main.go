package main

import (
	"context"
	"fmt"

	logger "github.com/ipfs/go-log"
	libp2p "github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	ma "github.com/multiformats/go-multiaddr"
)

func init() {
	logger.SetDebugLogging()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listen, err := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/4001")
	if err != nil {
		panic(err)
	}
	opts := []libp2p.Option{
		libp2p.ListenAddrs(listen),
		libp2p.EnableRelay(circuit.OptHop),
	}

	host, err := libp2p.New(ctx, opts...)
	if err != nil {
		panic(err)
	}

	fmt.Printf("/ip4/0.0.0.0/tcp/%v/p2p/%s\n", 4001, host.ID().Pretty())

	// Blocking wait forever.
	select {}
}
