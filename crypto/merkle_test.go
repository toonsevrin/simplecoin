package crypto

import (
	"testing"
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"bytes"
)
func TestGetMerkleRootTwoNodes(t *testing.T) {
	one := sha256.Sum256([]byte("test"))
	two := sha256.Sum256([]byte("ok"))
	expected := sha256.Sum256( append(one[:], two[:]...))

	actual := GetMerkleRoot([][32]byte{sha256.Sum256([]byte("test")), sha256.Sum256([]byte("ok"))})

	assert.True(t, bytes.Equal(expected[:], actual[:]))
}

func TestGetMerkleRootTwoNodesFalse(t *testing.T) {
	one := sha256.Sum256([]byte("test"))
	two := sha256.Sum256([]byte("ok"))
	expected := sha256.Sum256( append(one[:], two[:]...))

	actual := GetMerkleRoot([][32]byte{sha256.Sum256([]byte("test")), sha256.Sum256([]byte("ok?"))})

	assert.False(t, bytes.Equal(expected[:], actual[:]))
}
func TestGetMerkleRootThreeNodes(t *testing.T) {
	one := sha256.Sum256([]byte("test"))
	two := sha256.Sum256([]byte("ok"))
	oneTwo := sha256.Sum256( append(one[:], two[:]...))
	three := sha256.Sum256([]byte("hm"))
	expected := sha256.Sum256( append(oneTwo[:], three[:]...))
	actual := GetMerkleRoot([][32]byte{sha256.Sum256([]byte("test")), sha256.Sum256([]byte("ok")), sha256.Sum256([]byte("hm"))})

	assert.True(t, bytes.Equal(expected[:], actual[:]))
}
func TestGetMerkleRootFourNodes(t *testing.T) {
	one := sha256.Sum256([]byte("test"))
	two := sha256.Sum256([]byte("ok"))
	oneTwo := sha256.Sum256( append(one[:], two[:]...))
	three := sha256.Sum256([]byte("hm"))
	four := sha256.Sum256([]byte("four"))
	threeFour := sha256.Sum256( append(three[:], four[:]...))
	expected := sha256.Sum256( append(oneTwo[:], threeFour[:]...))
	actual := GetMerkleRoot([][32]byte{sha256.Sum256([]byte("test")), sha256.Sum256([]byte("ok")), sha256.Sum256([]byte("hm")), sha256.Sum256([]byte("four"))})

	assert.True(t, bytes.Equal(expected[:], actual[:]))
}