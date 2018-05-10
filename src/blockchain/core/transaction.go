package core

import "fmt"

type Address [20]byte

type Transaction struct { //블록 내용 (거래기록 등)
	From Address
	Data txdata
}

func (a Address) ToString() string {
	res := "0x"
	for _, v := range a {
		res += fmt.Sprintf("%02x", v)
	}
	return res
}

type txdata struct {
	To     Address
	Amount uint64
}
