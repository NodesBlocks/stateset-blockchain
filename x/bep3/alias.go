package bep3

import (
	"github.com/stateset/stateset-blockchain/x/bep3/keeper"
	"github.com/stateset/stateset-blockchain/x/bep3/types"
)


const (
	EventTypeCreateAtomicSwap      = types.EventTypeCreateAtomicSwap
	EventTypeClaimAtomicSwap       = types.EventTypeClaimAtomicSwap
	EventTypeRefundAtomicSwap      = types.EventTypeRefundAtomicSwap
	EventTypeSwapsExpired          = types.EventTypeSwapsExpired
	AttributeValueCategory         = types.AttributeValueCategory
	AttributeKeySender             = types.AttributeKeySender
	AttributeKeyRecipient          = types.AttributeKeyRecipient
	AttributeKeyAtomicSwapID       = types.AttributeKeyAtomicSwapID
	AttributeKeyRandomNumberHash   = types.AttributeKeyRandomNumberHash
	AttributeKeyTimestamp          = types.AttributeKeyTimestamp
	AttributeKeySenderOtherChain   = types.AttributeKeySenderOtherChain
	AttributeKeyExpireHeight       = types.AttributeKeyExpireHeight
	AttributeKeyAmount             = types.AttributeKeyAmount
	AttributeKeyDirection          = types.AttributeKeyDirection
	AttributeKeyClaimSender        = types.AttributeKeyClaimSender
	AttributeKeyRandomNumber       = types.AttributeKeyRandomNumber
	AttributeKeyRefundSender       = types.AttributeKeyRefundSender
	AttributeKeyAtomicSwapIDs      = types.AttributeKeyAtomicSwapIDs
	AttributeExpirationBlock       = types.AttributeExpirationBlock
	ModuleName                     = types.ModuleName
	StoreKey                       = types.StoreKey
	RouterKey                      = types.RouterKey
	QuerierRoute                   = types.QuerierRoute
	DefaultParamspace              = types.DefaultParamspace
	DefaultLongtermStorageDuration = types.DefaultLongtermStorageDuration
	CreateAtomicSwap               = types.CreateAtomicSwap
	ClaimAtomicSwap                = types.ClaimAtomicSwap
	RefundAtomicSwap               = types.RefundAtomicSwap
	CalcSwapID                     = types.CalcSwapID
	Int64Size                      = types.Int64Size
	RandomNumberHashLength         = types.RandomNumberHashLength
	RandomNumberLength             = types.RandomNumberLength
	AddrByteCount                  = types.AddrByteCount
	MaxOtherChainAddrLength        = types.MaxOtherChainAddrLength
	SwapIDLength                   = types.SwapIDLength
	MaxExpectedIncomeLength        = types.MaxExpectedIncomeLength
	QueryGetAssetSupply            = types.QueryGetAssetSupply
	QueryGetAssetSupplies          = types.QueryGetAssetSupplies
	QueryGetAtomicSwap             = types.QueryGetAtomicSwap
	QueryGetAtomicSwaps            = types.QueryGetAtomicSwaps
	QueryGetParams                 = types.QueryGetParams
	NULL                           = types.NULL
	Open                           = types.Open
	Completed                      = types.Completed
	Expired                        = types.Expired
	INVALID                        = types.INVALID
	Incoming                       = types.Incoming
	Outgoing                       = types.Outgoing
)

