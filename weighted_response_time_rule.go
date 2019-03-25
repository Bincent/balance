package balance

import (
	"errors"
	"github.com/bincent/balance"
	"math/rand"
)

func init() {
	manager := balance.Manager{}
	manager.Register("WeightedResponseTimeRule", &WeightedResponseTimeRule{})
}

type WeightedResponseTimeRule struct {

}

func (this *WeightedResponseTimeRule) Execute(insts [] *balance.Instance,key...string) (inst *balance.Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)

	index := rand.Intn(lens)
	inst = insts[index]

	return
}
