package tronApi

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"math/big"
	"unicode/utf8"
)

type ReqTriggerConstantContract struct {
	OwnerAddress     string `json:"owner_address"`
	ContractAddress  string `json:"contract_address"`
	FunctionSelector string `json:"function_selector"`
	Parameter        string `json:"parameter"`
	Visible          bool   `json:"visible"`
}

type RespTrc20TriggerConstantContract struct {
	Result         Result   `json:"result"`
	EnergyUsed     int64    `json:"energy_used"`
	ConstantResult []string `json:"constant_result"`
	Transaction    struct {
		Ret []struct {
		} `json:"ret"`
		Visible bool   `json:"visible"`
		TxID    string `json:"txID"`
		RawData struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Data            string `json:"data"`
						OwnerAddress    string `json:"owner_address"`
						ContractAddress string `json:"contract_address"`
					} `json:"value"`
					TypeUrl string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transaction"`
}

// Trc20TriggerConstantContract 获取Trc20余额

func (t *TronApiEngine) Trc20TriggerConstantContract(req *ReqTriggerConstantContract) (resp RespTrc20TriggerConstantContract, err error) {
	shortUrl := fmt.Sprintf("/wallet/triggerconstantcontract")
	jsonBytes, _ := json.Marshal(req)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &resp)
	return
}

type ReqTriggerSmartContract struct {
	ContractAddress  string `json:"contract_address"`
	OwnerAddress     string `json:"owner_address"`
	FunctionSelector string `json:"function_selector"`
	Parameter        string `json:"parameter"`
	CallValue        int64  `json:"call_value"`
	FeeLimit         int64  `json:"fee_limit"`
	Visible          bool   `json:"visible"`
}

