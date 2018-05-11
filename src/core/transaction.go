package core

import "fmt"

type Address [20]byte

type Transaction struct {
	From Address
	Data txdata
}

func (a Address) ToString() string {
	res := "0x"
	for _, v := range a {
		res += fmt.Sprintf("%02x", v)
	} //트랜잭션 출력

	return res
}

type txdata struct {
	To     Address
	Amount uint64
}

func NewTransaction(amount uint64, from Address, to Address) *Transaction {
	d := txdata{
		To:     to,
		Amount: amount,
	}
	return &Transaction{From: from, Data: d}
}

func (t *Transaction) ToBytes() []byte {
	res := t.From[:]
	res = append(res, t.Data.To[:]...)
	buf := make([]byte, 8)
	mask := uint64(0xff)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte((t.Data.Amount >> uint(56-(8*i))) & mask)
	}
	res = append(res, buf...)
	return res
}
