package main

import (

	"github.com/wix-private/BazelBuildForGo/project1/p1lib"
	"github.com/coreos/etcd/client"
)

type EtcdClientAPI struct {
	client  client.KeysAPI
}

func main(){
	_ = &EtcdClientAPI{
		client : nil,
	}
	println ("Project1 called")
	res := p1lib.DoSomeWorksP1()
	println (res)
}