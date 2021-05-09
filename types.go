package sideshiftai

type Pair struct {
	Max  string `json:"max"`
	Min  string `json:"min"`
	Rate string `json:"rate"`
}
type Address struct {
	Address string `json:"address"`
}
type SettleTx struct {
	Type   string `json:"type"`
	TxHash string `json:"txHash"`
}
type Deposit struct {
	DepositId        string   `json:"depositId"`
	CreatedAt        string   `json:"createdAt"`
	DepositAmount    string   `json:"depositAmount"`
	SettleRate       string   `json:"settleRate"`
	SettleAmount     string   `json:"settleAmount"`
	NetworkFeeAmount string   `json:"networkFeeAmount"`
	Status           string   `json:"status"`
	SettleTx         SettleTx `json:"settleTx"`
	RefundAddress    string   `json:"refundAddress"`
	RefundTx         string   `json:"refundTx"`
	Reason           string   `json:"reason"`
}

// RequestQuotes is the request for RequestQuotes()
type RequestQuotes struct {
	DepositMethod string `json:"depositMethod"`
	SettleMethod  string `json:"settleMethod"`
	DepositAmount string `json:"depositAmount"`
}

// ResponseQuotes is the response from RequestQuotes()
type ResponseQuotes struct {
	CreatedAt     string `json:"createdAt"`
	DepositAmount string `json:"depositAmount"`
	DepositMethod string `json:"depositMethod"`
	ExpiresAt     string `json:"expiresAt"`
	ID            string `json:"id"`
	Rate          string `json:"rate"`
	SettleAmount  string `json:"settleAmount"`
	SettleMethod  string `json:"settleMethod"`
}

// RequestFixedOrders is the request for RequestFixedOrders()
type RequestFixedOrders struct {
	Type          string `json:"type"`
	QuoteId       string `json:"quoteId"`
	SettleAddress string `json:"settleAddress"`
	AffiliateId   string `json:"affiliateId"`
	RefundAddress string `json:"refundAddress"`
}

// RequestVariableOrders is the request for RequestOrders()
type RequestVariableOrders struct {
	Type            string `json:"type"`
	DepositMethodId string `json:"depositMethodId"`
	SettleMethodId  string `json:"settleMethodId"`
	SettleAddress   string `json:"settleAddress"`
	AffiliateId     string `json:"affiliateId"`
	RefundAddress   string `json:"refundAddress"`
}

// ResponseOrders is the response from RequestOrders()
type ResponseOrders struct {
	CreatedAt       string    `json:"createdAt"`
	CreatedAtISO    string    `json:"createdAtISO"`
	ExpiresAt       string    `json:"expiresAt"`
	ExpiresAtISO    string    `json:"expiresAtISO"`
	DepositAddress  Address   `json:"depositAddress"`
	DepositMethodId string    `json:"depositMethodId"`
	ID              string    `json:"id"`
	OrderId         string    `json:"orderId"`
	SettleAddress   Address   `json:"settleAddress"`
	SettleMethodId  string    `json:"settleMethodId"`
	DepositMax      string    `json:"depositMax"`
	DepositMin      string    `json:"depositMin"`
	QuoteId         string    `json:"quoteId"`
	SettleAmount    string    `json:"settleAmount"`
	DepositAmount   string    `json:"depositAmount"`
	Deposits        []Deposit `json:"deposits"`
}

// ResponsePairs is the response from RequestPairs()
type ResponsePairs struct {
	Rate int    `json:"rate"`
	Min  string `json:"min"`
	Max  string `json:"max"`
}
