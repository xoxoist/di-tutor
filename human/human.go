package human

import (
	"github.com/xoxoist/di-tutor/human/body"
	"github.com/xoxoist/di-tutor/human/head"
)

type Human struct {
	HeadDependenciesHolder head.DependenciesHolder
	BodyDependenciesHolder body.DependenciesHolder
}

func (h *Human) RunningAllHumanFunction() {
	// head
	h.HeadDependenciesHolder.Ears.Hearing()
	h.HeadDependenciesHolder.Eyes.Seeing()

	// body
	h.BodyDependenciesHolder.Hands.TakeWithRightHand()
	h.BodyDependenciesHolder.Hands.TakeWithLeftHand()
	h.BodyDependenciesHolder.Hands.PunchWithRightHand()
	h.BodyDependenciesHolder.Hands.PunchWithLeftHand()
	h.BodyDependenciesHolder.Legs.WalkWithRightLeg()
	h.BodyDependenciesHolder.Legs.WalkWithLeftLeg()
	h.BodyDependenciesHolder.Legs.KickWithRightLeg()
	h.BodyDependenciesHolder.Legs.KickWithLeftLeg()
}

func NewHuman(
	headDependenciesHolder head.DependenciesHolder,
	bodyDependenciesHolder body.DependenciesHolder,
) *Human {
	return &Human{
		HeadDependenciesHolder: headDependenciesHolder,
		BodyDependenciesHolder: bodyDependenciesHolder,
	}
}
