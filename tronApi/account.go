package tronApi

type RespGetAccount struct {
	OwnerPermission struct {
		Keys []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		PermissionName string `json:"permission_name"`
	} `json:"owner_permission"`
	AccountResource struct {
		FrozenBalanceForEnergy struct {
			FrozenBalance int   `json:"frozen_balance"`
			ExpireTime    int64 `json:"expire_time"`
		} `json:"frozen_balance_for_energy"`
		LatestConsumeTimeForEnergy int64 `json:"latest_consume_time_for_energy"`
	} `json:"account_resource"`
	ActivePermission []struct {
		Operations string `json:"operations"`
		Keys       []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		Id             int    `json:"id"`
		Type           string `json:"type"`
		PermissionName string `json:"permission_name"`
	} `json:"active_permission"`
	AssetOptimized bool   `json:"asset_optimized"`
	Address        string `json:"address"`
	FrozenSupply   []struct {
		FrozenBalance int   `json:"frozen_balance"`
		ExpireTime    int64 `json:"expire_time"`
	} `json:"frozen_supply"`
	AssetIssuedID     string `json:"asset_issued_ID"`
	CreateTime        int64  `json:"create_time"`
	LatestConsumeTime int64  `json:"latest_consume_time"`
	Frozen            []struct {
		FrozenBalance int   `json:"frozen_balance"`
		ExpireTime    int64 `json:"expire_time"`
	} `json:"frozen"`
	Allowance           int64  `json:"allowance"`
	AssetIssuedName     string `json:"asset_issued_name"`
	LatestOprationTime  int64  `json:"latest_opration_time"`
	FreeAssetNetUsageV2 []struct {
		Value int    `json:"value"`
		Key   string `json:"key"`
	} `json:"free_asset_net_usageV2"`
	IsWitness bool `json:"is_witness"`
	AssetV2   []struct {
		Value int64  `json:"value"`
		Key   string `json:"key"`
	} `json:"assetV2"`
	Asset []struct {
		Value int64  `json:"value"`
		Key   string `json:"key"`
	} `json:"asset"`
	FrozenV2 []struct {
		Type string `json:"type,omitempty"`
	} `json:"frozenV2"`
	Balance               int64 `json:"balance"`
	LatestConsumeFreeTime int64 `json:"latest_consume_free_time"`
	WitnessPermission     struct {
		Keys []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		Id             int    `json:"id"`
		Type           string `json:"type"`
		PermissionName string `json:"permission_name"`
	} `json:"witness_permission"`
}

func (a *RespGetAccount) GetTrc10Balance(tokenId string) (balance int64) {
	for _, assetV2 := range a.AssetV2 {
		if assetV2.Key == tokenId {
			balance = assetV2.Value
			return
		}
	}
	for _, asset := range a.Asset {
		if asset.Key == tokenId {
			balance = asset.Value
			return
		}
	}
	return 0
}
