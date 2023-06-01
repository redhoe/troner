package tronApi

import "testing"

func TestGetNodeInfo(t *testing.T) {
	resp, err := apiEngine.GetNodeInfo()
	if err != nil {
		t.Log(err)
	}
	t.Log(resp.BeginSyncNum)
}

func TestGetLastBlock(t *testing.T) {
	resp, err := apiEngine.GetNowBlock()
	if err != nil {
		t.Log(err)
	}
	t.Log(resp.BlockHeader.RawData.Number)
}

func TestGetBlockByNum(t *testing.T) {
	resp, err := apiEngine.GetBlockByNum(35066070)
	if err != nil {
		t.Log(err)
	}
	t.Log(resp)
}
