package types

import (
	"encoding/json"
	"crypto/sha256"
)
type Transaction struct {
	Header TransactionHeader
	Hash Hash
}

type TransactionHeader struct {
	Sender Hash
	Receiver Hash
	Amount uint32
	Fee uint32
}

func (transaction Transaction) generateHash() Hash {
	data, err := transaction.Serialize()
	if err != nil {
		panic(err)
	}
	return sha256.Sum256(data)
}

func (transaction Transaction) Serialize() ([]byte, error) {
	return json.Marshal(transaction)
}


func (transaction *Transaction) Deserialize(data []byte) error{
	return json.Unmarshal(data, transaction)
}