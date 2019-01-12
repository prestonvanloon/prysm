package validator

import (
	"context"
	"fmt"

	pb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	p2pService "github.com/prysmaticlabs/prysm/shared/p2p"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type ValidatorService struct {
	ctx    context.Context
	cancel context.CancelFunc

	// Services
	p2p *p2pService.Server

	// Beacon chain gRPC
	rpcendpoint string
	conn        *grpc.ClientConn
	beacon      *pb.BeaconServiceClient
	proposer    *pb.ProposerServiceClient
	validator   *pb.ValidatorServiceClient
}

func NewValidatorService(
	ctx context.Context,
	beaconEndpoint string,
	p2p *p2pService.Server) *ValidatorService {

	ctx, cancel := context.WithCancel(ctx)

	return &ValidatorService{
		ctx:         ctx,
		cancel:      cancel,
		p2p:         p2p,
		rpcendpoint: beaconEndpoint,
	}
}

func (v *ValidatorService) Start() {
	var err error
	v.conn, err = grpc.DialContext(v.ctx, v.rpcendpoint, grpc.WithInsecure() /*TODO*/)
	if err != nil {
		// Non-blocking DialContext should only return an error if endpoint is
		// malformed or unacceptable.
		panic(err)
	}
}

func (v *ValidatorService) Stop() error {
	return nil
}

func (v *ValidatorService) Status() error {
	if v.p2p.Status() != nil {
		return v.p2p.Status()
	}

	if v.conn.GetState() != connectivity.Ready {
		return fmt.Errorf("grpc connection is %s", v.conn.GetState().String())
	}

	if v.ctx.Err() != nil {
		return v.ctx.Err()
	}

	return nil
}
