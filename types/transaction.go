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
	Inputs []Input//amount difference can be claimed as fee
	Outputs []Output

}

type Output struct {
	Address Hash
	Amount uint32
}

type Input struct {
	Address Hash
	Amount uint32
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