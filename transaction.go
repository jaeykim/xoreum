package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	data txdata
	hash atomic.Value
	size atomic.Value
	from atmoc.Value
}

type txdata struct {
	Participants		[]*common.Address
	ParticipantsNonce	[]uint64
	XORs				[]uint64
	Payload				[]byte
	
	// Signature values

	// This is only used when marshaling to JSON
	Hash *common.Hash
}
