package main

import (
	p1 "github.com/wix-private/BazelBuildForGo/project1/p1lib"
)

func main(){
	println ("Project2 called")
	res := p1.DoSomeWorksP1()
	println(res)
}