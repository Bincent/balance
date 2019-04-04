# Balance

#### 介绍
安装
```shell
go get -u -v github.com/bincent/balance
```

使用示例

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
```