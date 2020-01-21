package balance

import (
	"fmt"
	"hash/crc32"
	"math/rand"
)

func init() {
	RegisterBalancer("hash", &HashBalance{})
}

type HashBalance struct {
}

func (p *HashBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	defKey := fmt.Sprintf("%d", rand.Int())

	if len(key) > 0 {
		defKey = key[0]
		return
	}
	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}
	hashVal := crc32.ChecksumIEEE([]byte(defKey))
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
