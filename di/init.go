package di

import (
	"github.com/xoxoist/di-tutor/human"
	"github.com/xoxoist/di-tutor/human/body"
	"github.com/xoxoist/di-tutor/human/head"
	"go.uber.org/dig"
)

// Container master containers for all dependencies injected
// this global variable will be accessed from main function
// and will provide needed instances across functionalities
var Container = dig.New()

// Injected this struct represents dependencies injections
// bank whole injected instance will be accessed from this
// structure.
type Injected struct {
	FooBar string
	Head   head.DependenciesHolder
	Body   body.DependenciesHolder
	Human  *human.Human
}

// NewInjected initialize dependencies injection entries
// for all dependencies based what this function params
// needed will be injected again using Injected struct.
func NewInjected(
	hd head.DependenciesHolder,
	bd body.DependenciesHolder,
	hm *human.Human,
) *Injected {
	return &Injected{
		FooBar: "Hello This is FooBar content",
		Head:   hd,
		Body:   bd,
		Human:  hm,
	}
}

// init default initialization function from golang
func init() {
	var err error
	// Injecting needed dependencies across functionalities
	err = head.RegisterDependencies(Container)
	err = body.RegisterDependencies(Container)
	err = Container.Provide(human.NewHuman)

	// Wrapping up all injected dependencies
	err = Container.Provide(NewInjected)
	if err != nil {
		panic(err)
	}
}
