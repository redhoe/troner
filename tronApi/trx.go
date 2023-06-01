package tronApi

import (
	"encoding/json"
	"fmt"
)

type ReqCreateTx struct {
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
	Amount       int64  `json:"amount"`
	Visible      bool   `json:"visible"`
}

type RespCreateTx struct {
	Visible bool   `json:"visible"`
	TxID    string `json:"txID"`
	RawData struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Amount       int    `json:"amount"`
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
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
	Signature  string `json:"signature"`
}

type TxResp struct {
	Visible bool   `json:"visible"`
	TxID    string `json:"txID"`
	RawData struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Amount       int64  `json:"amount"`
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
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
	Signature  string `json:"signature"`
}

// CreateTx https://cn.developers.tron.network/reference/account-getaccount 创建主币交易
func (t *TronApiEngine) CreateTx(req *ReqCreateTx) (resp *TxResp, err error) {
	shortUrl := fmt.Sprintf("/wallet/createtransaction")
	jsonBytes, _ := json.Marshal(req)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &resp)
	return
}

// GetAccount https://cn.developers.tron.network/reference/account-getaccount 获取账号信息
func (t *TronApiEngine) GetAccount(address string) (resp *RespGetAccount, err error) {
	shortUrl := fmt.Sprintf("/wallet/getaccount")
	data := map[string]interface{}{
		"address": address,
		"visible": true,
	}
	jsonBytes, _ := json.Marshal(data)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &resp)
	if resp.Address == "" {
		err = fmt.Errorf("account not found")
		return
	}
	return
}
