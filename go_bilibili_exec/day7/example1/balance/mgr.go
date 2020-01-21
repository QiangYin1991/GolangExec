package balance

import "fmt"

var (
	mgr = BalanceMgr{
		allBalancer: make(map[string]IBalancer),
	}
)

func RegisterBalancer(name string, b IBalancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}

	fmt.Println("DoBalance " + name)
	inst, err = balancer.DoBalance(insts)
	return
}

type BalanceMgr struct {
	allBalancer map[string]IBalancer
}

func (p *BalanceMgr) registerBalancer(name string, b IBalancer) {
	p.allBalancer[name] = b
}
