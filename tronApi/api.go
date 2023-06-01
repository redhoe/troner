package tronApi

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type TronApiEngine struct {
	LogEngine *zap.Logger
	BaseUrl   string
	ApiKey    string
}

func NewTronApiEngine(apiUrl, apiKey string, log *zap.Logger) *TronApiEngine {
	return &TronApiEngine{log, apiUrl, apiKey}
}

func (t *TronApiEngine) get(shortUrl string) ([]byte, error) {
	url := t.BaseUrl + shortUrl
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("TRON_PRO_API_KEY", t.ApiKey)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (t *TronApiEngine) post(shortUrl string, jsonStr string) ([]byte, error) {
	url := t.BaseUrl + shortUrl
	payload := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("TRON_PRO_API_KEY", t.ApiKey)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// trc20转账编码
func (t *TronApiEngine) encodeParams(to string, amountBigInt *big.Int) (parameters string, err error) {
	parameters = ""
	err = nil
	addrB, err := address.Base58ToAddress(to)
	if err != nil {
		return
	}
	//amountBig := big.NewInt(amount)
	// trc20BalanceOf +
	//parameters = "0000000000000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-2:] + addrB.Hex()[2:]
	ab := common.LeftPadBytes(amountBigInt.Bytes(), 32)
	parameters = "0000000000000000000000000000000000000000000000000000000000000000"[len(addrB.Hex())-4:] + addrB.Hex()[4:]
	parameters += common.Bytes2Hex(ab)
	return
}

// hash 签名
func (t *TronApiEngine) signHashTx(hash string, privateKey string) (string, error) {
	if hash[:2] != "0x" {
		hash = "0x" + hash
	}
	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}
	hashBytes, err := hexutil.Decode(hash)
	if err != nil {
		return "", err
	}
	si, err := t.signTx(hashBytes, privateKey)
	if err != nil {
		return "", err
	}
	//logger.Infof("sign:%+v", si)
	return hexutil.Encode(si), err

}

func (t *TronApiEngine) signTx(hash []byte, privateKey string) ([]byte, error) {
	b, err := hexutil.Decode("0x" + privateKey)
	if err != nil {
		return []byte(""), err
	}
	//logger.Infof("%+v", b)
	key, err := crypto.ToECDSA(b)
	if err != nil {
		return []byte(""), err
	}
	sig, err := crypto.Sign(hash, key)
	if err != nil {
		return []byte(""), err
	}
	return sig, nil
}
