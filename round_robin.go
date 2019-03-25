package balance

import (
	"errors"
	"github.com/bincent/balance"
)

func init() {
	manager := balance.Manager{}
	manager.Register("RoundRobin", &RoundRobin{})
}

// 轮询调度算法
type RoundRobin struct {
	curIndex int
}

func (this *RoundRobin) Execute(insts [] *balance.Instance, key ...string) (inst *balance.Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)
	if this.curIndex >= lens {
		this.curIndex = 0
	}
	inst = insts[this.curIndex]
	this.curIndex++
	return

}