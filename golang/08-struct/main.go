package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main2() {
	aDoctor := struct{ name string }{name: "john"}
	pl(aDoctor)
}

type Animal struct {
	Name string
}

type Bird struct {
	Animal
	SpeedKPH float32
}

func main3() {
	bird1 := Bird{}
	bird1.Name = "b1"
	bird1.SpeedKPH = 1000.1
	pl(bird1)
	bird2 := Bird{
		Animal:   Animal{Name: "b2"},
		SpeedKPH: 2100,
	}
	pl(bird2)
}

type Person struct {
	Name string `todo:"test"`
}

func (pn *Person) SayMyName() string {
	return pn.Name
}

func main() {
	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("Name")
	pl(field.Tag)

	jack := Person{Name: "Jack"}

	pl(jack.SayMyName())
}
