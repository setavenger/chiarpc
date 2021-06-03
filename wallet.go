package chiarpc

import (
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

	data := struct {
		Amount   uint64 `json:"amount"`
		Address  string `json:"address"`
		Fee      uint64 `json:"fee"`
		WalletId uint64 `json:"wallet_id"`
	}{
		Amount:   amount,
		Address:  address,
		Fee:      fee,
		WalletId: walletId,
	}

	responseRaw, err := c.makeRPCCall(http.MethodPost, "send_transaction", WalletPort, data, nil)
	log.Println(responseRaw)
	if err != nil {
		log.Println(err)
	}
	return responseRaw, nil
}

type SendTransactionData struct {
}
