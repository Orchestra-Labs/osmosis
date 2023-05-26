package model

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	"github.com/osmosis-labs/osmosis/v15/x/cosmwasmpool/cosmwasm/msg"
	"github.com/osmosis-labs/osmosis/v15/x/cosmwasmpool/types"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v15/x/poolmanager/types"

	cosmwasmutils "github.com/osmosis-labs/osmosis/osmoutils/cosmwasm"
)

type Pool struct {
	CosmWasmPool
	WasmKeeper types.WasmKeeper
}

var (
	_ poolmanagertypes.PoolI  = &Pool{}
	_ types.CosmWasmExtension = &Pool{}
)

// NewCosmWasmPool creates a new CosmWasm pool with the specified parameters.
func NewCosmWasmPool(poolId uint64, codeId uint64, instantiateMsg []byte) Pool {
	pool := Pool{
		CosmWasmPool: CosmWasmPool{
			PoolAddress:     poolmanagertypes.NewPoolAddress(poolId).String(),
			ContractAddress: "", // N.B. This is to be set in InitializePool()
			PoolId:          poolId,
			CodeId:          codeId,
			InstantiateMsg:  instantiateMsg,
		},
		WasmKeeper: nil, // N.B.: this is set in InitializePool().
	}

	return pool
}

// poolmanager.PoolI interface implementation

// GetAddress returns the address of the concentrated liquidity pool
func (p Pool) GetAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(p.PoolAddress)
	if err != nil {
		panic(fmt.Sprintf("could not bech32 decode address of pool with id: %d", p.GetId()))
	}
	return addr
}

// GetId returns the id of the concentrated liquidity pool
func (p Pool) GetId() uint64 {
	return p.PoolId
}

// String returns the json marshalled string of the pool
func (p Pool) String() string {
	return p.CosmWasmPool.String()
}

// GetSwapFee returns the swap fee of the pool.
func (p Pool) GetSpreadFactor(ctx sdk.Context) sdk.Dec {
	request := msg.GetSwapFeeQueryMsg{}
	response := cosmwasmutils.MustQuery[msg.GetSwapFeeQueryMsg, msg.GetSwapFeeQueryMsgResponse](ctx, p.WasmKeeper, p.ContractAddress, request)
	return response.SwapFee
}

// IsActive returns true if the pool is active
func (p Pool) IsActive(ctx sdk.Context) bool {
	return true
}

// SpotPrice returns the spot price of the pool.
func (p Pool) SpotPrice(ctx sdk.Context, baseAssetDenom string, quoteAssetDenom string) (sdk.Dec, error) {
	request := msg.SpotPriceQueryMsg{
		SpotPrice: msg.SpotPrice{
			QuoteAssetDenom: quoteAssetDenom,
			BaseAssetDenom:  baseAssetDenom,
		},
	}
	response, err := cosmwasmutils.Query[msg.SpotPriceQueryMsg, msg.SpotPriceQueryMsgResponse](ctx, p.WasmKeeper, p.ContractAddress, request)
	if err != nil {
		return sdk.Dec{}, err
	}
	return sdk.MustNewDecFromStr(response.SpotPrice), nil
}

// GetTotalShares returns the total shares of the pool.
func (p Pool) GetTotalShares() sdk.Int {
	panic("TODO: remove from PoolI")
}

// GetType returns the type of the pool.
func (p Pool) GetType() poolmanagertypes.PoolType {
	return poolmanagertypes.CosmWasm
}

// GetTotalPoolLiquidity returns the total pool liquidity
func (p Pool) GetTotalPoolLiquidity(ctx sdk.Context) sdk.Coins {
	request := msg.GetTotalPoolLiquidityQueryMsg{}
	response := cosmwasmutils.MustQuery[msg.GetTotalPoolLiquidityQueryMsg, msg.GetTotalPoolLiquidityQueryMsgResponse](ctx, p.WasmKeeper, p.ContractAddress, request)
	return response.TotalPoolLiquidity
}

// types.CosmWasmExtension interface implementation

func (p Pool) GetCodeId() uint64 {
	return p.CodeId
}

func (p Pool) GetInstantiateMsg() []byte {
	return p.InstantiateMsg
}

func (p Pool) GetContractAddress() string {
	return p.ContractAddress
}

func (p *Pool) SetContractAddress(contractAddress string) {
	p.ContractAddress = contractAddress
}

func (p Pool) GetStoreModel() proto.Message {
	return &p.CosmWasmPool
}

// Set the wasm keeper.
func (p *Pool) SetWasmKeeper(wasmKeeper types.WasmKeeper) {
	p.WasmKeeper = wasmKeeper
}
