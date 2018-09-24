package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/prysmaticlabs/prysm/proto/validator-pubkey-manager/v1"
)

var (
	server  = flag.String("server", "", "The server address in the format of host:port")
	podname = flag.String("pod-name", "", "The name of the pod")
	out     = flag.String("out", "", "The file path to write the pubkey")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server '%s' %v", *server, err)
	}

	client := pb.NewPubkeyManagerClient(conn)

	log.Printf("Getting pubkey for pod %s", *podname)

	req := &pb.GetPubkeyRequest{PodName: *podname}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := client.GetPubkey(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get pubkey %v", err)
	}

	log.Printf("Received resp %v", res)

	if *out != "" {
		f, err := os.Create(*out)
		if err != nil {
			log.Fatalf("Failed to open file for write %v", err)
		}
		defer f.Close()
		_, err = f.Write(res.PublicKey)
		if err != nil {
			log.Fatalf("Failed to write public key %v", err)
		}
	}
}
