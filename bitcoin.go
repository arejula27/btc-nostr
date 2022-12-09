package main

import (
	"encoding/json"
	"log"
	"time"
)

type BitcoinConf struct {
	NodeUrl  string
	Port     int
	RpcUser  string
	Password string
	SSLmode  bool
}

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

const TIMEOUT = time.Second * 5

func NewBitcoinCli(conf BitcoinConf) *bitcoinCli {

	client, err := newRpcClient(conf.NodeUrl, conf.Port, conf.RpcUser, conf.Password, conf.SSLmode, TIMEOUT)

	if err != nil {
		log.Fatal(err)

	}
	return &bitcoinCli{
		client: client,
	}
}

func (b *bitcoinCli) getBestBlockhash() (bestBlockHash string, err error) {
	r, err := b.client.call("getbestblockhash", nil)
	if err = handleRpcError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &bestBlockHash)
	return
}

func (b *bitcoinCli) getBlockheader(blockHash string) (*BlockHeader, error) {
	r, err := b.client.call("getblockheader", []string{blockHash})
	if err = handleRpcError(err, &r); err != nil {
		return nil, err
	}

	var blockHeader BlockHeader
	err = json.Unmarshal(r.Result, &blockHeader)

	return &blockHeader, err
}
