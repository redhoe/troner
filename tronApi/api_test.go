package tronApi

import (
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"go.uber.org/zap"
	"math/big"
	"testing"
)

var apiEngine *TronApiEngine

var apiUrlNile = "https://api.nileex.io"

func init() {
	log, _ := zap.NewDevelopment()
	apiEngine = NewTronApiEngine("test", apiUrlNile, log)
}

// 构建交易
func TestBuildTx(t *testing.T) {
	req := &ReqCreateTx{
		OwnerAddress: "TYn62fLYjAW7azvkDeMQx9CVV5AqnFUDCH",
		ToAddress:    "TWUjvkTK4aaMZfMNnwuko6rJga3dE5C1Ls",
		Amount:       2250000,
		Visible:      true,
	}
	rep, err := apiEngine.CreateTx(req)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(rep)
}

// 预估能量消耗
func TestTrc20TriggerConstantContract(t *testing.T) {
	addrA, err := address.Base58ToAddress("TXaevvPuifR6TxrDwxK3XA4AN86LGaBHzS")
	if err != nil {
		return
	}
	addrB, err := address.Base58ToAddress("TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj")
	if err != nil {
		return
	}
	t.Log(addrA.Hex())
	t.Log(addrB.Hex())

	req := &ReqTriggerConstantContract{
		OwnerAddress:     addrA.Hex(),
		ContractAddress:  addrB.Hex(),
		FunctionSelector: "balanceOf(address)",
		Parameter:        "",
		Visible:          false,
	}
	req.Parameter, _ = apiEngine.encodeParams("TWUjvkTK4aaMZfMNnwuko6rJga3dE5C1Ls", big.NewInt(123456))

	t.Log(apiEngine.Trc20TriggerConstantContract(req))
}

// 广播交易TRX -- pass
func TestBroadcast(t *testing.T) {
	// tx参数
	req := &ReqCreateTx{
		OwnerAddress: "TYn62fLYjAW7azvkDeMQx9CVV5AqnFUDCH",
		ToAddress:    "TWUjvkTK4aaMZfMNnwuko6rJga3dE5C1Ls",
		Amount:       2233000,
		Visible:      true,
	}
	rep, err := apiEngine.CreateTx(req)
	if err != nil {
		t.Log(err.Error())
	}
	//rep.Signature = make([]string, 0)
	t.Log("广播前hash：", rep.TxID)

	rep.Signature, _ = apiEngine.signHashTx(rep.TxID,
		"66e1c2eb8eef356c8f10fba7b24da187ba25eeaebe487f473a66f7643aef61ee")
	//rep.Signature = append(rep.Signature, signature)
	rep.Visible = true
	hash, err := apiEngine.BroadcastTx(rep)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log("广播后hash：", hash)

}

// 广播交易TRC20
func TestTrc20Broadcast(t *testing.T) {
	// tx参数
	req := &ReqTriggerSmartContract{
		OwnerAddress:     "TYn62fLYjAW7azvkDeMQx9CVV5AqnFUDCH",
		ContractAddress:  "TMMa62GBJsWoTvxvqm4XAHuD2dCuKUUzrn",
		FunctionSelector: "transfer(address,uint256)",
		Parameter:        "",
		CallValue:        0,
		FeeLimit:         15000000,
		Visible:          true,
	}
	//req.Parameter, _ = apiEngine.encodeParams("TWUjvkTK4aaMZfMNnwuko6rJga3dE5C1Ls", 1230000000000000000)
	rep, err := apiEngine.Trc20TriggerSmartContract(req, "TWUjvkTK4aaMZfMNnwuko6rJga3dE5C1Ls", big.NewInt(1230000000000000000))
	if err != nil {
		t.Log(err.Error())
	}
	//t.Log(fmt.Sprintf("%+v", rep))
	t.Log("广播前hash：", rep.Transaction.TxID)
	rep.Transaction.Signature, err = apiEngine.signHashTx(rep.Transaction.TxID,
		"66e1c2eb8eef356c8f10fba7b24da187ba25eeaebe487f473a66f7643aef61ee")
	hash, err := apiEngine.BroadcastTx(rep.Transaction)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log("广播后hash：", hash)
}
