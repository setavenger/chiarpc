package chiarpc

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c Client) GetBlockchainState() (*BlockchainState, error) {
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_blockchain_state", FullNodePort, nil, nil)
	if err != nil {
		log.Println(err)
	}
	var parsedResponse struct {
		Success         bool            `json:"success"`
		BlockchainState BlockchainState `json:"blockchain_state"`
	}

	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &parsedResponse.BlockchainState, nil
}

func (c Client) GetBlock(headerHash string) (*Block, error) {
	data := map[string]interface{}{"header_hash": headerHash}
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_block", FullNodePort, data, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var parsedResponse struct {
		Success bool  `json:"success"`
		Block   Block `json:"block"`
	}
	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return &parsedResponse.Block, nil
}

func (c Client) GetBlocks(start, end uint64, excludeHeaderHash bool) (*[]Block, error) {
	data := map[string]interface{}{"start": start, "end": end, "exclude_header_hash": excludeHeaderHash}
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_blocks", FullNodePort, data, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var parsedResponse struct {
		Success bool    `json:"success"`
		Blocks  []Block `json:"blocks"`
	}
	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return &parsedResponse.Blocks, nil
}

func (c Client) GetBlockRecordByHeight(height uint64) (*BlockRecord, error) {
	data := map[string]interface{}{"height": height}
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_block_record_by_height", FullNodePort, data, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var parsedResponse struct {
		Success     bool        `json:"success"`
		BlockRecord BlockRecord `json:"block_record"`
	}
	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return &parsedResponse.BlockRecord, nil
}

func (c Client) GetBlockRecord(headerHash string) (*BlockRecord, error) {
	data := map[string]interface{}{"header_hash": headerHash}
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_block_record", FullNodePort, data, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var parsedResponse struct {
		Success     bool        `json:"success"`
		BlockRecord BlockRecord `json:"block_record"`
	}
	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return &parsedResponse.BlockRecord, nil
}

func (c Client) GetBlockRecords(start, end uint64) (*[]BlockRecord, error) {
	data := map[string]interface{}{"start": start, "end": end}
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_block_records", FullNodePort, data, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var parsedResponse struct {
		Success      bool          `json:"success"`
		BlockRecords []BlockRecord `json:"block_records"`
	}
	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return &parsedResponse.BlockRecords, nil
}

func (c Client) GetUnfinishedBlockHeaders() (*[]string, error) {
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_unfinished_block_headers", FullNodePort, nil, nil)
	if err != nil {
		log.Println(err)
	}
	var parsedResponse struct {
		Success bool     `json:"success"`
		Headers []string `json:"headers"`
	}
	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return &parsedResponse.Headers, nil
}

func (c Client) GetNetworkSpace(olderBlockHeaderHash, newerBlockHeaderHash string) (*interface{}, error) {
	//older_block_header_hash
	//newer_block_header_hash
	data := map[string]interface{}{"older_block_header_hash": olderBlockHeaderHash, "newer_block_header_hash": newerBlockHeaderHash}
	responseRaw, err := c.makeRPCCall(http.MethodPost, "get_network_space", FullNodePort, data, nil)
	if err != nil {
		log.Println(err)
	}
	var parsedResponse struct {
		Success bool        `json:"success"`
		Space   interface{} `json:"space,int"`
	}
	err = json.Unmarshal(responseRaw, &parsedResponse)
	if err != nil {
		log.Println(err)
	}
	return &parsedResponse.Space, nil
}

