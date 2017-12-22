package main

import "fmt"

type any interface {
}

type evalFunc func(any any) (any, any)

func main() {
	evenFunc := func(state any) (any, any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	even := buildlazyintevaluator(evenFunc, 0)
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth even :%v\n", i, even())
	}
}

func buildlazyevaluator(evalFunc evalFunc, initstate any) func() any {
	retvalchan := make(chan any)
	loopFunc := func() {
		var actstate any = initstate
		var retval any
		for {
			retval, actstate = evalFunc(actstate)
			retvalchan <- retval
		}
	}
	retFunc := func() any {
		return <-retvalchan
	}

	go loopFunc()
	return retFunc
}

func buildlazyintevaluator(evalFunc evalFunc, initstate any) func() int {
	ef := buildlazyevaluator(evalFunc, initstate)
	return func() int {
		return ef().(int)
	}
}
