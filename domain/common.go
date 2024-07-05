package domain

type State = int

const (
	_ State = iota
	Renew
	Redeem
	Normal
)
