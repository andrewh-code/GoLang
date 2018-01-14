package model

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// populate the Person struct
func (p *Person) PopulatePerson() (personOut Person) {
	p.Firstname = "Andrew"
	p.Lastname = "Code"

	return
}
