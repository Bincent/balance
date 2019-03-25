package balance

import (
	"errors"
	"math/rand"
)

func init()  {
	manager.register("random", &RandomRule{})
}

// 随机策略
type RandomRule struct {
}

func (this *RandomRule) Execute(insts [] *Instance,key...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)

	index := rand.Intn(lens)
	inst = insts[index]

	return
}