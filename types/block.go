package types
import (
	"encoding/json"
	"crypto/sha256"
	"simplecoin/crypto"
)
type Block struct {
	Header BlockHeader
	Hash Hash //header hash
	Transactions []Transaction
}

type BlockHeader struct {
	PrevBlock Hash
	TransactionsRoot Hash //Merkle hash of all transactions in this block
	StateRoot Hash//Merkle hash of all transactions in this block
	Nonce uint32
}

func (blockHeader BlockHeader) generateHash() Hash{
	data, err := blockHeader.Serialize()
	if err != nil {
		panic(err)
	}
	return sha256.Sum256(data)
}

func (blockHeader BlockHeader) Serialize() ([]byte, error){
	return json.Marshal(blockHeader)
}

func (block Block) IsValid() bool{
	if block.Header.generateHash() != block.Hash {
		return false
	}
	if GenerateMerkleRoot(block.Transactions) != block.Header.TransactionsRoot {
		return false
	}
	return true
}

func GenerateMerkleRoot(transactions []Transaction) Hash {
	hashes := make([][32]byte, len(transactions))
	for i := 0; i < len(transactions); i++{
		hashes[i] = transactions[i].Hash
	}
	return crypto.GetMerkleRoot(hashes)
}
func (block Block) Serialize() ([]byte, error){
	return json.Marshal(block)
}

func (block *Block) Deserialize(data []byte) error {
	return json.Unmarshal(data, block)
}


func GetGenesis() Block {
	block := Block{Header: BlockHeader{}}

	block.Hash = block.Header.generateHash()
	return block
}