//get_unfinished_block_headers
type BlockchainState struct {
	Difficulty                  uint64 `json:"difficulty"`
	GenesisChallengeInitialized bool   `json:"genesis_challenge_initialized"`
	MempoolSize                 uint64 `json:"mempool_size"`
	Peak                        struct {
		SignagePointIndex                  int64       `json:"signage_point_index"`
		Fees                               interface{} `json:"fees"`
		PrevTransactionBlockHash           interface{} `json:"prev_transaction_block_hash"`
		FarmerPuzzleHash                   string      `json:"farmer_puzzle_hash"`
		PrevHash                           string      `json:"prev_hash"`
		FinishedChallengeSlotHashes        interface{} `json:"finished_challenge_slot_hashes"`
		PoolPuzzleHash                     string      `json:"pool_puzzle_hash"`
		PrevTransactionBlockHeight         int64       `json:"prev_transaction_block_height"`
		HeaderHash                         string      `json:"header_hash"`
		Overflow                           bool        `json:"overflow"`
		FinishedInfusedChallengeSlotHashes interface{} `json:"finished_infused_challenge_slot_hashes"`
		InfusedChallengeVdfOutput          struct {
			Data string `json:"data"`
		} `json:"infused_challenge_vdf_output"`
		ChallengeVdfOutput struct {
			Data string `json:"data"`
		} `json:"challenge_vdf_output"`
		Deficit                    int64       `json:"deficit"`
		FinishedRewardSlotHashes   interface{} `json:"finished_reward_slot_hashes"`
		Height                     uint64      `json:"height"`
		Timestamp                  uint64      `json:"timestamp"`
		SubEpochSummaryIncluded    interface{} `json:"sub_epoch_summary_included"`
		RewardInfusionNewChallenge string      `json:"reward_infusion_new_challenge"`
		RequiredIters              interface{} `json:"required_iters"`
		TotalIters                 interface{} `json:"total_iters"`
		RewardClaimsIncorporated   interface{} `json:"reward_claims_incorporated"`
		Weight                     int64       `json:"weight"`
		ChallengeBlockInfoHash     string      `json:"challenge_block_info_hash"`
		SubSlotIters               interface{} `json:"sub_slot_iters"`
	} `json:"peak"`
	Space        interface{} `json:"space"`
	SubSlotIters int64       `json:"sub_slot_iters"`
	Sync         struct {
		SyncMode           bool  `json:"sync_mode"`
		SyncProgressHeight int64 `json:"sync_progress_height"`
		Synced             bool  `json:"synced"`
		SyncTipHeight      int64 `json:"sync_tip_height"`
	} `json:"sync"`
}

type Block struct {
	RewardChainBlock struct {
		SignagePointIndex          int `json:"signage_point_index"`
		InfusedChallengeChainIpVdf struct {
			Output struct {
				Data string `json:"data"`
			} `json:"output"`
			NumberOfIterations int    `json:"number_of_iterations"`
			Challenge          string `json:"challenge"`
		} `json:"infused_challenge_chain_ip_vdf"`
		RewardChainSpVdf struct {
			Output struct {
				Data string `json:"data"`
			} `json:"output"`
			NumberOfIterations int    `json:"number_of_iterations"`
			Challenge          string `json:"challenge"`
		} `json:"reward_chain_sp_vdf"`
		ProofOfSpace struct {
			PoolContractPuzzleHash interface{} `json:"pool_contract_puzzle_hash"`
			PlotPublicKey          string      `json:"plot_public_key"`
			Size                   int         `json:"size"`
			Challenge              string      `json:"challenge"`
			Proof                  string      `json:"proof"`
			PoolPublicKey          string      `json:"pool_public_key"`
		} `json:"proof_of_space"`
		TotalIters          int64 `json:"total_iters"`
		Weight              int   `json:"weight"`
		ChallengeChainSpVdf struct {
			Output struct {
				Data string `json:"data"`
			} `json:"output"`
			NumberOfIterations int    `json:"number_of_iterations"`
			Challenge          string `json:"challenge"`
		} `json:"challenge_chain_sp_vdf"`
		PosSsCcChallengeHash      string `json:"pos_ss_cc_challenge_hash"`
		ChallengeChainSpSignature string `json:"challenge_chain_sp_signature"`
		ChallengeChainIpVdf       struct {
			Output struct {
				Data string `json:"data"`
			} `json:"output"`
			NumberOfIterations int    `json:"number_of_iterations"`
			Challenge          string `json:"challenge"`
		} `json:"challenge_chain_ip_vdf"`
		IsTransactionBlock bool `json:"is_transaction_block"`
		RewardChainIpVdf   struct {
			Output struct {
				Data string `json:"data"`
			} `json:"output"`
			NumberOfIterations int    `json:"number_of_iterations"`
			Challenge          string `json:"challenge"`
		} `json:"reward_chain_ip_vdf"`
		RewardChainSpSignature string `json:"reward_chain_sp_signature"`
		Height                 int    `json:"height"`
	} `json:"reward_chain_block"`
	RewardChainIpProof struct {
		Witness              string `json:"witness"`
		WitnessType          int    `json:"witness_type"`
		NormalizedToIDentity bool   `json:"normalized_to_identity"`
	} `json:"reward_chain_ip_proof"`
	TransactionsGenerator        interface{}   `json:"transactions_generator"`
	FinishedSubSlots             []interface{} `json:"finished_sub_slots"`
	TransactionsGeneratorRefList []interface{} `json:"transactions_generator_ref_list"`
	HeaderHash                   string        `json:"header_hash"`
	InfusedChallengeChainIpProof struct {
		Witness              string `json:"witness"`
		WitnessType          int    `json:"witness_type"`
		NormalizedToIDentity bool   `json:"normalized_to_identity"`
	} `json:"infused_challenge_chain_ip_proof"`
	FoliageTransactionBlock interface{} `json:"foliage_transaction_block"`
	Foliage                 struct {
		PrevBlockHash                    string      `json:"prev_block_hash"`
		FoliageBlockDataSignature        string      `json:"foliage_block_data_signature"`
		FoliageTransactionBlockSignature interface{} `json:"foliage_transaction_block_signature"`
		FoliageBlockData                 struct {
			ExtensionData          string `json:"extension_data"`
			FarmerRewardPuzzleHash string `json:"farmer_reward_puzzle_hash"`
			PoolTarget             struct {
				MaxHeight  int    `json:"max_height"`
				PuzzleHash string `json:"puzzle_hash"`
			} `json:"pool_target"`
			UnfinishedRewardBlockHash string `json:"unfinished_reward_block_hash"`
			PoolSignature             string `json:"pool_signature"`
		} `json:"foliage_block_data"`
		FoliageTransactionBlockHash interface{} `json:"foliage_transaction_block_hash"`
		RewardBlockHash             string      `json:"reward_block_hash"`
	} `json:"foliage"`
	ChallengeChainIpProof struct {
		Witness              string `json:"witness"`
		WitnessType          int    `json:"witness_type"`
		NormalizedToIDentity bool   `json:"normalized_to_identity"`
	} `json:"challenge_chain_ip_proof"`
	TransactionsInfo   interface{} `json:"transactions_info"`
	RewardChainSpProof struct {
		Witness              string `json:"witness"`
		WitnessType          int    `json:"witness_type"`
		NormalizedToIDentity bool   `json:"normalized_to_identity"`
	} `json:"reward_chain_sp_proof"`
	ChallengeChainSpProof struct {
		Witness              string `json:"witness"`
		WitnessType          int    `json:"witness_type"`
		NormalizedToIDentity bool   `json:"normalized_to_identity"`
	} `json:"challenge_chain_sp_proof"`
}

