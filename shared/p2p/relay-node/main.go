package main

import (
	"context"
	"log"

	logger "github.com/ipfs/go-log"
	insecure "github.com/libp2p/go-conn-security/insecure"
	libp2p "github.com/libp2p/go-libp2p"
	relay "github.com/libp2p/go-libp2p-circuit"
	tptu "github.com/libp2p/go-libp2p-transport-upgrader"
	ma "github.com/multiformats/go-multiaddr"
	mplex "github.com/whyrusleeping/go-smux-multiplex"
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
	}

	host, err := libp2p.New(ctx, opts...)
	if err != nil {
		panic(err)
	}
	// wat is this ???
	upgrader := &tptu.Upgrader{
		Secure: insecure.New("whatever"),
		Muxer:  new(mplex.Transport),
	}

	ropts := []relay.RelayOpt{}

	node, err := relay.NewRelay(ctx, host, upgrader, ropts...)

	if err != nil {
		panic(err)
	}

	_ = node

	log.Println("Listening on port 4000, i think")

	// Blocking wait forever.
	select {}
}
