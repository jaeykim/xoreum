package main

import "fmt"

type Transaction struct {
	data txdata
	hash atomic.Value
	size atomic.Value
	from atmoc.Value
}

type txdata struct {
	SenderNonce		uint64
	RecipientNonce	uint64
	Sender			*common.Address
	Recipient		*common.Address
	Payload			[]byte
	
	// Signature values

	// This is only used when marshaling to JSON
	Hash *common.Hash
}
