package queryserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	synth "github.com/allora-network/allora-chain/x/emissions/keeper/inference_synthesis"
	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetWorkerLatestInferenceByTopicId handles the query for the latest inference by a specific worker for a given topic.
func (qs queryServer) GetWorkerLatestInferenceByTopicId(ctx context.Context, req *types.QueryWorkerLatestInferenceRequest) (*types.QueryWorkerLatestInferenceResponse, error) {
	if err := qs.k.ValidateStringIsBech32(req.WorkerAddress); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	topicExists, err := qs.k.TopicExists(ctx, req.TopicId)
	if !topicExists {
		return nil, status.Errorf(codes.NotFound, "topic %v not found", req.TopicId)
	} else if err != nil {
		return nil, err
	}

	inference, err := qs.k.GetWorkerLatestInferenceByTopicId(ctx, req.TopicId, req.WorkerAddress)
	if err != nil {
		return nil, err
	}

	return &types.QueryWorkerLatestInferenceResponse{LatestInference: &inference}, nil
}

func (qs queryServer) GetInferencesAtBlock(ctx context.Context, req *types.QueryInferencesAtBlockRequest) (*types.QueryInferencesAtBlockResponse, error) {
	topicExists, err := qs.k.TopicExists(ctx, req.TopicId)
	if !topicExists {
		return nil, status.Errorf(codes.NotFound, "topic %v not found", req.TopicId)
	} else if err != nil {
		return nil, err
	}

	inferences, err := qs.k.GetInferencesAtBlock(ctx, req.TopicId, req.BlockHeight)
	if err != nil {
		return nil, err
	}

	return &types.QueryInferencesAtBlockResponse{Inferences: inferences}, nil
}

// Return full set of inferences in I_i from the chain
func (qs queryServer) GetNetworkInferencesAtBlock(
	ctx context.Context,
	req *types.QueryNetworkInferencesAtBlockRequest,
) (*types.QueryNetworkInferencesAtBlockResponse, error) {
	topic, err := qs.k.GetTopic(ctx, req.TopicId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "topic %v not found", req.TopicId)
	}
	if topic.EpochLastEnded == 0 {
		return nil, status.Errorf(codes.NotFound, "network inference not available for topic %v", req.TopicId)
	}

	networkInferences, _, _, _, err := synth.GetNetworkInferencesAtBlock(
		sdk.UnwrapSDKContext(ctx),
		qs.k,
		req.TopicId,
		req.BlockHeightLastInference,
		req.BlockHeightLastReward,
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryNetworkInferencesAtBlockResponse{NetworkInferences: networkInferences}, nil
}

// Return full set of inferences in I_i from the chain, as well as weights and forecast implied inferences
func (qs queryServer) GetLatestNetworkInference(
	ctx context.Context,
	req *types.QueryLatestNetworkInferencesRequest,
) (
	*types.QueryLatestNetworkInferencesResponse,
	error,
) {
	topicExists, err := qs.k.TopicExists(ctx, req.TopicId)
	if !topicExists {
		return nil, status.Errorf(codes.NotFound, "topic %v not found", req.TopicId)
	} else if err != nil {
		return nil, err
	}

	networkInferences, forecastImpliedInferenceByWorker, infererWeights, forecasterWeights, inferenceBlockHeight, lossBlockHeight, err := synth.GetLatestNetworkInference(
		sdk.UnwrapSDKContext(ctx),
		qs.k,
		req.TopicId,
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryLatestNetworkInferencesResponse{
		NetworkInferences:         networkInferences,
		InfererWeights:            synth.ConvertWeightsToArrays(infererWeights),
		ForecasterWeights:         synth.ConvertWeightsToArrays(forecasterWeights),
		ForecastImpliedInferences: synth.ConvertForecastImpliedInferencesToArrays(forecastImpliedInferenceByWorker),
		InferenceBlockHeight:      inferenceBlockHeight,
		LossBlockHeight:           lossBlockHeight,
	}, nil
}

func (qs queryServer) GetLatestAvailableNetworkInference(
	ctx context.Context,
	req *types.QueryLatestNetworkInferencesRequest,
) (
	*types.QueryLatestNetworkInferencesResponse,
	error,
) {

	// Find the latest inference block height
	_, inferenceBlockHeight, err := qs.k.GetLatestTopicInferences(ctx, req.TopicId)
	if err != nil {
		return nil, err
	}

	topic, err := qs.k.GetTopic(ctx, req.TopicId)
	if err != nil {
		return nil, err
	}

	params, err := qs.k.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	// calculate the previous loss block height
	previousLossBlockHeight := inferenceBlockHeight - topic.EpochLength

	for i := 0; i < int(params.DefaultPageLimit); i++ {
		if previousLossBlockHeight < 0 {
			return nil, status.Errorf(codes.NotFound, "losses not yet available for topic %v", req.TopicId)
		}

		// check for inferences
		inferences, err := qs.k.GetInferencesAtBlock(ctx, req.TopicId, inferenceBlockHeight)
		if err == nil && inferences != nil && len(inferences.Inferences) > 0 {
			// check for losses. If they exist, then break so we can query
			_, err = qs.k.GetNetworkLossBundleAtBlock(ctx, req.TopicId, previousLossBlockHeight)
			if err == nil {
				break
			}
		}

		inferenceBlockHeight -= topic.EpochLength
		previousLossBlockHeight -= topic.EpochLength
	}

	networkInferences, forecastImpliedInferenceByWorker, infererWeights, forecasterWeights, err :=
		synth.GetNetworkInferencesAtBlock(
			sdk.UnwrapSDKContext(ctx),
			qs.k,
			req.TopicId,
			inferenceBlockHeight,
			previousLossBlockHeight,
		)

	if err != nil {
		return nil, err
	}

	return &types.QueryLatestNetworkInferencesResponse{
		NetworkInferences:         networkInferences,
		InfererWeights:            synth.ConvertWeightsToArrays(infererWeights),
		ForecasterWeights:         synth.ConvertWeightsToArrays(forecasterWeights),
		ForecastImpliedInferences: synth.ConvertForecastImpliedInferencesToArrays(forecastImpliedInferenceByWorker),
		InferenceBlockHeight:      inferenceBlockHeight,
		LossBlockHeight:           previousLossBlockHeight,
	}, nil
}
