package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Ashish",
		Price: 1.00,
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
