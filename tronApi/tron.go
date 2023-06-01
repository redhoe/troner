package tronApi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RespNodeInfo struct {
	ActiveConnectCount  int    `json:"activeConnectCount"`
	BeginSyncNum        int64  `json:"beginSyncNum"`
	Block               string `json:"block"`
	CheatWitnessInfoMap struct {
	} `json:"cheatWitnessInfoMap"`
	ConfigNodeInfo struct {
		ActiveNodeSize           int     `json:"activeNodeSize"`
		AllowAdaptiveEnergy      int     `json:"allowAdaptiveEnergy"`
		AllowCreationOfContracts int     `json:"allowCreationOfContracts"`
		BackupListenPort         int     `json:"backupListenPort"`
		BackupMemberSize         int     `json:"backupMemberSize"`
		BackupPriority           int     `json:"backupPriority"`
		CodeVersion              string  `json:"codeVersion"`
		DbVersion                int     `json:"dbVersion"`
		DiscoverEnable           bool    `json:"discoverEnable"`
		ListenPort               int     `json:"listenPort"`
		MaxConnectCount          int     `json:"maxConnectCount"`
		MaxTimeRatio             float64 `json:"maxTimeRatio"`
		MinParticipationRate     int     `json:"minParticipationRate"`
		MinTimeRatio             float64 `json:"minTimeRatio"`
		P2PVersion               string  `json:"p2pVersion"`
		PassiveNodeSize          int     `json:"passiveNodeSize"`
		SameIpMaxConnectCount    int     `json:"sameIpMaxConnectCount"`
		SendNodeSize             int     `json:"sendNodeSize"`
		SupportConstant          bool    `json:"supportConstant"`
		VersionNum               string  `json:"versionNum"`
	} `json:"configNodeInfo"`
	CurrentConnectCount int `json:"currentConnectCount"`
	MachineInfo         struct {
		CpuCount               int           `json:"cpuCount"`
		CpuRate                float64       `json:"cpuRate"`
		DeadLockThreadCount    int           `json:"deadLockThreadCount"`
		DeadLockThreadInfoList []interface{} `json:"deadLockThreadInfoList"`
		FreeMemory             int           `json:"freeMemory"`
		JavaVersion            string        `json:"javaVersion"`
		JvmFreeMemory          int64         `json:"jvmFreeMemory"`
		JvmTotalMemory         int64         `json:"jvmTotalMemory"`
		MemoryDescInfoList     []struct {
			InitSize int     `json:"initSize"`
			MaxSize  int64   `json:"maxSize"`
			Name     string  `json:"name"`
			UseRate  float64 `json:"useRate"`
			UseSize  int     `json:"useSize"`
		} `json:"memoryDescInfoList"`
		OsName         string  `json:"osName"`
		ProcessCpuRate float64 `json:"processCpuRate"`
		ThreadCount    int     `json:"threadCount"`
		TotalMemory    int64   `json:"totalMemory"`
	} `json:"machineInfo"`
	PassiveConnectCount int `json:"passiveConnectCount"`
	PeerList            []struct {
		Active                  bool    `json:"active"`
		AvgLatency              float64 `json:"avgLatency"`
		BlockInPorcSize         int     `json:"blockInPorcSize"`
		ConnectTime             int64   `json:"connectTime"`
		DisconnectTimes         int     `json:"disconnectTimes"`
		HeadBlockTimeWeBothHave int     `json:"headBlockTimeWeBothHave"`
		HeadBlockWeBothHave     string  `json:"headBlockWeBothHave"`
		Host                    string  `json:"host"`
		InFlow                  int     `json:"inFlow"`
		LastBlockUpdateTime     int64   `json:"lastBlockUpdateTime"`
		LastSyncBlock           string  `json:"lastSyncBlock"`
		LocalDisconnectReason   string  `json:"localDisconnectReason"`
		NeedSyncFromPeer        bool    `json:"needSyncFromPeer"`
		NeedSyncFromUs          bool    `json:"needSyncFromUs"`
		NodeCount               int     `json:"nodeCount"`
		NodeId                  string  `json:"nodeId"`
		Port                    int     `json:"port"`
		RemainNum               int     `json:"remainNum"`
		RemoteDisconnectReason  string  `json:"remoteDisconnectReason"`
		Score                   int     `json:"score"`
		SyncBlockRequestedSize  int     `json:"syncBlockRequestedSize"`
		SyncFlag                bool    `json:"syncFlag"`
		SyncToFetchSize         int     `json:"syncToFetchSize"`
		SyncToFetchSizePeekNum  int     `json:"syncToFetchSizePeekNum"`
		UnFetchSynNum           int     `json:"unFetchSynNum"`
	} `json:"peerList"`
	SolidityBlock string `json:"solidityBlock"`
	TotalFlow     int    `json:"totalFlow"`
}

// 部分节点未开放该接口 不推荐使用
func (t *TronApiEngine) GetNodeInfo() (resp *RespNodeInfo, err error) {
	shortUrl := fmt.Sprintf("/wallet/getnodeinfo")
	bodyBytes, err := t.get(shortUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &resp)
	if resp.BeginSyncNum == 0 {
		err = errors.New("apiUrl Error")
	}
	return
}

