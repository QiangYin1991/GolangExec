package balance

type IBalancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}
