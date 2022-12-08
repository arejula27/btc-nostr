package main

import (
	"encoding/json"
)

type BlockHeader struct {
	Hash              string
	Confirmations     int
	Height            int
	Version           uint32
	VersionHex        string
	Merkleroot        string
	Time              int64
	Mediantime        int64
	Nonce             uint32
	Bits              string
	Difficulty        float64
	Chainwork         string
	Txes              int    `json:"nTx"`
	Previousblockhash string `json:"omitempty"`
	Nextblockhash     string `json:"omitempty"`
}

type bitcoinCli struct {
	client    *rpcClient
	lastBlock string
}

func NewBitcoinCli(client *rpcClient) *bitcoinCli {
	return &bitcoinCli{
		client: client,
	}
}

func (b *bitcoinCli) getBestBlockhash() (bestBlockHash string, err error) {
	r, err := b.client.call("getbestblockhash", nil)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &bestBlockHash)
	return
}

func (b *bitcoinCli) getBlockheader(blockHash string) (*BlockHeader, error) {
	r, err := b.client.call("getblockheader", []string{blockHash})
	if err = handleError(err, &r); err != nil {
		return nil, err
	}

	var blockHeader BlockHeader
	err = json.Unmarshal(r.Result, &blockHeader)

	return &blockHeader, err
}
