package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Avenger struct {
	ID            int
	SuperHeroName string
	RealName      string
}

var ID int = 0

func newAvenger(HeroName string, NormalName string) *Avenger {
	ID++ // note it by default this value starts at zero.
	return &Avenger{
		ID:            ID,
		SuperHeroName: HeroName,
		RealName:      NormalName,
	}

}

/*
Here's a method we can apply once our constructor has instantiated our Avengers custom type
*/
func (MyHero Avenger) UpperCase() string {

	// Notice how we drill down the custom object's (struct) field using the dot notaton:
	return strings.ToUpper(MyHero.SuperHeroName)
}

func main() {

	// First we create our custom type using the newAvenger
	avengerdetails := newAvenger("Ironman", "Tony Stark")

	// Let's confirm that has worked.
	fmt.Println(reflect.TypeOf(avengerdetails))
	fmt.Println(avengerdetails.SuperHeroName)

	// Now let's apply an method to our object.
	fmt.Println(avengerdetails.UpperCase())

}
