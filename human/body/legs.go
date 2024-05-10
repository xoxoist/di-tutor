package body

import "fmt"

type Legs struct{}

func (l *Legs) WalkWithRightLeg() {
	fmt.Println("Walking with right leg")
}

func (l *Legs) WalkWithLeftLeg() {
	fmt.Println("Walking with left leg")
}

func (l *Legs) KickWithRightLeg() {
	fmt.Println("Kicking with right leg")
}

func (l *Legs) KickWithLeftLeg() {
	fmt.Println("Kicking with left leg")
}

func NewLegs() *Legs {
	return &Legs{}
}
