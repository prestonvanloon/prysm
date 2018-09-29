package main

import (
	"context"
	"fmt"

	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	libp2p "github.com/libp2p/go-libp2p"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listen, err := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/4001")
	if err != nil {
		panic(err)
	}
	opts := []libp2p.Option{
		libp2p.ListenAddrs(listen),
	}

	host, err := libp2p.New(ctx, opts...)
	if err != nil {
		panic(err)
	}

	// Construct a datastore (needed by the DHT). This is just a simple, in-memory thread-safe datastore.
	dstore := dsync.MutexWrap(ds.NewMapDatastore())

	// Make the DHT
	dht := kaddht.NewDHT(ctx, host, dstore)
	if err := dht.Bootstrap(ctx); err != nil {
		panic(err)
	}

	fmt.Printf("Bootstrap node running: /ip4/0.0.0.0/tcp/%v/p2p/%s\n", 4001, host.ID().Pretty())

	// TODO: Enable monitoring metrics
	// TODO: Log peer connections?

	// Blocking wait forever.
	select {}
}
