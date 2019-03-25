package balance

import "fmt"

type Manager struct {
	allBalance map[string]Balance
}

var manager = Manager{
	allBalance: make(map[string]Balance),
}

func Random(insts []*Instance) (inst *Instance, err error)  {
	return manager.execute("random", insts)
}

func  Round(insts []*Instance) (inst *Instance, err error)  {
	return manager.execute("round", insts)
}

func (this *Manager) register(name string, balance Balance) {
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