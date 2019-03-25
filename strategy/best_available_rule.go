package strategy

import (
	"errors"
	"math/rand"
	"reepu.com/balance"
)

func init() {
	balance.manager.register("BestAvailableRule", &BestAvailableRule{})
}

type BestAvailableRule struct {

}

func (this *BestAvailableRule) Execute(insts [] *Instance,key...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)

	index := rand.Intn(lens)
	inst = insts[index]

	return
}