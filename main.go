package main

import (
	"github.com/xoxoist/di-tutor/di"
)

func main() {
	err := di.Container.Invoke(func(inj *di.Injected) {
		inj.Human.RunningAllHumanFunction()
	})
	if err != nil {
		panic(err)
	}
}
