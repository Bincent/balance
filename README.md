# Balance

#### 介绍

```go
// 预置数组
var insts []*balance.Instance

host := "192.168.1.1"
port, _ := 8080
one := balance.NewInstance(host, port)
insts = append(insts, one)

// 调用负载均衡说明
// RandomRule表示随机策略
inst, err := balance.Random(insts)

// RoundRobin表示轮询策略
inst, err := balance.Round(insts)

//WeightedResponseTimeRule表示加权策略
inst, err := balance.Weighted(insts)

//BestAvailableRule表示请求数最少策略
inst, err := balance.BestAvailable(insts)

```