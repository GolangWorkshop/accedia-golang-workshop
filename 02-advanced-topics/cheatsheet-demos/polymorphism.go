// TODO: rework example, shapes and surface area might be a better example

package main

import "fmt"

type Country interface {
	getCapital() string
}

// a structure that has the required by the interface field
type Bulgaria struct {
	Capital string
}

// implementing methods of interface 'Country'
// for structure 'Bulgaria'
func (bulgaria Bulgaria) getCapital() string {
	return bulgaria.Capital
}

// a second structure that has the required by the interface field
type Germany struct {
	Capital string
}

// implementing methods of interface 'Country'
// for structure 'Germany'
func (germany Germany) getCapital() string {
	return germany.Capital
}

func main() {

	// creating an instance of interface 'Country'
	var ICountry Country

	bulgaria := Bulgaria{Capital: "Sofia"}
	germany := Germany{Capital: "Berlin"}

	// assigning object 'bulgaria' to 'ICountry'
	// and invoking getCapital()
	ICountry = bulgaria
	fmt.Println(ICountry.getCapital()) // this prints Sofia in the console

	// assigning object 'germany' to 'ICountry'
	// and invoking getDevelopedBy()
	ICountry = germany
	fmt.Println(ICountry.getCapital()) // this prints Berlin in the console
}
