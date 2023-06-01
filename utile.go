package troner

import (
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/shopspring/decimal"
	"math/big"
)

func CheckAddressIsBase58(addr string) bool {
	_, err := address.Base58ToAddress(addr)
	if err != nil {
		return false
	}
	return true
}

// 数值类型转化

func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)
	return n
}

func BitIntToWeiDecimal(n *big.Int, size int32) decimal.Decimal {
	return decimal.NewFromBigInt(n, 0).Div(decimal.New(1, size))
}

func Int64ToBigInt(n int64) *big.Int {
	return big.NewInt(n)
}
