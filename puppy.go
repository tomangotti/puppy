package puppy

import (
	"github.com/tomangotti/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	s1:= dog.WhenGrownUps(Bark())

	return s1
}