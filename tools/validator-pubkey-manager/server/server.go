package main

import (
	"context"

	beaconpb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	pb "github.com/prysmaticlabs/prysm/proto/validator-pubkey-manager/v1"
)

type pubkeyManagerServer struct {
	storage storage
	pow     *powchainclient
}

func newServer(httpPath, address, privKey string) *pubkeyManagerServer {
	return &pubkeyManagerServer{
		storage: newKubernetesStorage(),
		pow:     newPowchainclient(httpPath, address, privKey),
	}
}

func (s *pubkeyManagerServer) GetPubkey(ctx context.Context, req *pb.GetPubkeyRequest) (*beaconpb.PublicKey, error) {
	// 1) Fetch the config map
	cm, err := s.storage.PubkeyMap()
	if err != nil {
		return nil, err
	}

	podName := req.PodName
	pkey, ok := cm[podName]

	// 2) If the podname is in the config map, return that pubkey
	if ok {
		return &beaconpb.PublicKey{
			PublicKey: pkey,
		}, nil
	}

	// 3) Otherwise, generate a new pubkey, update the map, and return the value.
	// Note: this could return a unallocated pubkey from the genesis set, in desired in the future.
	pkey = RandomPubkey()
	if err := s.pow.Deposit(ctx, pkey); err != nil {
		return nil, err
	}
	if err := s.storage.SetPubkey(podName, pkey); err != nil {
		return nil, err
	}
	return &beaconpb.PublicKey{PublicKey: pkey}, nil
}
