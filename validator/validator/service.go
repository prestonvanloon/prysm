package validator

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	ptypes "github.com/gogo/protobuf/types"
	p2ppb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	p2pService "github.com/prysmaticlabs/prysm/shared/p2p"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/slotticker"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var log = logrus.WithField("prefix", "validator")

type ValidatorService struct {
	ctx    context.Context
	cancel context.CancelFunc

	pubKey           *pb.PublicKey
	assignmentLock   sync.RWMutex
	genesisTime      time.Time
	epochAssignments []*pb.Assignment

	// Services
	p2p *p2pService.Server

	// Beacon chain gRPC
	rpcendpoint string
	conn        *grpc.ClientConn
	beacon      pb.BeaconServiceClient
	//proposer    *pb.ProposerServiceClient
	validator pb.ValidatorServiceClient
}

func NewValidatorService(
	ctx context.Context,
	beaconEndpoint string,
	p2p *p2pService.Server,
	pubKey []byte) *ValidatorService {

	ctx, cancel := context.WithCancel(ctx)

	return &ValidatorService{
		ctx:         ctx,
		cancel:      cancel,
		p2p:         p2p,
		rpcendpoint: beaconEndpoint,
		pubKey:      &pb.PublicKey{PublicKey: pubKey},
	}
}

func (v *ValidatorService) Start() {
	log.WithField(
		"pubkey",
		fmt.Sprintf("%#x", v.pubKey.PublicKey),
	).Info("Starting service")

	v.p2p.RegisterTopic(p2ppb.Topic_BEACON_BLOCK_ANNOUNCE.String(), &p2ppb.BeaconBlock{})

	var err error
	v.conn, err = grpc.DialContext(v.ctx, v.rpcendpoint, grpc.WithInsecure() /*TODO*/)
	if err != nil {
		// Non-blocking DialContext should only return an error if endpoint is
		// malformed or unacceptable.
		panic(err)
	}

	v.beacon = pb.NewBeaconServiceClient(v.conn)
	v.validator = pb.NewValidatorServiceClient(v.conn)

	go v.MainRoutine()
}

func (v *ValidatorService) Stop() error {
	v.cancel()
	return nil
}

func (v *ValidatorService) Status() error {
	if v.conn.GetState() != connectivity.Ready {
		return fmt.Errorf("grpc connection is %s", v.conn.GetState().String())
	}
	if v.ctx.Err() != nil {
		return v.ctx.Err()
	}
	if v.p2p.Status() != nil {
		return fmt.Errorf("unhealthy dependency on p2p: %v", v.p2p.Status())
	}

	return nil
}

func (v *ValidatorService) MainRoutine() {
	v.initialize()
	go v.monitorAssignmentUpdates()

	// monitor slot updates
	v.assignmentLock.RLock()
	slotTicker := slotticker.GetSlotTicker(v.genesisTime, params.BeaconConfig().SlotDuration)
	v.assignmentLock.RUnlock()

	for {
		select {
		case slot := <-slotTicker.C():
			v.assignmentLock.RLock()
			defer v.assignmentLock.RUnlock()

			// If assignment is thing, the do the thing.
			log.WithField("slot", slot).Info("Processing slot")

			// TODO: Improve this to be O(1) lookup
			for _, a := range v.epochAssignments {
				if a.AssignedSlot == slot {
					log.WithFields(logrus.Fields{
						"role":  a.Role,
						"shard": a.ShardId,
					}).Info("Performing role")
				}

				switch a.Role {
				case pb.ValidatorRole_PROPOSER:
					// Propose a new block
					//
					blk := &p2ppb.BeaconBlock{
						Slot: slot,

						ParentRootHash32:   []byte{}, // TODO
						StateRootHash32:    []byte{},
						RandaoRevealHash32: []byte{},
						DepositRootHash32:  []byte{},
						Body: &p2ppb.BeaconBlockBody{
							Attestations:      []*p2ppb.Attestation{},
							ProposerSlashings: []*p2ppb.ProposerSlashing{},
							CasperSlashings:   []*p2ppb.CasperSlashing{},
							Deposits:          []*p2ppb.Deposit{},
							Exits:             []*p2ppb.Exit{},
						},
					}

					blk.Signature = [][]byte{}

					v.p2p.Broadcast(blk)
					break
				case pb.ValidatorRole_ATTESTER:
					// do attestor stuff
					break
				case pb.ValidatorRole_UNKNOWN:
				default:
					log.WithField("role", a.Role).Warn("Unknown role")
					break
				}
			}

		case <-v.ctx.Done():
			slotTicker.Done()

		default:
		}
	}
}

func (v *ValidatorService) initialize() {
	v.assignmentLock.Lock()

	resp, err := v.beacon.CurrentAssignmentsAndGenesisTime(
		v.ctx,
		&pb.ValidatorAssignmentRequest{
			//AllValidators: true,
			PublicKeys: []*pb.PublicKey{
				v.pubKey,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	v.epochAssignments = resp.Assignments
	gt, err := ptypes.TimestampFromProto(resp.GenesisTimestamp)
	if err != nil {
		panic(err)
	}
	v.genesisTime = gt
	v.assignmentLock.Unlock()

	log.WithFields(logrus.Fields{
		"assignments": resp.Assignments,
		"genesis":     gt,
	}).Info("Received current assignments")

}

func (v *ValidatorService) monitorAssignmentUpdates() {
	// Start streaming assignment updates
	stream, err := v.beacon.ValidatorAssignments(v.ctx, &pb.ValidatorAssignmentRequest{})
	if err != nil {
		panic(err) // TODO
	}

	for {
		assignment, err := stream.Recv()
		if err == io.EOF {
			// Server terminated stream
			// TODO: Print warning or do something about this.
			break
		}

		if err != nil {
			fmt.Printf("Received error from assignment stream: %v", err)
			break
		}

		log.WithField("assignment", assignment).Info("Received assignment")
	}
}
