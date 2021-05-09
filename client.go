package sideshiftai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Documentation: https://documenter.getpostman.com/

// Client is a sideshift.ai API client
type Client interface {
	// 	Fetch facts, such as assets (currencies), deposit and settle methods
	GetFetchFacts()
	// Fetch current rate/price, min deposit, and max deposit for a shift.
	GetFetchPairs()
	// Retrieves the permissions of the caller, indicating what actions they're allowed to take on SideShift.ai
	GetPermissions()
	// For fixed rate orders, a quote should be requested first.
	// A quote expires after 15 minutes.
	// After the quote request, a fixed rate order should be requested using the id returned by the /quotes endpoint
	PostRequestQuotes()
	// After requesting a quote, use the quoteId to create a fixed rate order with the quote.
	// For fixed rate orders, a deposit of exactly the amount of depositAmount must be made before the expiresAt timestamp, otherwise the deposit will be refunded.
	// For XLM deposits, the API response will contain an additional memo field under depositAddress. The XLM transaction must contain this memo, otherwise the deposit might be lost.
	// affiliateId is optional, but should be defined to get a commission after the shift is complete. It can be obtained here. Commissions are paid in SideShift.ai native XAI token, read more about XAI here.
	// refundAddress is optional, if not defined, user will be prompted to enter a refund address manually on the SideShift.ai order page if the order needs to be refunded
	// orderId field is deprecated, use id instead.
	// To create an order for XAI balance as settle method, the request should include the sessionSecret, settleMethod should be saibal and the settleAddress should be the affiliateId.
	PostCreateFixedOrders()
	// For variable rate orders, the settlement rate is determined when the user's payment is received.
	// For XLM deposits, the API response will contain an additional memo field under depositAddress. The XLM transaction must contain this memo, otherwise the deposit might be lost.
	// affiliateId is optional, but should be defined to get a commission after the shift is complete. It can be obtained here. Commissions are paid in SideShift.ai native XAI token, read more about XAI here.
	// refundAddress is optional, if not defined, user will be prompted to enter a refund address manually on the SideShift.ai order page if the order needs to be refunded
	// orderId field is deprecated, use id instead.
	// To create an order for XAI balance as settle method, the request should include the sessionSecret, settleMethod should be saibal and the settleAddress should be the affiliateId.
	PostCreateVariableOrders()
	// Fetch the order data, including the deposits. orderId field is deprecated, use id instead.
	GetFetchOrders()
}

type client struct {
	httpcl *http.Client
	config *Config
}

// New returns a new sideshift.ai API client
func New(config *Config) *client {
	cfg := &Config{
		APIBaseAddress: APIBaseAddress,
		APIVersion:     APIVersion,
	}
	if len(config.APIBaseAddress) == 0 {
		cfg.APIBaseAddress = APIBaseAddress
	}
	if len(config.APIVersion) == 0 {
		cfg.APIVersion = APIVersion
	}
	return &client{config: cfg}
}

// Helper function
func (c *client) do(method, path string, in, out interface{}) error {
	var payload []byte
	var err error

	payload = nil
	if in != nil {
		payload, err = json.Marshal(in)
		if err != nil {
			return err
		}
	}
	endpoint, err := url.Parse(fmt.Sprintf("%s%s", c.config.APIBaseAddress, c.config.APIVersion))
	if err != nil {
		return nil
	}

	endpoint.Path = fmt.Sprintf("%s/%s", endpoint.Path, path)

	req, err := http.NewRequest(method, fmt.Sprintf("%s", endpoint), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	httpcl := http.DefaultClient
	resp, err := httpcl.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var apierr APIError
		err = json.Unmarshal(body, &apierr)
		if err != nil {
			return err
		}
		return &apierr
	}

	return json.NewDecoder(resp.Body).Decode(out)
}

// Methods
func (c *client) GetFetchFacts() (interface{}, error) {
	var out interface{}
	err := c.do("GET", "facts", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetFetchPairs(pair string) (interface{}, error) {
	var out Pair
	err := c.do("GET", fmt.Sprintf("pairs/%s", pair), nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *client) GetPermissions() (interface{}, error) {
	var out interface{}
	err := c.do("GET", "permissions", nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *client) PostRequestQuotes(req *RequestQuotes) (resp *ResponseQuotes, err error) {
	err = c.do("POST", "quotes", &req, &resp)
	if err != nil {
		return nil, err
	}

	return
}

func (c *client) PostCreateFixedOrders(req *RequestFixedOrders) (resp *ResponseOrders, err error) {
	err = c.do("POST", "orders", &req, &resp)
	if err != nil {
		return nil, err
	}

	return
}

func (c *client) PostCreateVariableOrders(req *RequestVariableOrders) (resp *ResponseOrders, err error) {
	err = c.do("POST", "orders", &req, &resp)
	if err != nil {
		return nil, err
	}

	return
}

func (c *client) GetFetchOrders(orderid string) (resp *ResponseOrders, err error) {
	err = c.do("GET", fmt.Sprintf("orders/%s", orderid), nil, &resp)
	if err != nil {
		return nil, err
	}

	return
}
