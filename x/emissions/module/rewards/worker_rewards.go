package rewards

import (
	alloraMath "github.com/allora-network/allora-chain/math"
	"github.com/allora-network/allora-chain/x/emissions/keeper"
	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TaskRewards struct {
	Address sdk.AccAddress
	Reward  alloraMath.Dec
}

var TASK_FORECAST = true
var TASK_INFERENCE = false

func GetInferenceTaskEntropy(
	ctx sdk.Context,
	k keeper.Keeper,
	topicId uint64,
	emaAlpha alloraMath.Dec,
	pRewardSpread alloraMath.Dec,
	betaEntropy alloraMath.Dec,
) (
	entropy alloraMath.Dec,
	modifiedRewardFractions []alloraMath.Dec,
	workers []sdk.AccAddress,
	err error,
) {
	return getInferenceOrForecastTaskEntropy(ctx, k, topicId, emaAlpha, pRewardSpread, betaEntropy, TASK_INFERENCE)
}

func GetForecastingTaskEntropy(
	ctx sdk.Context,
	k keeper.Keeper,
	topicId uint64,
	emaAlpha alloraMath.Dec,
	pRewardSpread alloraMath.Dec,
	betaEntropy alloraMath.Dec,
) (
	entropy alloraMath.Dec,
	modifiedRewardFractions []alloraMath.Dec,
	workers []sdk.AccAddress,
	err error,
) {
	return getInferenceOrForecastTaskEntropy(ctx, k, topicId, emaAlpha, pRewardSpread, betaEntropy, TASK_FORECAST)
}

func getInferenceOrForecastTaskEntropy(
	ctx sdk.Context,
	k keeper.Keeper,
	topicId uint64,
	emaAlpha alloraMath.Dec,
	pRewardSpread alloraMath.Dec,
	betaEntropy alloraMath.Dec,
	which bool,
) (
	entropy alloraMath.Dec,
	modifiedRewardFractions []alloraMath.Dec,
	workers []sdk.AccAddress,
	err error,
) {
	var scoresAtBlock types.Scores
	if which == TASK_INFERENCE {
		scoresAtBlock, err = k.GetWorkerInferenceScoresAtBlock(ctx, topicId, ctx.BlockHeight())
		if err != nil {
			return alloraMath.Dec{}, nil, nil, err
		}
	} else { // TASK_FORECAST
		scoresAtBlock, err = k.GetWorkerForecastScoresAtBlock(ctx, topicId, ctx.BlockHeight())
		if err != nil {
			return alloraMath.Dec{}, nil, nil, err
		}
	}
	numWorkers := len(scoresAtBlock.Scores)
	scores := make([]alloraMath.Dec, numWorkers)
	workers = make([]sdk.AccAddress, numWorkers)
	for i, scorePtr := range scoresAtBlock.Scores {
		scores[i] = scorePtr.Score
		addrStr := scorePtr.Address
		workerAddr, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return alloraMath.Dec{}, nil, nil, err
		}
		workers[i] = workerAddr
	}
	var previousRewardFraction alloraMath.Dec
	rewardFractions, err := GetScoreFractions(scores, pRewardSpread)
	if err != nil {
		return alloraMath.Dec{}, nil, nil, err
	}
	emaRewardFractions := make([]alloraMath.Dec, numWorkers)
	for i, fraction := range rewardFractions {
		noPriorScore := false
		if which == TASK_INFERENCE {
			previousRewardFraction, noPriorScore, err = k.GetPreviousInferenceRewardFraction(ctx, topicId, workers[i])
			if err != nil {
				return alloraMath.Dec{}, nil, nil, err
			}
		} else { // TASK_FORECAST
			previousRewardFraction, noPriorScore, err = k.GetPreviousForecastRewardFraction(ctx, topicId, workers[i])
			if err != nil {
				return alloraMath.Dec{}, nil, nil, err
			}
		}
		emaRewardFractions[i], err = alloraMath.CalcEma(
			emaAlpha,
			fraction,
			previousRewardFraction,
			noPriorScore,
		)
		if err != nil {
			return alloraMath.Dec{}, nil, nil, err
		}
	}
	numberRatio, err := NumberRatio(rewardFractions)
	if err != nil {
		return alloraMath.Dec{}, nil, nil, err
	}
	modifiedRewardFractions, err = ModifiedRewardFractions(emaRewardFractions)
	if err != nil {
		return alloraMath.Dec{}, nil, nil, err
	}
	entropy, err = Entropy(
		modifiedRewardFractions,
		numberRatio,
		alloraMath.NewDecFromInt64(int64(numWorkers)),
		betaEntropy,
	)
	if err != nil {
		return alloraMath.Dec{}, nil, nil, err
	}
	return entropy, modifiedRewardFractions, workers, nil
}

