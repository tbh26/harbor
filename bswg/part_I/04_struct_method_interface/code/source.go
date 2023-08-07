package code

import (
	"errors"
	"fmt"
	"reflect"
)

func Demo() {
	fmt.Println("Hello code.Demo() world! (ch4)")
	//
	structDemo()
	methodDemo()
	interfaceDemo()
	//
	fmt.Println()
}

type Item struct {
	Topic       string
	Description string
	Code        int64
}

func NewItem(topic string, desciption string, code int64) (*Item, error) {
	if len(topic) == 0 || code < 0 {
		return nil, errors.New("invalid input(s)/param(s)")
	}
	return &Item{topic, desciption, code}, nil
}

type DayOfBirth struct {
	Day   int8
	Month int8
	Year  int
}

type Person struct {
	FullName string
	Dob      DayOfBirth // nested struct
}

type Person2 struct {
	FullName   string
	DayOfBirth // embedded struct (as long as uniq)
}

func structDemo() {
	fmt.Println()
	//
	i1 := Item{}
	fmt.Printf(" i1 == %+v \n", i1)
	//
	i2 := Item{"chair", "", 12345}
	i2.Code = 54321
	fmt.Printf(" i2 == %v \n", i2)
	//
	i3 := Item{Topic: "table", Description: "workdesk", Code: 98765}
	fmt.Printf(" i3 == %+v \n", i3)
	//
	i4 := Item{Topic: "Closet"}
	fmt.Printf(" i4 == %v \n", i4)
	//
	//	i5, e := NewItem("", "", -12) // panics!
	i5, e := NewItem("stool", "n", 12)
	if e != nil {
		panic(e)
	}
	fmt.Printf(" *i5 == %+v \n", *i5)
	//
	//anonymous struct
	as := struct {
		label string
		val   int
	}{"thing", 42}
	fmt.Printf(" as == %+v \n", as)
	//
	fmt.Printf(" var 'i4' type-of '%s', var 'as' type-of '%s' \n", reflect.TypeOf(i4), reflect.TypeOf(as))
	//
	dob := DayOfBirth{21, 3, 1987}
	name := "John Doe"
	jd := Person{name, dob}
	fmt.Printf(" jd == %+v \n", jd)
	p2 := Person2{name, dob}
	fmt.Printf(" p2 == %+v \n", p2)
}

func (i Item) info() string { //Item receiver
	return fmt.Sprintf("Item{topic: '%s', description: '%s', code: '%d'}", i.Topic, i.Description, i.Code)
}

func (i *Item) reverseSome() {
	i.Topic = Reverse(i.Topic)
	i.Description = Reverse(i.Description)
}

func Reverse(s string) string {
	runes := []rune(s)
	for index, back := 0, len(runes)-1; index < back; index, back = index+1, back-1 {
		runes[index], runes[back] = runes[back], runes[index]
	}
	return string(runes)
}

func (dob DayOfBirth) ymd() string { // maybe an embedded method?
	return fmt.Sprintf("%04d-%02d-%02d", dob.Year, dob.Month, dob.Day)
}

func (p Person2) info() string {
	//	return fmt.Sprintf("Person2{full name: '%s', date of birth: '%s'}", p.FullName, p.DayOfBirth.ymd())
	return fmt.Sprintf("Person2{full name: '%s', date of birth: '%s'}", p.FullName, p.ymd()) //<-embedded-method
}

func methodDemo() {
	fmt.Println()
	//
	i1, _ := NewItem("stool", "🪑 n", 357)
	var i Item = *i1
	fmt.Println(" item:", i.info())
	//
	//	(&i).reverseSome()
	i.reverseSome()
	fmt.Println(" item:", i.info())
	//
	dob := DayOfBirth{21, 3, 1987}
	name := "John Doe"
	p := Person2{name, dob}
	fmt.Println(" person:", p.info())
}

type Alive interface {
	Greet() string
	Move(string) string
}

func (p Person2) Greet() string {
	return fmt.Sprintf("Hello '%s'!", p.FullName)
}

func (p Person2) Move(direction string) string {
	return fmt.Sprintf("'%s' moved '%s'", p.FullName, direction)
}

type Animal struct {
	Tag string
}

func (a *Animal) Greet() string {
	return fmt.Sprintf("Hi '%s'.", a.Tag)
}

func (a *Animal) Move(direction string) string {
	return fmt.Sprintf("'%s' moves '%s'", a.Tag, direction)
}

func demonstrate(a Alive, d string) { //interface as param
	fmt.Println(a.Greet())
	fmt.Println(a.Move(d))
}

func interfaceDemo() {
	fmt.Println()
	//
	p := Person2{"John Doe", DayOfBirth{21, 3, 1987}}
	direction := "north"
	fmt.Println(p.Greet())
	fmt.Println(p.Move(direction))
	demonstrate(p, direction)
	demonstrate(&p, direction)
	//
	cat := &Animal{Tag: "Kitty"}
	direction = "away"
	fmt.Println((*cat).Greet())
	fmt.Println(cat.Move(direction))
	//	demonstrate(*cat, direction)
	demonstrate(cat, direction)
	//
}