type BlockRecord struct {
	SignagePointIndex                  int         `json:"signage_point_index"`
	Fees                               interface{} `json:"fees"`
	PrevTransactionBlockHash           interface{} `json:"prev_transaction_block_hash"`
	FarmerPuzzleHash                   string      `json:"farmer_puzzle_hash"`
	PrevHash                           string      `json:"prev_hash"`
	FinishedChallengeSlotHashes        interface{} `json:"finished_challenge_slot_hashes"`
	PoolPuzzleHash                     string      `json:"pool_puzzle_hash"`
	PrevTransactionBlockHeight         int         `json:"prev_transaction_block_height"`
	HeaderHash                         string      `json:"header_hash"`
	Overflow                           bool        `json:"overflow"`
	FinishedInfusedChallengeSlotHashes interface{} `json:"finished_infused_challenge_slot_hashes"`
	InfusedChallengeVdfOutput          struct {
		Data string `json:"data"`
	} `json:"infused_challenge_vdf_output"`
	ChallengeVdfOutput struct {
		Data string `json:"data"`
	} `json:"challenge_vdf_output"`
	Deficit                    int         `json:"deficit"`
	FinishedRewardSlotHashes   interface{} `json:"finished_reward_slot_hashes"`
	Height                     int         `json:"height"`
	Timestamp                  interface{} `json:"timestamp"`
	SubEpochSummaryIncluded    interface{} `json:"sub_epoch_summary_included"`
	RewardInfusionNewChallenge string      `json:"reward_infusion_new_challenge"`
	RequiredIters              int         `json:"required_iters"`
	TotalIters                 int64       `json:"total_iters"`
	RewardClaimsIncorporated   interface{} `json:"reward_claims_incorporated"`
	Weight                     int         `json:"weight"`
	ChallengeBlockInfoHash     string      `json:"challenge_block_info_hash"`
	SubSlotIters               int         `json:"sub_slot_iters"`
}
