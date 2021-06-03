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
	interimData, err := json.Marshal(responseRaw)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(interimData, &parsedResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &parsedResponse.BlockchainState, nil
}

func (c Client) GetBlock(headerHash string) (map[string]interface{}, error) {
	return nil, nil
}

func (c Client) GetBlocks() (map[string]interface{}, error) {
	return nil, nil
}

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