type RespBlockInfo struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		RawData struct {
			Number         int64  `json:"number"`
			TxTrieRoot     string `json:"txTrieRoot"`
			WitnessAddress string `json:"witness_address"`
			ParentHash     string `json:"parentHash"`
			Version        int64  `json:"version"`
			Timestamp      int64  `json:"timestamp"`
		} `json:"raw_data"`
		WitnessSignature string `json:"witness_signature"`
	} `json:"block_header"`
	Transactions []struct {
		Ret []struct {
			ContractRet string `json:"contractRet"`
		} `json:"ret"`
		Signature []string `json:"signature"`
		TxID      string   `json:"txID"`
		RawData   struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Amount          int64  `json:"amount,omitempty"`
						OwnerAddress    string `json:"owner_address"`
						ToAddress       string `json:"to_address,omitempty"`
						Resource        string `json:"resource,omitempty"`
						FrozenDuration  int    `json:"frozen_duration,omitempty"`
						FrozenBalance   int    `json:"frozen_balance,omitempty"`
						ReceiverAddress string `json:"receiver_address,omitempty"`
						Data            string `json:"data,omitempty"`
						ContractAddress string `json:"contract_address,omitempty"`
						AssetName       string `json:"asset_name,omitempty"`
						AccountAddress  string `json:"account_address,omitempty"`
					} `json:"value"`
					TypeUrl string `json:"type_url"`
				} `json:"parameter"`
				Type         string `json:"type"`
				PermissionId int    `json:"Permission_id,omitempty"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			Timestamp     int64  `json:"timestamp,omitempty"`
			FeeLimit      int    `json:"fee_limit,omitempty"`
			Data          string `json:"data,omitempty"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transactions"`
}

// GetNowBlock 获取当前区块
func (t *TronApiEngine) GetNowBlock() (resp *RespBlockInfo, err error) {
	shortUrl := fmt.Sprintf("/wallet/getnowblock")
	bodyBytes, err := t.get(shortUrl)
	err = json.Unmarshal(bodyBytes, &resp)
	//logger.Infof("%+v", resp)
	//logger.Info(string(bodyBytes))
	return
}

// GetBlockByNum 根据num获取区块
func (t *TronApiEngine) GetBlockByNum(num int64) (resp *RespBlockInfo, err error) {
	shortUrl := fmt.Sprintf("/wallet/getblockbynum")
	data := map[string]interface{}{"num": num}
	jsonBytes, _ := json.Marshal(data)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	err = json.Unmarshal(bodyBytes, &resp)
	//logger.Infof("%+v", resp)
	//logger.Info(fmt.Sprintf("%+v", string(bodyBytes)))
	return
}

type RespTransaction struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature []string `json:"signature"`
	TxID      string   `json:"txID"`
	RawData   struct {
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
		FeeLimit      int    `json:"fee_limit"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	RawDataHex string `json:"raw_data_hex"`
}

// GetTransactionById 通过Hash获取交易信息
func (t *TronApiEngine) GetTransactionById(hash string) (resp *RespTransaction, err error) {
	shortUrl := fmt.Sprintf("/wallet/gettransactionbyid")
	data := map[string]interface{}{"value": hash}
	jsonBytes, _ := json.Marshal(data)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	err = json.Unmarshal(bodyBytes, &resp)
	//logger.Infof("%+v", resp)
	//logger.Info(fmt.Sprintf("%+v", string(bodyBytes)))
	return
}

type RespTransactionInfo struct {
	Id              string   `json:"id"`
	Fee             int      `json:"fee"`
	BlockNumber     int      `json:"blockNumber"`
	BlockTimeStamp  int64    `json:"blockTimeStamp"`
	ContractResult  []string `json:"contractResult"`
	ContractAddress string   `json:"contract_address"`
	Receipt         struct {
		OriginEnergyUsage  int    `json:"origin_energy_usage"`
		NetUsage           int    `json:"net_usage"`
		EnergyUsage        int    `json:"energy_usage"`
		EnergyFee          int    `json:"energy_fee"`
		EnergyUsageTotal   int    `json:"energy_usage_total"`
		NetFee             int    `json:"net_fee"`
		Result             string `json:"result"`
		EnergyPenaltyTotal int    `json:"energy_penalty_total"`
	} `json:"receipt"`
	Log []struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
		Data    string   `json:"data"`
	} `json:"log"`
	PackingFee int `json:"packingFee"`
}

// GetTransactionInfoById 通过Hash获取详细信息
func (t *TronApiEngine) GetTransactionInfoById(hash string) (resp *RespTransactionInfo, err error) {
	shortUrl := fmt.Sprintf("/wallet/gettransactionbyid")
	if hash[:2] == "0x" {
		hash = hash[2:]
	}
	data := map[string]interface{}{"value": hash}
	jsonBytes, _ := json.Marshal(data)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	err = json.Unmarshal(bodyBytes, &resp)
	//logger.Infof("%+v", resp)
	//logger.Info(fmt.Sprintf("%+v", string(bodyBytes)))
	return
}

type RespBroadcastTx struct {
	Txid        string `json:"txID"`
	Result      bool   `json:"result"`
	Code        string `json:"code"`
	Message     string `json:"message"`
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

// BroadcastTx https://cn.developers.tron.network/reference/broadcasttransaction
func (t *TronApiEngine) BroadcastTx(req any) (hash string, err error) {
	shortUrl := fmt.Sprintf("/wallet/broadcasttransaction")
	jsonBytes, _ := json.Marshal(req)
	bodyBytes, err := t.post(shortUrl, string(jsonBytes))
	if err != nil {
		return
	}
	resp := new(RespBroadcastTx)
	err = json.Unmarshal(bodyBytes, &resp)
	hash = resp.Txid
	if !resp.Result {
		err = errors.New(resp.Message)
	}
	return
}
