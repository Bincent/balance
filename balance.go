package balance

type Balance interface {
	/**
	 * 负载均衡算法
	**/
	Execute([] *Instance,...string) (*Instance,error)
}