// The performance score of the entire forecasting task T_i
// is positive if the removal of the forecasting task would
// increase the network loss, and is negative if its removal
// would decrease the network loss
// We subtract the log-loss of the complete network inference
// (L_i) from that of the naive network (L_i^-), which is
// obtained by omitting all forecast-implied inferences
// T_i = log L_i^- - log L_i
// However we store the log based forms in the keeper
// so we do not need to take the logs again
func ForecastingPerformanceScore(
	naiveNetworkInferenceLoss,
	networkInferenceLoss alloraMath.Dec,
) (alloraMath.Dec, error) {
	return naiveNetworkInferenceLoss.Sub(networkInferenceLoss)
}

// we apply a utility function to the forecasting performance score
// to let the forecasting task utility range from the interval [0.1, 0.5]
// χ = 0.1 + 0.4σ(a*T_i − b)
// sigma is the sigmoid function
// a has fiduciary value of 8
// b has fiduciary value of 0.5
func ForecastingUtility(
	forecastingTaskUtilityScore,
	a,
	b alloraMath.Dec,
) (alloraMath.Dec, error) {
	aTimesForecastigPerformanceScore, err := a.Mul(forecastingTaskUtilityScore)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	aTimesForecastigPerformanceScoreMinusB, err := aTimesForecastigPerformanceScore.Sub(b)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	ret, err := Sigmoid(aTimesForecastigPerformanceScoreMinusB)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	zeroPointOne := alloraMath.MustNewDecFromString("0.1")
	zeroPointFour := alloraMath.MustNewDecFromString("0.4")
	ret, err = zeroPointFour.Mul(ret)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	ret, err = zeroPointOne.Add(ret)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	return ret, nil
}

// renormalize with a factor γ to ensure that the
// total reward allocated to workers (Ui + Vi)
// remains constant (otherwise, this would go at the expense of reputers)
// γ = (F_i + G_i) / ( (1 − χ)*F_i + χ*G_i)
func NormalizationFactor(
	entropyInference alloraMath.Dec,
	entropyForecasting alloraMath.Dec,
	forecastingUtility alloraMath.Dec,
) (alloraMath.Dec, error) {
	numerator, err := entropyInference.Add(entropyForecasting)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	oneMinusForecastingUtility, err := alloraMath.OneDec().Sub(forecastingUtility)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	oneMinusForecastingUtilityTimesEntropyInference, err := oneMinusForecastingUtility.Mul(entropyInference)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	forecastingUtilityTimesEntropyForecasting, err := forecastingUtility.Mul(entropyForecasting)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	denominator, err := oneMinusForecastingUtilityTimesEntropyInference.Add(forecastingUtilityTimesEntropyForecasting)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	ret, err := numerator.Quo(denominator)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	return ret, nil
}

// helper function to get chi and gamma
func getChiAndGamma(
	niaveNetworkInferenceLoss alloraMath.Dec,
	networkInferenceLoss alloraMath.Dec,
	entropyInference,
	entropyForecasting,
	a,
	b alloraMath.Dec,
) (chi alloraMath.Dec, gamma alloraMath.Dec, err error) {
	forecastingTaskUtilityScore, err := ForecastingPerformanceScore(
		niaveNetworkInferenceLoss,
		networkInferenceLoss,
	)
	if err != nil {
		return alloraMath.Dec{}, alloraMath.Dec{}, err
	}
	chi, err = ForecastingUtility(
		forecastingTaskUtilityScore,
		a,
		b,
	)
	if err != nil {
		return alloraMath.Dec{}, alloraMath.Dec{}, err
	}
	gamma, err = NormalizationFactor(
		entropyInference,
		entropyForecasting,
		chi,
	)
	if err != nil {
		return alloraMath.Dec{}, alloraMath.Dec{}, err
	}
	return chi, gamma, nil
}

