package validate

import (
	"crypto/sha256"
)

func MerkleRootHash(b [][]byte) [32]byte {
	tr := b //transaction
	mask := 0x1
	for mask < len(tr) {
		mask = (mask << 1) //bit shift
	}

	rem := make([][]byte, mask-len(tr))

	for i := 0; i < len(rem); i++ {
		rem[i] = []byte{} //rem 은 처음에 null이므로 null byte 로 초기화
	}

	if len(rem) > 0 {
		tr = append(tr, rem...)
	}

	for i := 0; i < len(tr); i++ {
		tr[i] = hash(tr[i])
	}

	for len(tr) > 1 {
		ts := make([][]byte, 0)
		for i := 0; i < len(tr); i += 2 {
			ts = append(ts, hash(tr[i], tr[i+1]))
		}
		tr = ts
	}

	var res [32]byte
	copy(res[:], tr[0][:32])

	return res //루트 해시
}

func hash(b ...[]byte) []byte {
	i := make([]byte, 0)
	for _, v := range b {
		i = append(i, v...)
	}
	res := sha256.Sum256(i) //암호화
	return res[:]
}
