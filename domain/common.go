package domain

type State = string

const (
	Renew  State = "1" // 续费
	Redeem State = "2" // 赎回
	Normal State = "3" // 正常
)