// inference rewards calculation
// U_i = ((1 - χ) * γ * F_i * E_i ) / (F_i + G_i + H_i)
func GetRewardForInferenceTaskInTopic(
	niaveNetworkInferenceLoss alloraMath.Dec,
	networkInferenceLoss alloraMath.Dec,
	entropyInference alloraMath.Dec, // F_i
	entropyForecasting alloraMath.Dec, // G_i
	entropyReputer alloraMath.Dec, // H_i
	totalReward alloraMath.Dec, // E_i
	a alloraMath.Dec, // global param used for chi χ
	b alloraMath.Dec, // global param used for chi χ
) (alloraMath.Dec, error) {
	chi, gamma, err := getChiAndGamma(
		niaveNetworkInferenceLoss,
		networkInferenceLoss,
		entropyInference,
		entropyForecasting,
		a,
		b,
	)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	oneMinusChi, err := alloraMath.OneDec().Sub(chi)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	oneMinusChiGamma, err := oneMinusChi.Mul(gamma)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	oneMinusChiGammaEntropyInference, err := oneMinusChiGamma.Mul(entropyInference)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	numerator, err := oneMinusChiGammaEntropyInference.Mul(totalReward)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	entropyInferencePlusForecasting, err := entropyInference.Add(entropyForecasting)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	denominator, err := entropyInferencePlusForecasting.Add(entropyReputer)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	ret, err := numerator.Quo(denominator)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	return ret, nil
}

// forecaster rewards calculation
// V_i = (χ * γ * G_i * E_i) / (F_i + G_i + H_i)
func GetRewardForForecastingTaskInTopic(
	niaveNetworkInferenceLoss alloraMath.Dec,
	networkInferenceLoss alloraMath.Dec,
	entropyInference alloraMath.Dec, // F_i
	entropyForecasting alloraMath.Dec, // G_i
	entropyReputer alloraMath.Dec, // H_i
	totalReward alloraMath.Dec, // E_i
	sigmoidA alloraMath.Dec, // a used for sigmoid
	sigmoidB alloraMath.Dec, // b used for sigmoid
) (alloraMath.Dec, error) {
	chi, gamma, err := getChiAndGamma(
		niaveNetworkInferenceLoss,
		networkInferenceLoss,
		entropyInference,
		entropyForecasting,
		sigmoidA,
		sigmoidB,
	)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	chiGamma, err := chi.Mul(gamma)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	chiGammaEntropyForecasting, err := chiGamma.Mul(entropyForecasting)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	numerator, err := chiGammaEntropyForecasting.Mul(totalReward)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	entropyInferencePlusForecasting, err := entropyInference.Add(entropyForecasting)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	denominator, err := entropyInferencePlusForecasting.Add(entropyReputer)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	ret, err := numerator.Quo(denominator)
	if err != nil {
		return alloraMath.Dec{}, err
	}
	return ret, nil
}

