package model

type Person struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

// populate the Person struct
func (p *Person) PopulatePerson() (personOut Person) {
	p.Firstname = "Andrew"
	p.Lastname = "Code"

	return
}
