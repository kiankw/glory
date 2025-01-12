package mq

import (
	"fmt"
)

func GetMQInstance(name string) MQService {
	if !inited {
		panic("mq not init, please call mq.Init() before get instance")
	}
	srv, ok := mqInstance[name]
	if !ok {
		panic(fmt.Sprintf("mq with name %s not found", name))
	}
	err := srv.Connect()
	if err != nil {
		panic(fmt.Sprintf("fail to connect to mq with name %s, err is %v", name, err))
	}
	return srv
}
