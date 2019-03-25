package balance

import "fmt"

type Manager struct {
	allBalance map[string]Balance
}

var manager = Manager{
	allBalance: make(map[string]Balance),
}

// 随机策略
func Random(insts []*Instance) (address string, err error)  {
	return manager.execute("RandomRule", insts)
}

// 轮询调度算法
func Round(insts []*Instance) (address string, err error)  {
	return manager.execute("RoundRobin", insts)
}

// 加权策略
func Weighted(insts []*Instance) (address string, err error)  {
	return manager.execute("WeightedResponseTimeRule", insts)
}

// 请求数最少策略
func BestAvailable(insts []*Instance) (address string, err error)  {
	return manager.execute("BestAvailableRule", insts)
}

func (this *Manager) register(name string, balance Balance) {
	this.allBalance[name] = balance
}

func (this *Manager) execute(name string, insts []*Instance) (address string, err error) {
	balance, ok := manager.allBalance[name]
	if !ok {
		return  "", fmt.Errorf("not fount %s", name)
	}

	return balance.Execute(insts)
}