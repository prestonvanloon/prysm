package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	ptypes "github.com/gogo/protobuf/types"
	pbp2p "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	"google.golang.org/grpc"
)

var notImplementedErr = errors.New("not implemented")

func main() {
	fmt.Println("Starting gRPC server on port 4000")

	s := grpc.NewServer()
	srv := &beaconServer{}

	pb.RegisterBeaconServiceServer(s, srv)

	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type beaconServer struct {
}

func (s *beaconServer) CanonicalHead(_ context.Context, _ *ptypes.Empty) (*pbp2p.BeaconBlock, error) {
	return nil, notImplementedErr
}

func (s *beaconServer) CurrentAssignmentsAndGenesisTime(_ context.Context, req *pb.ValidatorAssignmentRequest) (*pb.CurrentAssignmentsResponse, error) {

	resp := &pb.CurrentAssignmentsResponse{}

	for i := uint64(0); i < 10000; i++ {
		resp.Assignments = append(resp.Assignments, &pb.Assignment{
			PublicKey:    req.PublicKeys[0],
			ShardId:      0,
			Role:         pb.ValidatorRole_PROPOSER,
			AssignedSlot: i,
		}, &pb.Assignment{
			PublicKey:    req.PublicKeys[0],
			ShardId:      0,
			Role:         pb.ValidatorRole_ATTESTER,
			AssignedSlot: i,
		})
	}

	resp.GenesisTimestamp = ptypes.TimestampNow()

	return resp, nil
}

func (s *beaconServer) LatestAttestation(_ *ptypes.Empty, _ pb.BeaconService_LatestAttestationServer) error {
	return notImplementedErr
}

func (s *beaconServer) ValidatorAssignments(_ *pb.ValidatorAssignmentRequest, _ pb.BeaconService_ValidatorAssignmentsServer) error {
	return notImplementedErr
}
