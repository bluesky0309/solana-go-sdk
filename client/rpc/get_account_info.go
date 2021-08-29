package rpc

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetAccountInfoConfigEncoding is account's data encode format
type GetAccountInfoConfigEncoding string

const (
	// GetAccountInfoConfigEncodingBase58 limited to Account data of less than 128 bytes
	GetAccountInfoConfigEncodingBase58     GetAccountInfoConfigEncoding = "base58"
	GetAccountInfoConfigEncodingJsonParsed GetAccountInfoConfigEncoding = "jsonParsed"
	GetAccountInfoConfigEncodingBase64     GetAccountInfoConfigEncoding = "base64"
	GetAccountInfoConfigEncodingBase64Zstd GetAccountInfoConfigEncoding = "base64+zstd"
)

// GetAccountInfoConfig is an option config for `getAccountInfo`
type GetAccountInfoConfig struct {
	Commitment Commitment                     `json:"commitment,omitempty"`
	Encoding   GetAccountInfoConfigEncoding   `json:"encoding,omitempty"`
	DataSlice  *GetAccountInfoConfigDataSlice `json:"dataSlice,omitempty"`
}

// GetAccountInfoResponse is a full raw rpc response of `getAccountInfo`
type GetAccountInfoResponse struct {
	GeneralResponse
	Result GetAccountInfoResult `json:"result"`
}

// GetAccountInfoConfigDataSlice is a part of GetAccountInfoConfig
type GetAccountInfoConfigDataSlice struct {
	Offset uint64 `json:"offset,omitempty"`
	Length uint64 `json:"length,omitempty"`
}

// GetAccountInfoResult is rpc result of `getAccountInfo`
type GetAccountInfoResult struct {
	Context Context                   `json:"context"`
	Value   GetAccountInfoResultValue `json:"value"`
}

// GetAccountInfoResultValue is rpc result of `getAccountInfo`
type GetAccountInfoResultValue struct {
	Lamports  uint64      `json:"lamports"`
	Owner     string      `json:"owner"`
	Excutable bool        `json:"excutable"`
	RentEpoch uint64      `json:"rentEpoch"`
	Data      interface{} `json:"data"`
}

// GetAccountInfo returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetAccountInfo(ctx context.Context, base58Addr string) (GetAccountInfoResponse, error) {
	body, err := c.Call(ctx, "getAccountInfo", base58Addr)
	if err != nil {
		return GetAccountInfoResponse{}, fmt.Errorf("rpc: call error, err: %v", err)
	}

	var res GetAccountInfoResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return GetAccountInfoResponse{}, fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}
	return res, nil
}

// GetAccountInfo returns all information associated with the account of provided Pubkey
func (c *RpcClient) GetAccountInfoWithCfg(ctx context.Context, base58Addr string, cfg GetAccountInfoConfig) (GetAccountInfoResponse, error) {
	body, err := c.Call(ctx, "getAccountInfo", base58Addr, cfg)
	if err != nil {
		return GetAccountInfoResponse{}, fmt.Errorf("rpc: call error, err: %v", err)
	}

	var res GetAccountInfoResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return GetAccountInfoResponse{}, fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}
	return res, nil
}