var (
	// functions aliases
	NewKeeper                  = keeper.NewKeeper
	NewQuerier                 = keeper.NewQuerier
	NewAssetSupply             = types.NewAssetSupply
	RegisterCodec              = types.RegisterCodec
	NewGenesisState            = types.NewGenesisState
	DefaultGenesisState        = types.DefaultGenesisState
	GenerateSecureRandomNumber = types.GenerateSecureRandomNumber
	CalculateRandomHash        = types.CalculateRandomHash
	CalculateSwapID            = types.CalculateSwapID
	GetAtomicSwapByHeightKey   = types.GetAtomicSwapByHeightKey
	NewMsgCreateAtomicSwap     = types.NewMsgCreateAtomicSwap
	NewMsgClaimAtomicSwap      = types.NewMsgClaimAtomicSwap
	NewMsgRefundAtomicSwap     = types.NewMsgRefundAtomicSwap
	NewParams                  = types.NewParams
	DefaultParams              = types.DefaultParams
	NewAssetParam              = types.NewAssetParam
	ParamKeyTable              = types.ParamKeyTable
	NewQueryAssetSupply        = types.NewQueryAssetSupply
	NewQueryAssetSupplies      = types.NewQueryAssetSupplies
	NewQueryAtomicSwapByID     = types.NewQueryAtomicSwapByID
	NewQueryAtomicSwaps        = types.NewQueryAtomicSwaps
	NewAtomicSwap              = types.NewAtomicSwap
	NewSwapStatusFromString    = types.NewSwapStatusFromString
	NewSwapDirectionFromString = types.NewSwapDirectionFromString
	NewAugmentedAtomicSwap     = types.NewAugmentedAtomicSwap

	// variable aliases
	ModuleCdc                       = types.ModuleCdc
	ErrInvalidTimestamp             = types.ErrInvalidTimestamp
	ErrInvalidHeightSpan            = types.ErrInvalidHeightSpan
	ErrInsufficientAmount           = types.ErrInsufficientAmount
	ErrAssetNotSupported            = types.ErrAssetNotSupported
	ErrAssetNotActive               = types.ErrAssetNotActive
	ErrAssetSupplyNotFound          = types.ErrAssetSupplyNotFound
	ErrExceedsSupplyLimit           = types.ErrExceedsSupplyLimit
	ErrExceedsAvailableSupply       = types.ErrExceedsAvailableSupply
	ErrInvalidCurrentSupply         = types.ErrInvalidCurrentSupply
	ErrInvalidIncomingSupply        = types.ErrInvalidIncomingSupply
	ErrInvalidOutgoingSupply        = types.ErrInvalidOutgoingSupply
	ErrInvalidClaimSecret           = types.ErrInvalidClaimSecret
	ErrAtomicSwapAlreadyExists      = types.ErrAtomicSwapAlreadyExists
	ErrAtomicSwapNotFound           = types.ErrAtomicSwapNotFound
	ErrSwapNotRefundable            = types.ErrSwapNotRefundable
	ErrSwapNotClaimable             = types.ErrSwapNotClaimable
	ErrInvalidAmount                = types.ErrInvalidAmount
	ErrInvalidSwapAccount           = types.ErrInvalidSwapAccount
	AtomicSwapKeyPrefix             = types.AtomicSwapKeyPrefix
	AtomicSwapByBlockPrefix         = types.AtomicSwapByBlockPrefix
	AtomicSwapLongtermStoragePrefix = types.AtomicSwapLongtermStoragePrefix
	AtomicSwapCoinsAccAddr          = types.AtomicSwapCoinsAccAddr
	KeyAssetParams                  = types.KeyAssetParams
	DefaultBnbDeputyFixedFee        = types.DefaultBnbDeputyFixedFee
	DefaultMinAmount                = types.DefaultMinAmount
	DefaultMaxAmount                = types.DefaultMaxAmount
	DefaultMinBlockLock             = types.DefaultMinBlockLock
	DefaultMaxBlockLock             = types.DefaultMaxBlockLock
	DefaultPreviousBlockTime        = types.DefaultPreviousBlockTime
	ModulePermissionsUpgradeTime    = types.ModulePermissionsUpgradeTime
)

type (
	Keeper               = keeper.Keeper
	AssetSupply          = types.AssetSupply
	AssetSupplies        = types.AssetSupplies
	GenesisState         = types.GenesisState
	MsgCreateAtomicSwap  = types.MsgCreateAtomicSwap
	MsgClaimAtomicSwap   = types.MsgClaimAtomicSwap
	MsgRefundAtomicSwap  = types.MsgRefundAtomicSwap
	Params               = types.Params
	AssetParam           = types.AssetParam
	AssetParams          = types.AssetParams
	QueryAssetSupply     = types.QueryAssetSupply
	QueryAssetSupplies   = types.QueryAssetSupplies
	QueryAtomicSwapByID  = types.QueryAtomicSwapByID
	QueryAtomicSwaps     = types.QueryAtomicSwaps
	AtomicSwap           = types.AtomicSwap
	AtomicSwaps          = types.AtomicSwaps
	SwapStatus           = types.SwapStatus
	SwapDirection        = types.SwapDirection
	SupplyLimit          = types.SupplyLimit
	AugmentedAtomicSwap  = types.AugmentedAtomicSwap
	AugmentedAtomicSwaps = types.AugmentedAtomicSwaps
)