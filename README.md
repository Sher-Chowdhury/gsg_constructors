# gsg_constructors


A constructor is just an ordinary function, but it is a function that's specifically designed to create a struct object. 

E.g. if successive struct objects are to have incrementing id value, then constructor function will house the logic to handle that for you. Here's how to create it, first we create a custom variable:


```go
type Avenger struct {
	ID            int
	SuperHeroName string
	RealName      string
}
```

now we need to ensure that all Avenger objects variables have valid content, and to do that we create a constructor function that creates all Avenger variables for us. By convention the constructor functio is named after the struct object, but with 'new' preceding it:


```go
func newAvenger(HeroName string, NormalName string) *Avenger {
  AvengerDetails := Avenger{
		ID:            ID,
		SuperHeroName: HeroName,
		RealName:      NormalName,
	}
  ID++   // note it by default this value starts at zero. 

  // This returns the object' memory location, only because there's no need to return anything. 
  return &AvengerDetails
}
```



