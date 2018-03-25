package main

import (

	"github.com/wix-private/BazelBuildForGo/project1/p1lib"
)

func main(){
	println ("Project1 called")
	res := p1lib.DoSomeWorksP1()
	println (res)
}