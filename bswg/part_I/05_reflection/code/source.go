package code

import (
	"fmt"
	"reflect"
	"strings"
)

func Demo() {
	fmt.Println("Hello code.Demo() world! (ch5)")
	//
	reflectIntro()
	reflectValue()
	reflectTag()
	lawsOfReflection()
	//
	fmt.Println()
}

type T struct {
	N int
	S string
}

type Adder interface {
	Add(int, int) int
}

type Calculator struct{}

func (c *Calculator) Add(n int, m int) int {
	return n + m
}

type S struct {
	R map[string]int
	S string
	T
}

func printerReflect(offset int, typeOfX reflect.Type) {
	indent := strings.Repeat(" ", offset)
	fmt.Printf("%s %s: %s {\n", indent, typeOfX.Name(), typeOfX.Kind())
	if typeOfX.Kind() == reflect.Struct {
		for i := 0; i < typeOfX.NumField(); i++ {
			innerIndent := strings.Repeat(" ", offset+4)
			f := typeOfX.Field(i)
			if f.Type.Kind() == reflect.Struct {
				printerReflect(offset+4, f.Type)
			} else {
				fmt.Printf("%s %s: %s\n", innerIndent, f.Name, f.Type)
			}
		}
	}
	fmt.Printf("%s }\n", indent)
}

func reflectIntro() {
	fmt.Println()
	//
	var unknown interface{}
	n := 42
	s := "forty two"
	//
	unknown = n
	typeOfU := reflect.TypeOf(unknown)
	fmt.Printf(" '%v', type: %s (%T) %s \n", unknown, typeOfU, unknown, reflect.ValueOf(unknown).Kind())
	//
	unknown = s
	typeOfU = reflect.TypeOf(unknown)
	fmt.Printf(" '%v', type: %s (%T) %s \n", unknown, typeOfU, unknown, reflect.ValueOf(unknown).Kind())
	//
	t := T{n, s}
	typeT := reflect.TypeOf(t)
	fmt.Printf(" typeOf var 't': %s \n", typeT)
	for i := 0; i < typeT.NumField(); i += 1 {
		field := typeT.Field(i)
		fmt.Printf(" - [%d] name: %s, type: %s \n", i, field.Name, field.Type)
	}
	//
	fmt.Println(" =-=")
	//
	var addPtr *Adder
	addPtrType := reflect.TypeOf(addPtr).Elem()
	c := Calculator{}
	calcType := reflect.TypeOf(c)
	calcTypePtr := reflect.TypeOf(&c)
	fmt.Printf(" addPtrType: %s \n", addPtrType)
	fmt.Printf(" calcType: %s, calcTypePtr: %s \n", calcType, calcTypePtr)
	fmt.Printf(" calcType Implements addPtrType: %t, calcTypePtr Implements addPtrType: %t \n", calcType.Implements(addPtrType), calcTypePtr.Implements(addPtrType))
	//
	fmt.Println(" =-=")
	//
	x := S{
		make(map[string]int),
		s,
		T{n, s},
	}
	printerReflect(4, reflect.TypeOf(x))
	//
}

func ValuePrint(i interface{}) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int:
		fmt.Printf(" int with value: %d \n", v.Interface())
	case reflect.String:
		fmt.Printf(" string, content: '%s' \n", v.Interface())
	default:
		fmt.Printf(" %s, value: %v \n", v.Kind(), v.Interface())
	}
}

type Xyz struct {
	X int64
	Y string
	Z float32
}

type Zyx struct {
	Z string
	Y int
	X string
	q string
}

