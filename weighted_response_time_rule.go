package balance

import (
	"errors"
	"math/rand"
	"net"
	"strconv"
)

func init() {
	manager.register("WeightedResponseTimeRule", &WeightedResponseTimeRule{})
}

type WeightedResponseTimeRule struct {

}

func (this *WeightedResponseTimeRule) Execute(insts [] *Instance,key...string) (address string, err error) {
	lens := len(insts)
	if lens == 0 {
		return "", errors.New("no instance")
	}

	index := rand.Intn(lens)
	inst := insts[index]

	return net.JoinHostPort(inst.host, strconv.Itoa(inst.port)), nil
}
