package types



type Account struct {
	PublicAccount
	Hash Hash
}

type PublicAccount struct {
	Address Hash//A hash of the pubkey
	Balance uint32
}