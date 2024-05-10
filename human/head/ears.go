package head

import "fmt"

type Ears struct{}

func (e *Ears) Hearing() {
	fmt.Println("Hearing...")
}

func NewEars() *Ears {
	return &Ears{}
}
