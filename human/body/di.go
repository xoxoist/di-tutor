package body

import (
	"go.uber.org/dig"
)

type DependenciesHolder struct {
	dig.In
	Hands *Hands
	Legs  *Legs
}

func RegisterDependencies(container *dig.Container) error {
	var err error
	err = container.Provide(NewHands)
	err = container.Provide(NewLegs)
	if err != nil {
		return err
	}
	return nil
}
