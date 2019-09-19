package main

import (
	"fmt"
	"reflect"
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

func main() {
	avengerdetails := newAvenger("Ironman", "Tony Stark")

	fmt.Println(reflect.TypeOf(avengerdetails))
	fmt.Println(avengerdetails.SuperHeroName)
}
