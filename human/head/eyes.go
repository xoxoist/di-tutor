package head

import "fmt"

type Eyes struct{}

func (e *Eyes) Seeing() {
	fmt.Println("Seeing...")
}

func NewEyes() *Eyes {
	return &Eyes{}
}
