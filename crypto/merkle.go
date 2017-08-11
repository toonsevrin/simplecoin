package crypto

import (
	"crypto/sha256"
)
func GetMerkleRoot(nodes [][32]byte) [32]byte{
	if len(nodes) == 1 {
		return nodes[0]
	}
	newNodes := make([][32]byte, len(nodes)/2)
	for i := 0; 2*i + 1 < len(nodes); i ++ {
		result := append(nodes[2*i][:], nodes[2*i+ 1][:]...)
		newNodes[i] = sha256.Sum256(result)
	}
	if len(nodes) % 2 == 1 {
		newNodes = append(newNodes, nodes[len(nodes) - 1])
	}
	return GetMerkleRoot(newNodes)
}
