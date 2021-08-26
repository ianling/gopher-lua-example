package dog

// Dog is a Go representation of a dog.
type Dog struct {
	Name  string
	Speak func()
}

// NewDog creates a new Dog object.
func NewDog(name string) Dog {
	return Dog{Name: name}
}
