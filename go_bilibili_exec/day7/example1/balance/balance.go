package balance

type IBalancer interface {
	DoBalance([]*Instance) (*Instance, error)
}
