package main

import (
	"fmt"
	"github.com/bincent/balance"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	var insts []*balance.Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	for {
		inst, err := balance.Random(insts)
		if err != nil {
			fmt.Println("do balance err")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}