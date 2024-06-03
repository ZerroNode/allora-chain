package inference_synthesis

import (
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	cosmosMath "cosmossdk.io/math"
	"github.com/allora-network/allora-chain/x/emissions/keeper"
	emissions "github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Calculates all network inferences in the set I_i given historical state (e.g. regrets)
// and data from workers (e.g. inferences, forecast-implied inferences)
// as of a specified block height
func GetNetworkInferencesAtBlock(
	ctx sdk.Context,
	k keeper.Keeper,
	topicId TopicId,
	inferencesNonce BlockHeight,
	previousLossNonce BlockHeight,
) (*emissions.ValueBundle, error) {
	networkInferences := &emissions.ValueBundle{
		TopicId:          topicId,
		InfererValues:    make([]*emissions.WorkerAttributedValue, 0),
		ForecasterValues: make([]*emissions.WorkerAttributedValue, 0),
	}

	inferences, err := k.GetInferencesAtBlock(ctx, topicId, inferencesNonce)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "no inferences found for topic %v at block %v", topicId, inferencesNonce)
	}
	// Add inferences in the bundle -> this bundle will be used as a fallback in case of error
	for _, infererence := range inferences.Inferences {
		networkInferences.InfererValues = append(networkInferences.InfererValues, &emissions.WorkerAttributedValue{
			Worker: infererence.Inferer,
			Value:  infererence.Value,
		})
	}

	forecasts, err := k.GetForecastsAtBlock(ctx, topicId, inferencesNonce)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			forecasts = &emissions.Forecasts{
				Forecasts: make([]*emissions.Forecast, 0),
			}
		} else {
			return nil, err
		}
	}

	if len(inferences.Inferences) > 1 {
		moduleParams, err := k.GetParams(ctx)
		if err != nil {
			return nil, err
		}

		reputerReportedLosses, err := k.GetReputerLossBundlesAtBlock(ctx, topicId, previousLossNonce)
		if err != nil {
			ctx.Logger().Warn(fmt.Sprintf("Error getting reputer losses: %s", err.Error()))
			return networkInferences, nil
		}

		// Map list of stakesOnTopic to map of stakesByReputer
		stakesByReputer := make(map[string]cosmosMath.Int)
		for _, bundle := range reputerReportedLosses.ReputerValueBundles {
			stakeAmount, err := k.GetStakeOnReputerInTopic(ctx, topicId, bundle.ValueBundle.Reputer)
			if err != nil {
				ctx.Logger().Warn(fmt.Sprintf("Error getting stake on reputer: %s", err.Error()))
				return networkInferences, nil
			}
			stakesByReputer[bundle.ValueBundle.Reputer] = stakeAmount
		}

		networkCombinedLoss, err := CalcCombinedNetworkLoss(
			stakesByReputer,
			reputerReportedLosses,
			moduleParams.Epsilon,
		)
		if err != nil {
			ctx.Logger().Warn(fmt.Sprintf("Error calculating network combined loss: %s", err.Error()))
			return networkInferences, nil
		}
		topic, err := k.GetTopic(ctx, topicId)
		if err != nil {
			ctx.Logger().Warn(fmt.Sprintf("Error getting topic: %s", err.Error()))
			return networkInferences, nil
		}
		networkInferenceBuilder, err := NewNetworkInferenceBuilderFromSynthRequest(
			SynthRequest{
				Ctx:                 ctx,
				K:                   k,
				TopicId:             topicId,
				Inferences:          inferences,
				Forecasts:           forecasts,
				NetworkCombinedLoss: networkCombinedLoss,
				Epsilon:             moduleParams.Epsilon,
				PNorm:               topic.PNorm,
				CNorm:               moduleParams.CNorm,
			},
		)
		if err != nil {
			ctx.Logger().Warn(fmt.Sprintf("Error constructing network inferences builder topic: %s", err.Error()))
			return nil, err
		}
		networkInferences = networkInferenceBuilder.CalcAndSetNetworkInferences().Build()
	} else {
		// If there is only one valid inference, then the network inference is the same as the single inference
		// For the forecasts to be meaningful, there should be at least 2 inferences
		singleInference := inferences.Inferences[0]

		networkInferences = &emissions.ValueBundle{
			TopicId:       topicId,
			CombinedValue: singleInference.Value,
			InfererValues: []*emissions.WorkerAttributedValue{
				{
					Worker: singleInference.Inferer,
					Value:  singleInference.Value,
				},
			},
			ForecasterValues:       []*emissions.WorkerAttributedValue{},
			NaiveValue:             singleInference.Value,
			OneOutInfererValues:    []*emissions.WithheldWorkerAttributedValue{},
			OneOutForecasterValues: []*emissions.WithheldWorkerAttributedValue{},
			OneInForecasterValues:  []*emissions.WorkerAttributedValue{},
		}
	}

	return networkInferences, nil
}