type RespTriggerSmartContract struct {
	Result struct {
		Result bool `json:"result"`
	} `json:"result"`
	Transaction struct {
		Visible bool   `json:"visible"`
		TxID    string `json:"txID"`
		RawData struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Data            string `json:"data"`
						OwnerAddress    string `json:"owner_address"`
						ContractAddress string `json:"contract_address"`
						CallValue       int    `json:"call_value"`
					} `json:"value"`
					TypeUrl string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			FeeLimit      int    `json:"fee_limit"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
		Signature  string `json:"signature"`
	} `json:"transaction"`
}

// Trc20TriggerSmartContract 构建trc20交易
func (t *TronApiEngine) Trc20TriggerSmartContract(req *ReqTriggerSmartContract, toAddress string, tokenValue *big.Int) (resp RespTriggerSmartContract, err error) {
	shortUrl := fmt.Sprintf("/wallet/triggersmartcontract")
	req.Parameter, _ = t.encodeParams(toAddress, tokenValue)
	jsonBytes, _ := json.Marshal(req)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &resp)
	return
}

// balanceOf(address)
func (t *TronApiEngine) Trc20Balance(address, contract string) (nums *big.Int, err error) {
	req := &ReqTriggerConstantContract{
		OwnerAddress:     address,
		ContractAddress:  contract,
		FunctionSelector: "balanceOf(address)",
		Parameter:        "",
		Visible:          true,
	}
	req.Parameter, _ = t.encodeBalanceParams(address)
	resp, err := t.Trc20TriggerConstantContract(req)
	if err != nil {
		return nil, err
	}
	if !resp.Result.Result {
		return nil, errors.New(resp.Result.Message)
	}
	amountBigInt, err := t.parseTRC20NumericProperty(resp.ConstantResult[0])
	if err != nil {
		return nil, err
	}

	return amountBigInt, nil

}

// trc20 查询余额编码
func (t *TronApiEngine) encodeBalanceParams(base58Address string) (parameters string, err error) {
	parameters = ""
	err = nil
	addrB, err := address.Base58ToAddress(base58Address)
	if err != nil {
		return
	}
	parameters = "0000000000000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-2:] + addrB.Hex()[2:]
	return
}

func (t *TronApiEngine) parseTRC20NumericProperty(data string) (*big.Int, error) {
	if common.Has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) == 64 {
		var n big.Int
		_, ok := n.SetString(data, 16)
		if ok {
			return &n, nil
		}
	}
	return nil, fmt.Errorf("Cannot parse %s", data)
}

type Trc20Info struct {
	Name    string
	Address string
	Decimal int32
	Symbol  string
}

// decimals() name() symbol()
func (t *TronApiEngine) Trc20Info(contract string) (trc20Info Trc20Info, err error) {
	trc20Info = Trc20Info{}
	trc20Info.Address = contract
	addressObj, err := address.Base58ToAddress(contract)
	req := &ReqTriggerConstantContract{
		OwnerAddress:     "410000000000000000000000000000000000000000",
		ContractAddress:  addressObj.Hex(),
		FunctionSelector: "decimals()",
		Parameter:        "",
		Visible:          false,
	}
	//req.Parameter, _ = apiEngine.encodeBalanceParams(address)
	resp, err := t.Trc20TriggerConstantContract(req)
	if err != nil {
		return
	}
	if !resp.Result.Result {
		return trc20Info, errors.New(resp.Result.Message)
	}
	decimals, err := t.parseTRC20NumericProperty(resp.ConstantResult[0])
	if err != nil {
		return trc20Info, err
	}
	trc20Info.Decimal = int32(decimals.Int64())

	req.FunctionSelector = "name()"
	respName, err := t.Trc20TriggerConstantContract(req)
	if err != nil {
		return
	}
	if !respName.Result.Result {
		return trc20Info, errors.New(respName.Result.Message)
	}
	name, err := t.parseTRC20StringProperty(respName.ConstantResult[0])
	if err != nil {
		return trc20Info, err
	}
	trc20Info.Name = name

	req.FunctionSelector = "symbol()"

	respSymbol, err := t.Trc20TriggerConstantContract(req)
	if err != nil {
		return
	}
	if !respSymbol.Result.Result {
		return trc20Info, errors.New(respSymbol.Result.Message)
	}
	symbol, err := t.parseTRC20StringProperty(respSymbol.ConstantResult[0])
	if err != nil {
		return trc20Info, err
	}
	trc20Info.Symbol = symbol
	return trc20Info, nil

}

func (t *TronApiEngine) parseTRC20StringProperty(data string) (string, error) {
	if common.Has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) > 128 {
		n, _ := t.parseTRC20NumericProperty(data[64:128])
		if n != nil {
			l := n.Uint64()
			if 2*int(l) <= len(data)-128 {
				b, err := hex.DecodeString(data[128 : 128+2*l])
				if err == nil {
					return string(b), nil
				}
			}
		}
	} else if len(data) == 64 {
		// allow string properties as 32 bytes of UTF-8 data
		b, err := hex.DecodeString(data)
		if err == nil {
			i := bytes.Index(b, []byte{0})
			if i > 0 {
				b = b[:i]
			}
			if utf8.Valid(b) {
				return string(b), nil
			}
		}
	}
	return "", fmt.Errorf("Cannot parse %s,", data)
}

func (t *TronApiEngine) Trc20Gas(from, to, contract string, amount *big.Int) (energyUsed int64, err error) {
	energyUsed = 0
	contractAddress, err := address.Base58ToAddress(contract)
	fromAddress, _ := address.Base58ToAddress(from)
	toAddress, _ := address.Base58ToAddress(to)

	req := &ReqTriggerConstantContract{
		OwnerAddress:     fromAddress.Hex(),
		ContractAddress:  contractAddress.Hex(),
		FunctionSelector: "transfer(address,uint256)",
		Parameter:        "",
		Visible:          false,
	}
	req.Parameter, _ = t.encodeParams(toAddress.Hex(), amount)
	//req.Parameter, _ = apiEngine.encodeBalanceParams(address)
	resp, err := t.Trc20TriggerConstantContract(req)
	if err != nil {
		return
	}
	if !resp.Result.Result {
		return energyUsed, errors.New(resp.Result.Message)
	}

	//gas, err := t.parseTRC20NumericProperty(resp.ConstantResult[0])
	//if err != nil {
	//	return
	//}

	energyUsed = resp.EnergyUsed
	return

}
