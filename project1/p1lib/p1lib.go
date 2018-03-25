package p1lib

import "github.com/wix-private/BazelBuildForGo/project1/dep1"

func DoSomeWorksP1()(retVal string){
	retVal = dep1.DoSomeDep1()
	retVal = retVal + " DoSomeWorksP1 called"
	return retVal
}