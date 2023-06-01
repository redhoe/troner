package tronApi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RespTRC10Issue struct {
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"` // 全称
	Abbr         string `json:"abbr"` // 简写
	TotalSupply  int64  `json:"total_supply"`
	TrxNum       int    `json:"trx_num"`
	Precision    int32  `json:"precision"` // 精度
	Num          int    `json:"num"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	Id           string `json:"id"`
}

func (t *TronApiEngine) GetAssetIssueById(trc10id interface{}) (resp *RespTRC10Issue, err error) {
	shortUrl := fmt.Sprintf("/wallet/getassetissuebyid")
	data := map[string]interface{}{"value": trc10id, "visible": true}
	jsonBytes, _ := json.Marshal(data)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	err = json.Unmarshal(bodyBytes, &resp)
	return
}

type ReqBuildTrc10 struct {
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
	AssetName    string `json:"asset_name"`
	Amount       int64  `json:"amount"`
	Visible      bool   `json:"visible"`
}

type RespTxTrc10 struct {
	Visible bool   `json:"visible"`
	TxID    string `json:"txID"`
	RawData struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Amount       int    `json:"amount"`
					AssetName    string `json:"asset_name"`
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

func (t *TronApiEngine) BuildTrc10(req *ReqBuildTrc10) (resp *RespTxTrc10, err error) {
	shortUrl := fmt.Sprintf("/wallet/transferasset")
	jsonBytes, _ := json.Marshal(req)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	err = json.Unmarshal(bodyBytes, &resp)
	return
}

// GetTrc10Balance 获取指定trc10余额
func (t *TronApiEngine) GetTrc10Balance(address, tokenId string) (balance int64, err error) {
	resp, err := t.GetAccount(address)
	for _, assetV2 := range resp.AssetV2 {
		if assetV2.Key == tokenId {
			balance = assetV2.Value
			return
		}
	}
	for _, asset := range resp.Asset {
		if asset.Key == tokenId {
			balance = asset.Value
			return
		}
	}
	return
}

type RespTrc10Info struct {
	OwnerTAddress string `json:"owner_address"`
	Name          string `json:"name"`
	Abbr          string `json:"abbr"`
	TotalSupply   int64  `json:"total_supply"`
	TrxNum        int    `json:"trx_num"`
	Precision     int    `json:"precision"`
	Num           int    `json:"num"`
	StartTime     int64  `json:"start_time"`
	EndTime       int64  `json:"end_time"`
	Description   string `json:"description"`
	Url           string `json:"url"`
	Id            string `json:"id"`
}

type Trc10Info struct {
	Name    string
	Address string
	Decimal int32
	Symbol  string
}

func (t *TronApiEngine) GetTrc10Info(address string) (trc10Info Trc10Info, err error) {
	resp := new(RespTrc10Info)
	shortUrl := fmt.Sprintf("/wallet/getassetissuebyid")
	data := map[string]interface{}{
		"value":   address,
		"visible": true,
	}
	jsonBytes, _ := json.Marshal(data)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	err = json.Unmarshal(bodyBytes, &resp)
	if resp.Id == "" {
		err = errors.New("account not found")
		return
	}
	trc10Info.Name = resp.Name
	trc10Info.Address = resp.Id
	trc10Info.Decimal = int32(resp.Precision)
	trc10Info.Symbol = resp.Abbr
	return
}
