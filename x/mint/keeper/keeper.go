package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/allora-network/allora-chain/app/params"
	"github.com/allora-network/allora-chain/x/mint/types"

	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeService     storetypes.KVStoreService
	accountKeeper    types.AccountKeeper
	stakingKeeper    types.StakingKeeper
	bankKeeper       types.BankKeeper
	emissionsKeeper  types.EmissionsKeeper
	feeCollectorName string

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	Schema                                    collections.Schema
	Params                                    collections.Item[types.Params]
	PreviousRewardEmissionsPerUnitStakedToken collections.Item[math.Int]
	EcosystemTokensMinted                     collections.Item[math.Int]
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	sk types.StakingKeeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	ek types.EmissionsKeeper,
	feeCollectorName string,
	authority string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("the x/%s module account has not been set", types.ModuleName))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:              cdc,
		storeService:     storeService,
		stakingKeeper:    sk,
		accountKeeper:    ak,
		bankKeeper:       bk,
		emissionsKeeper:  ek,
		feeCollectorName: feeCollectorName,
		authority:        authority,
		Params:           collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		PreviousRewardEmissionsPerUnitStakedToken: collections.NewItem(sb, types.PreviousRewardsKey, "previousrewardsemissionsperunitstakedtoken", sdk.IntValue),
		EcosystemTokensMinted:                     collections.NewItem(sb, types.EcosystemTokensMintedKey, "ecosystemtokensminted", sdk.IntValue),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
	return k
}

// GetAuthority returns the x/mint module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", "x/"+types.ModuleName)
}

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx context.Context) (math.Int, error) {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx context.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) PayEmissionsFromEcosystemAccount(ctx context.Context, rewards sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.EcosystemModuleName, emissionstypes.AlloraRewardsAccountName, rewards)
}

// GetSupply implements an alias call to the underlying supply keeper's
// GetSupply to be used in BeginBlocker.
func (k Keeper) GetSupply(ctx context.Context) sdk.Coin {
	return k.bankKeeper.GetSupply(ctx, params.BaseCoinUnit)
}

func (k Keeper) GetEcosystemAddress() sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(types.EcosystemModuleName)
}

func (k Keeper) GetEcosystemBalance(ctx context.Context, mintDenom string) (math.Int, error) {
	ecosystemAddr := k.GetEcosystemAddress()
	return k.bankKeeper.GetBalance(ctx, ecosystemAddr, mintDenom).Amount, nil
}
