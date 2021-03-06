package service

import (
	"context"
	pb "github.com/MinterTeam/node-grpc-gateway/api_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Candidates returns list of candidates.
func (s *Service) Candidates(ctx context.Context, req *pb.CandidatesRequest) (*pb.CandidatesResponse, error) {
	cState, err := s.blockchain.GetStateForHeight(req.Height)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if req.Height != 0 {
		cState.Candidates().LoadCandidates()
	}

	candidates := cState.Candidates().GetCandidates()

	response := &pb.CandidatesResponse{}
	for _, candidate := range candidates {

		if timeoutStatus := s.checkTimeout(ctx); timeoutStatus != nil {
			return nil, timeoutStatus.Err()
		}

		isValidator := false
		if cState.Validators().GetByPublicKey(candidate.PubKey) != nil {
			isValidator = true
		}

		if req.Status != pb.CandidatesRequest_all {
			if req.Status == pb.CandidatesRequest_validator {
				if !isValidator {
					continue
				}
			} else if req.Status != pb.CandidatesRequest_CandidateStatus(candidate.Status) {
				continue
			}
		}

		if req.Height != 0 {
			cState.Candidates().LoadStakesOfCandidate(candidate.PubKey)
		}

		responseCandidate := makeResponseCandidate(cState, candidate, req.IncludeStakes, req.NotShowStakes)
		responseCandidate.Validator = isValidator

		response.Candidates = append(response.Candidates, responseCandidate)
	}

	return response, nil
}