func GetWorkersRewardsInferenceTask(
	ctx sdk.Context,
	keeper keeper.Keeper,
	topicId uint64,
	block int64,
	preward alloraMath.Dec,
	totalInferenceRewards alloraMath.Dec,
) ([]TaskRewards, error) {
	// Get network loss
	networkLosses, _, err := keeper.GetNetworkLossBundleAtOrBeforeBlock(ctx, topicId, block)
	if err != nil {
		return nil, err
	}

	// Get last score for each worker
	var scoresDec [][]alloraMath.Dec
	var workerAddresses []sdk.AccAddress
	for _, oneOutLoss := range networkLosses.OneOutInfererValues {
		workerAddr, err := sdk.AccAddressFromBech32(oneOutLoss.Worker)
		if err != nil {
			return nil, err
		}

		// Get worker last scores
		workerLastScores, err := keeper.GetWorkerInferenceScoresUntilBlock(ctx, topicId, block, workerAddr)
		if err != nil {
			return nil, err
		}

		// Add worker address in the worker addresses array
		workerAddresses = append(workerAddresses, workerAddr)

		var workerLastScoresDec []alloraMath.Dec
		for _, score := range workerLastScores {
			workerLastScoresDec = append(workerLastScoresDec, score.Score)
		}
		scoresDec = append(scoresDec, workerLastScoresDec)
	}

	// Get worker portion of rewards
	rewards, err := GetWorkerPortionOfRewards(scoresDec, preward, totalInferenceRewards, workerAddresses)
	if err != nil {
		return nil, err
	}

	return GetRewardsWithOutTax(ctx, keeper, rewards, topicId)
}

func GetWorkersRewardsForecastTask(
	ctx sdk.Context,
	keeper keeper.Keeper,
	topicId uint64,
	block int64,
	preward alloraMath.Dec,
	totalForecastRewards alloraMath.Dec,
) ([]TaskRewards, error) {
	// Get network loss
	networkLosses, _, err := keeper.GetNetworkLossBundleAtOrBeforeBlock(ctx, topicId, block)
	if err != nil {
		return nil, err
	}

	// Get new score for each worker
	var scoresDec [][]alloraMath.Dec
	var workerAddresses []sdk.AccAddress
	for _, oneOutLoss := range networkLosses.OneOutForecasterValues {
		workerAddr, err := sdk.AccAddressFromBech32(oneOutLoss.Worker)
		if err != nil {
			return nil, err
		}

		// Get worker last scores
		workerLastScores, err := keeper.GetWorkerForecastScoresUntilBlock(ctx, topicId, block, workerAddr)
		if err != nil {
			return nil, err
		}

		// Add worker address in the worker addresses array
		workerAddresses = append(workerAddresses, workerAddr)

		// Convert scores to alloraMath.Dec
		var workerLastScoresDec []alloraMath.Dec
		for _, score := range workerLastScores {
			workerLastScoresDec = append(workerLastScoresDec, score.Score)
		}
		scoresDec = append(scoresDec, workerLastScoresDec)
	}

	// Get worker portion of rewards
	rewards, err := GetWorkerPortionOfRewards(scoresDec, preward, totalForecastRewards, workerAddresses)

	if err != nil {
		return nil, err
	}

	return GetRewardsWithOutTax(ctx, keeper, rewards, topicId)
}

func GetRewardsWithOutTax(
	ctx sdk.Context,
	keeper keeper.Keeper,
	rewards []TaskRewards,
	topicId uint64,
) ([]TaskRewards, error) {
	params, err := keeper.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	var result []TaskRewards
	// Get average reward for this worker
	for _, reward := range rewards {
		avg, err := keeper.GetAverageWorkerReward(ctx, topicId, reward.Address)
		if err != nil {
			continue
		}
		avgValueTimesCount, err := avg.Value.Mul(alloraMath.NewDecFromInt64(int64(avg.Count)))
		if err != nil {
			continue
		}
		totalRewards, err := avgValueTimesCount.Add(reward.Reward)
		if err != nil {
			continue
		}
		avg.Count += 1
		avg.Value, err = totalRewards.Quo(alloraMath.NewDecFromInt64(int64(avg.Count)))
		if err != nil {
			continue
		}
		_ = keeper.SetAverageWorkerReward(ctx, topicId, reward.Address, avg)
		fee, err := CalculateWorkerTax(avg.Value, params.SybilTaxExponent, params.NumberExpectedInferenceSybils)
		if err != nil {
			continue
		}
		reward.Reward, err = reward.Reward.Sub(fee)
		if err != nil {
			continue
		}
		if reward.Reward.Lt(alloraMath.ZeroDec()) {
			reward.Reward = alloraMath.ZeroDec()
		}
		result = append(result, TaskRewards{
			Address: reward.Address,
			Reward:  reward.Reward,
		})
	}

	return result, nil
}
