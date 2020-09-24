package service

import (
	"context"
	"encoding/hex"
	"github.com/MinterTeam/minter-go-node/core/types"
	pb "github.com/MinterTeam/node-grpc-gateway/api_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// WaitList returns the list of address stakes in waitlist.
func (s *Service) WaitList(ctx context.Context, req *pb.WaitListRequest) (*pb.WaitListResponse, error) {
	if !strings.HasPrefix(strings.Title(req.Address), "Mx") {
		return new(pb.WaitListResponse), status.Error(codes.InvalidArgument, "invalid address")
	}

	decodeString, err := hex.DecodeString(req.Address[2:])
	if err != nil {
		return new(pb.WaitListResponse), status.Error(codes.InvalidArgument, "invalid address")
	}

	address := types.BytesToAddress(decodeString)

	if !strings.HasPrefix(req.PublicKey, "Mp") {
		return new(pb.WaitListResponse), status.Error(codes.InvalidArgument, "public key don't has prefix 'Mp'")
	}

	publickKey := types.HexToPubkey(req.PublicKey)

	cState, err := s.blockchain.GetStateForHeight(req.Height)
	if err != nil {
		return new(pb.WaitListResponse), status.Error(codes.NotFound, err.Error())
	}

	cState.RLock()
	defer cState.RUnlock()

	if timeoutStatus := s.checkTimeout(ctx); timeoutStatus != nil {
		return new(pb.WaitListResponse), timeoutStatus.Err()
	}

	response := new(pb.WaitListResponse)
	items := cState.WaitList().GetByAddressAndPubKey(address, publickKey)
	response.List = make([]*pb.WaitListResponse_Wait, 0, len(items))
	for _, item := range items {
		if timeoutStatus := s.checkTimeout(ctx); timeoutStatus != nil {
			return new(pb.WaitListResponse), timeoutStatus.Err()
		}

		response.List = append(response.List, &pb.WaitListResponse_Wait{
			Coin: &pb.Coin{
				Id:     uint64(item.Coin),
				Symbol: cState.Coins().GetCoin(item.Coin).CSymbol.String(),
			},
			Value: item.Value.String(),
		})
	}

	return response, nil
}
