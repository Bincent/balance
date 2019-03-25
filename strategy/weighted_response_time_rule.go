package strategy

import (
	"errors"
	"github.com/bincent/balance"
	"math/rand"
)

func init() {
	balance.manager.register("WeightedResponseTimeRule", &WeightedResponseTimeRule{})
}

type WeightedResponseTimeRule struct {

}

func (this *WeightedResponseTimeRule) Execute(insts [] *Instance,key...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)

	index := rand.Intn(lens)
	inst = insts[index]

	return
}
