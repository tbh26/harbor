package code

import (
	"encoding/xml"
	"errors"
	"fmt"
)

func XmlDemo() {
	fmt.Println("Hello code.XmlDemo() world! (ch8)")

	marshalXml()
	nextDemo()
	lastDemo()

	fmt.Println()
}

func marshalXml() {
	fmt.Println("\n=-= marshal xml ... =-=")

	number, err := xml.Marshal(42)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(number))

	float, _ := xml.Marshal(3.14)
	fmt.Println(string(float))

	msg, _ := xml.Marshal("This is a msg!!!")
	fmt.Println(string(msg))

	numbers, _ := xml.Marshal([]int{1, 2, 2, 3, 5, 8})
	fmt.Println(string(numbers))

	aMap, err := xml.Marshal(map[string]int{"one": 1, "two": 2})
	fmt.Println(err)
	fmt.Println("-", string(aMap), "-")

	fmt.Println()
}

type MyMap map[string]string

func (s MyMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}

	for key, value := range s {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return e.Flush()
}

func (a MyMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	key := ""
	val := ""

	for {

		t, _ := d.Token()
		switch tt := t.(type) {

		case xml.StartElement:
			key = tt.Name.Local
		case xml.CharData:
			val = string(tt)
		case xml.EndElement:
			if len(key) != 0 {
				a[key] = val
				key, val = "", ""
			}
			if tt.Name == start.Name {
				return nil
			}

		default:
			return errors.New(fmt.Sprintf("uknown %T", t))
		}
	}

	return nil
}

func nextDemo() {
	fmt.Println("\n=-= marshal and unmarshal xml =-=")

	var theMap MyMap = map[string]string{"one": "1", "two": "2", "three": "3"}
	aMap, _ := xml.MarshalIndent(&theMap, "", "    ")
	fmt.Println(string(aMap))

	var recoveredMap MyMap = make(map[string]string)

	err := xml.Unmarshal(aMap, &recoveredMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(recoveredMap)

	fmt.Println()
}

type User3 struct {
	//UserId   string `xml:"userId,omitempty"`
	UserId   string `xml:"user_id,omitempty"`
	Score    int    `xml:"score,omitempty"`
	password string `xml:"password,omitempty"`
}

type UsersArray struct {
	Users []User3 `xml:"users,omitempty"`
}

func lastDemo() {
	fmt.Println("\n=-= marshal and unmarshal xml again =-=")

	userA := User3{"Gopher", 1000, "admin"}
	userB := User3{"BigJ", 10, "1234"}
	userC := User3{UserId: "GGBoom", password: "1111"}

	db := UsersArray{[]User3{userA, userB, userC}}
	dbXML, err := xml.Marshal(&db)
	if err != nil {
		panic(err)
	}

	var recovered UsersArray
	err = xml.Unmarshal(dbXML, &recovered)
	if err != nil {
		panic(err)
	}
	fmt.Println(recovered)

	var indented []byte
	indented, err = xml.MarshalIndent(recovered, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indented))

	fmt.Println()
}
