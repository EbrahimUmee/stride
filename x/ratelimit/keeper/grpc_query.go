package keeper

import (
	"context"

	"github.com/Stride-Labs/stride/v3/x/ratelimit/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Paths(c context.Context, req *types.QueryPathsRequest) (*types.QueryPathsResponse, error) {
	// TODO:
	return &types.QueryPathsResponse{}, nil
}

func (k Keeper) Path(c context.Context, req *types.QueryPathRequest) (*types.QueryPathResponse, error) {
	// TODO:
	return &types.QueryPathResponse{}, nil
}
func (k Keeper) RateLimits(c context.Context, req *types.QueryRateLimitsRequest) (*types.QueryRateLimitsResponse, error) {
	// TODO:
	return &types.QueryRateLimitsResponse{}, nil
}

func (k Keeper) RateLimit(c context.Context, req *types.QueryRateLimitRequest) (*types.QueryRateLimitResponse, error) {
	// TODO:
	return &types.QueryRateLimitResponse{}, nil
}