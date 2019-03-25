package balance

import "fmt"

type Manager struct {
	allBalance map[string]Balance
}

var manager = Manager{
	allBalance: make(map[string]Balance),
}

func Random(insts []*Instance) (inst *Instance, err error)  {
	return manager.execute("RandomRule", insts)
}

func  Round(insts []*Instance) (inst *Instance, err error)  {
	return manager.execute("RoundRobin", insts)
}

func Weighted(insts []*Instance) (inst *Instance, err error)  {
	return manager.execute("WeightedResponseTimeRule", insts)
}

func BestAvailable(insts []*Instance) (inst *Instance, err error)  {
	return manager.execute("BestAvailableRule", insts)
}

func (this *Manager) Register(name string, balance Balance) {
	this.allBalance[name] = balance
}

func (this *Manager) execute(name string, insts []*Instance) (inst *Instance, err error) {
	balance, ok := manager.allBalance[name]
	if !ok {
		err = fmt.Errorf("not fount %s", name)
		fmt.Println("not found ",name)
		return
	}

	inst, err = balance.Execute(insts)
	if err != nil {
		err = fmt.Errorf(" %s erros", name)
		return
	}

	return
}