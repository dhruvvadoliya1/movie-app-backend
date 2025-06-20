package helpers

import (
	"context"
	"fmt"

	"github.com/dhruvvadoliya1/movie-app-backend/config"

	flipt "go.flipt.io/flipt-grpc"
	"google.golang.org/grpc"
)

var (
	fliptClient flipt.FliptClient
	initError   error
)

type BooleanFlagResponse struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled,omitempty"`
}

type Context struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type VarientFlagResponse struct {
	RequestContext Context `json:"request_context"`
	Match          bool    `json:"match,omitempty"`
	FlagKey        string  `json:"flag_key"`
	SegmentKey     string  `json:"segment_key,omitempty"`
	Value          string  `json:"value,omitempty"`
}

// InitFlizentClient make connection to flipt server and return flipt client if flipt functionality is enabled
func InitFliptClient() error {
	cfg := config.AllConfig.Flipt
	if !cfg.Enabled {
		initError = fmt.Errorf("flipt is not enabled")
		return initError
	}

	fliptServer := cfg.Host + ":" + cfg.Port
	conn, err := grpc.Dial(fliptServer, grpc.WithInsecure())
	if err != nil {
		initError = err
		return initError
	}

	fliptClient = flipt.NewFliptClient(conn)
	return nil
}

// GetBooleanFlag get boolean flag from flipt server by flag key
func GetBooleanFlag(flagKey string) (BooleanFlagResponse, error) {

	// check error while initializing flipt client
	if initError != nil {
		if config.AllConfig.Flipt.Enabled {
			return BooleanFlagResponse{}, initError
		}
		return BooleanFlagResponse{}, nil
	}

	var response BooleanFlagResponse
	flagResp, err := fliptClient.GetFlag(context.Background(), &flipt.GetFlagRequest{
		Key: flagKey,
	})
	if err != nil {
		return BooleanFlagResponse{}, fmt.Errorf("failed to get flag: %w", err)
	}

	if flagResp == nil {
		return BooleanFlagResponse{}, fmt.Errorf("flag is not found")
	}

	response.Name = flagResp.Name
	response.Key = flagResp.Key
	response.Description = flagResp.Description
	if flagResp.Enabled {
		response.Enabled = flagResp.Enabled
	}
	return response, nil
}

// GetVarientFlag get varient flag from flipt server by flagKey and constraint(contextMap)
func GetVarientFlag(flagKey string, entityId string, contextMap map[string]string) (VarientFlagResponse, error) {

	// check error while initializing flipt client
	if initError != nil {
		if config.AllConfig.Flipt.Enabled {
			return VarientFlagResponse{}, initError
		}
		return VarientFlagResponse{}, nil
	}

	var response VarientFlagResponse
	resp, err := fliptClient.Evaluate(context.Background(), &flipt.EvaluationRequest{
		FlagKey:  flagKey,
		EntityId: entityId,
		Context:  contextMap,
	})
	if err != nil {
		return VarientFlagResponse{}, fmt.Errorf("failed to evaluate flag: %w", err)
	}

	if resp == nil {
		return VarientFlagResponse{}, fmt.Errorf("flag is not found")
	}

	response.FlagKey = resp.FlagKey
	response.RequestContext = Context{Key: flagKey, Value: flagKey}

	if resp.Match {
		response.Match = resp.Match
		response.SegmentKey = resp.SegmentKey
		response.Value = resp.Value
	}

	return response, nil
}
