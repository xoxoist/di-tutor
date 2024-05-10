package head

import (
	"go.uber.org/dig"
)

type DependenciesHolder struct {
	dig.In
	Ears *Ears
	Eyes *Eyes
}

func RegisterDependencies(container *dig.Container) error {
	var err error
	err = container.Provide(NewEars)
	err = container.Provide(NewEyes)
	if err != nil {
		return err
	}
	return nil
}
