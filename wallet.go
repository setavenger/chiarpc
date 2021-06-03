package chiarpc

import (
	"encoding/json"
	"log"
	"net/http"
)

// SendTransaction amount and fee is in mojos
// the address is a xch1.... address
// walletId defaults to 1 if nothing else is specified
func (c Client) SendTransaction(amount uint64, address string, fee uint64, walletId uint64) (map[string]interface{}, error) {
	if walletId == 0 {
		walletId = 1
	}

	data := map[string]interface{}{"amount": amount, "address": address, "fee": fee, "wallet_id": walletId}
	responseRaw, err := c.makeRPCCall(http.MethodPost, "send_transaction", WalletPort, data, nil)
	log.Println(responseRaw)
	if err != nil {
		log.Println(err)
	}
	var respData map[string]interface{}
	err = json.Unmarshal(responseRaw, &respData)
	if err != nil {
		log.Println(err)
	}

	return respData, nil
}
