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
	fmt.Printf(" element var 'd': %v, kind: %s \n", elementD, elementD.Kind())
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
	fmt.Printf(" updated var 'd': %v \n", d)
	//
}
