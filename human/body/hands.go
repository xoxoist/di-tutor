package body

import "fmt"

type Hands struct{}

func (h *Hands) TakeWithRightHand() {
	fmt.Println("Taking with right hand")
}

func (h *Hands) TakeWithLeftHand() {
	fmt.Println("Taking with left hand")
}

func (h *Hands) PunchWithRightHand() {
	fmt.Println("Punching with right hand")
}

func (h *Hands) PunchWithLeftHand() {
	fmt.Println("Punching with left hand")
}

func NewHands() *Hands {
	return &Hands{}
}
