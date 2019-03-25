package balance

import (
	"errors"
	"net"
	"strconv"
)

func init() {
	manager.register("RoundRobin", &RoundRobin{})
}

// 轮询调度算法
type RoundRobin struct {
	index int
}

func (this *RoundRobin) Execute(insts [] *Instance, key ...string) (address string, err error) {
	lens := len(insts)

	if lens == 0 {
		return "", errors.New("no instance")
	}

	if this.index >= lens { this.index = 0 }

	inst := insts[this.index]
	this.index++

	return net.JoinHostPort(inst.host, strconv.Itoa(inst.port)), nil
}