func reflectValue() {
	fmt.Println()
	//
	a := 42
	b := "forty two"
	valueOfA := reflect.ValueOf(a)
	valueOfB := reflect.ValueOf(b)
	fmt.Printf(" value-of a: %v \n", valueOfA.Interface())
	fmt.Printf(" value-of b: %v \n", valueOfB.Interface())
	//
	ValuePrint(a)
	ValuePrint(b)
	ValuePrint(1.2)
	//
	c := Xyz{42, b, 2.1}
	valueOfC := reflect.ValueOf(c)
	fmt.Printf(" valueOf var 'c': %v, kind: %s \n", valueOfC, valueOfC.Kind())
	for i := 0; i < valueOfC.NumField(); i += 1 {
		field := valueOfC.Field(i)
		fmt.Printf(" - [%d] Kind(): %s, String(): %s, value/Interface(): %v \n", i, field.Kind(), field.String(), field.Interface())
	}
	//
	d := Zyx{"HOWDY", 42, "GoodBye", "Oops?"}
	elementD := reflect.ValueOf(&d).Elem()
	fmt.Printf(" element var 'd': %+v, kind: %s \n", elementD, elementD.Kind())
	for i := 0; i < elementD.NumField(); i += 1 {
		field := elementD.Field(i)
		if field.Kind() == reflect.String {
			current := field.String()
			if field.CanSet() {
				field.SetString(strings.ToLower(current))
			} else {
				fmt.Printf(" - can NOT update '%s' \n", current)
			}
		} else {
			fmt.Printf(" - skip %s (other Kind)\n", field.Kind())
		}
	}
	fmt.Printf(" updated var 'd': %+v \n", d)
	//
}

type User struct {
	UserId   string `tagA:"valueA1" tagB:"valueA2"`
	Email    string `tagB:"value"`
	Password string `tagC:"v1 v2"`
}

func reflectTag() {
	fmt.Println()
	//
	userType := reflect.TypeOf(User{})
	fmt.Printf(" user-type: %v \n", userType)
	//
	fieldUserId := userType.Field(0)
	t := fieldUserId.Tag
	fmt.Println(" StructTag is:", t)
	v, _ := t.Lookup("tagA")
	fmt.Printf(" userId tagA: %s\n", v)
	v, _ = t.Lookup("tagB")
	fmt.Printf(" userId tagB: %s\n", v)

	fieldEmail, _ := userType.FieldByName("Email")
	vEmail := fieldEmail.Tag.Get("tagB")
	fmt.Println(" email tagB:", vEmail)

	fieldPassword, _ := userType.FieldByName("Password")
	fmt.Printf(" Password tags: [%s]\n", fieldPassword.Tag)
	fmt.Println("  ", fieldPassword.Tag.Get("tagC"))
	//
	u1 := User{"ABC DEF", "pete@example.org", "Secret42!"}
	fmt.Printf(" u1: %v \n", u1)
	// tags?
}

func lawsOfReflection() {
	fmt.Println()
	//
	fmt.Println("=-= Three laws of Reflection =-=")
	//
	fmt.Println(" - 1st; Reflection goes from interface value to reflection object.")
	var a int32 = 42
	fmt.Printf(" var a init32 = 42, reflect.TypeOf(a): %s \n", reflect.TypeOf(a))
	//
	fmt.Println(" - 2nd; Reflection goes from reflection object to interface value.")
	v := reflect.ValueOf(a)
	fmt.Println(" var a init32 = 42, v := reflect.ValueOf(a), v: ", v)
	fmt.Printf(" var a init32 = 42, v := reflect.ValueOf(a), v.Interface(): %%d==%d \n", v.Interface())
	fmt.Printf(" v.CanSet(): %t   (boolean) \n", v.CanSet())
	//
	fmt.Println(" - 3rd; To modify a reflection object, the value must be settable.")
	//
	//v2 := reflect.ValueOf(a) // note, not an address, just a value
	//v2.SetInt(13)            // <-- panic!
	//
	//v3 := reflect.ValueOf(&a) // note, an address! (not a value)
	//v3.SetInt(23)             // <-- also panic
	//
	v4 := reflect.ValueOf(&a).Elem() // note, an address! (not a value)
	v4.SetInt(24)                    // <-- also panic
	fmt.Printf(" var a init32 = 42, v4 := reflect.ValueOf(&a).Elem(), v4.SetInt(24), v4.Interface(): %%d==%d \n", v4.Interface())
	fmt.Printf(" v4.CanSet(): %t   (boolean) \n", v4.CanSet())
	//
}
