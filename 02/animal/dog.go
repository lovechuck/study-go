package animal

//Fox ...
type Fox struct {
	Name string
}

//Woo ...
func (fox Fox) Woo() string {
	return fox.Name + " wooooooo"
}
