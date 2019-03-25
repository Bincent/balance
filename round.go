package balance

import "errors"

func init() {
	manager.register("round", &RoundRobin{})
}

// 轮询调度算法
type RoundRobin struct {
	curIndex int
}

func (this *RoundRobin) Execute(insts [] *Instance, key ...string) (inst *Instance, err error) {
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