// example made by matryer on medium

package main

import (
	"fmt"
	s "github.com/oxequa/realize-examples/interface/sub"
)

// Speaker types can say things.
type Speaker interface {
	Say(string)
}

type Person struct {
	name string
}

func (p *Person) Say(message string) {
	fmt.Println(p.name+":", message)
}

func SpeakAlphabet(via Speaker) {
	via.Say("abcdefghijklmnopqrstuvwxyz")
}

func main() {
	first := new(Person)
	first.name = "Anonymous"
	SpeakAlphabet(first)
	second := new(s.SpeakWriter)
	SpeakAlphabet(second)
}
