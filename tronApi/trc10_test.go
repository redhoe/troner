package tronApi

import "testing"

func TestGetAssetIssueById(t *testing.T) {
	resp, err := apiEngine.GetAssetIssueById("1000016")
	if err != nil {
		t.Log(err)
	}
	t.Log(resp.Name)
	t.Log(resp.Abbr)
	t.Log(resp.Precision)
	t.Log(resp.Id)
	t.Log(resp.Description)
}

// 广播交易TRX -- pass
func TestBroadcastTrc10(t *testing.T) {
	// tx参数
	req := &ReqBuildTrc10{
		OwnerAddress: "",
		ToAddress:    "",
		AssetName:    "1000016", //   string `json:"asset_name"`
		Amount:       2233000,
		Visible:      true,
	}
	rep, err := apiEngine.BuildTrc10(req)
	if err != nil {
		t.Log(err.Error())
	}
	//rep.Signature = make([]string, 0)
	t.Log("广播前hash：", rep.TxID)

	rep.Signature, _ = apiEngine.signHashTx(rep.TxID,
		"")
	//rep.Signature = append(rep.Signature, signature)
	rep.Visible = true
	hash, err := apiEngine.BroadcastTx(rep)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log("广播后hash：", hash)

}
