package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	b := pb.NewBeaconServiceClient(conn)

	resp, err := b.CurrentAssignmentsAndGenesisTime(context.Background(), &pb.ValidatorAssignmentRequest{AllValidators: true})
	if err != nil {
		panic(err)
	}

	m := make(map[uint64][]*pb.Assignment)

	var maxSlot uint64
	for _, a := range resp.Assignments {
		m[a.AssignedSlot] = append(m[a.AssignedSlot], a)
		if a.AssignedSlot > maxSlot {
			maxSlot = a.AssignedSlot
		}
	}

	for i := uint64(0); i <= maxSlot; i++ {
		as := m[i]
		for _, a := range as {
			fmt.Printf("Slot=%d shard=%d role=%s pubkey=%#x\n", a.AssignedSlot, a.ShardId, a.Role, a.PublicKey.PublicKey)
		}
	}
}
