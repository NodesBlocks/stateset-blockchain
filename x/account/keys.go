package account

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	AppAccountKeyPrefix = []byte{0x00}

	JailEndTimeAccountPrefix = []byte{0x10}
)

func key(addr sdk.AccAddress) []byte {
	return append(AppAccountKeyPrefix, addr.Bytes()...)
}

func jailEndTimeAccountsKey(endTime time.Time) []byte {
	return append(JailEndTimeAccountPrefix, sdk.FormatTimeBytes(endTime)...)
}

func jailEndTimeAccountKey(endTime time.Time, addr sdk.AccAddress) []byte {
	return append(jailEndTimeAccountsKey(endTime), addr.Bytes()...)
}