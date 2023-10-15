package code

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func JsonDemo() {
	fmt.Println("Hello code.JsonDemo() world! (ch8)")

	JsonIntro()
	JsonRecover()
	LastJsonDemo()
	JsonWithTags()

	fmt.Println()
}

func JsonIntro() {
	fmt.Println("\n=-= json intro =-=")

	number, err := json.Marshal(42)
	if err != nil {
		//panic(err)
		fmt.Println("error: ", err)
	}
	fmt.Println(string(number))

	float, _ := json.Marshal(3.14)
	fmt.Println(string(float))

	msg, _ := json.Marshal("This is a msg!!!")
	fmt.Println(string(msg))
	fmt.Printf("msg: %#v \n", msg)

	numbers, _ := json.Marshal([]int{1, 1, 2, 3, 5, 8})
	fmt.Println(string(numbers))
	fmt.Printf("numbers: %#v \n", numbers)

	aMap, _ := json.Marshal(map[string]int{"one": 1, "two": 2})
	fmt.Println(string(aMap))
	fmt.Printf("aMap: %#v \n", aMap)
	//fmt.Printf("aMap: %+v \n", aMap)

	fmt.Println()
}

func JsonRecover() {
	fmt.Println("\n=-= json unmarshal (byte[] -> recoverThing) =-=")

	aNumber, _ := json.Marshal(42)

	var recoveredNumber int = -1
	err := json.Unmarshal(aNumber, &recoveredNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println(recoveredNumber)
	fmt.Printf("recoveredNumber: %v, aNumber: %v \n", recoveredNumber, aNumber)

	aMap, _ := json.Marshal(map[string]int{"one": 1, "two": 2})

	recoveredMap := make(map[string]int)
	err = json.Unmarshal(aMap, &recoveredMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(recoveredMap)
	fmt.Printf("recoveredMap: %v, aMap: %v \n", recoveredMap, aMap)

	fmt.Println()
}

type User struct {
	UserId   string `json:"user_id,omitempty"`
	Score    int    `json:"score,omitempty"`
	password string `json:"password,omitempty"` // lowercase!
}

func LastJsonDemo() {
	fmt.Println("\n=-= another json demo =-=")

	userA := User{"Gopher", 1000, "admin"}
	userB := User{"BigJ", 10, "1234"}
	userC := User{UserId: "GGBoom", password: "1111"}

	db := []User{userA, userB, userC}
	dbJson, err := json.Marshal(&db)
	if err != nil {
		panic(err)
	}

	fmt.Printf("dbJson as string: %s \n", string(dbJson))
	//fmt.Printf("dbJson: %v \n", dbJson)

	var indented bytes.Buffer
	err = json.Indent(&indented, dbJson, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(indented.String())

	var recovered []User
	err = json.Unmarshal(dbJson, &recovered)
	if err != nil {
		panic(err)
	}

	//fmt.Println(recovered)
	//fmt.Printf("recovered; %v \n", recovered)
	fmt.Printf("recovered; %+v \n", recovered)

	fmt.Println()
}

func Marshal(input interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	buffer.WriteString("{")
	for i := 0; i < t.NumField(); i++ {
		encodedField, err := encodeField(t.Field(i), v.Field(i))

		if err != nil {
			return nil, err
		}
		if len(encodedField) != 0 {
			if i > 0 && i <= t.NumField()-1 {
				buffer.WriteString(", ")
			}
			buffer.WriteString(encodedField)
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func encodeField(f reflect.StructField, v reflect.Value) (string, error) {

	if f.PkgPath != "" {
		return "", nil
	}

	if f.Type.Kind() != reflect.String {
		return "", nil
	}

	tag, found := f.Tag.Lookup("pretty")
	if !found {
		return "", nil
	}

	result := f.Name + ":"
	var err error = nil
	switch tag {
	case "upper":
		result = result + strings.ToUpper(v.String())
	case "lower":
		result = result + strings.ToLower(v.String())
	default:
		err = errors.New("invalid tag value")
	}
	if err != nil {
		return "", err
	}

	return result, nil
}

type User2 struct {
	UserId   string `pretty:"upper"`
	Email    string `pretty:"lower"`
	password string `pretty:"lower"`
}

type Record struct {
	Name    string `pretty:"lower" json:"name"`
	Surname string `pretty:"upper" json:"surname"`
	Age     int    `pretty:"other" json:"age"`
}

func JsonWithTags() {
	fmt.Println("\n=-= json with tag(s) =-=")

	u := User2{"John", "John@Gmail.com", "admin"}

	marSer, _ := Marshal(u)
	fmt.Println("pretty user", string(marSer))

	r := Record{"John", "Johnson", 33}
	marRec, _ := Marshal(r)
	fmt.Println("pretty rec", string(marRec))

	jsonRec, _ := json.Marshal(r)
	fmt.Println("json rec", string(jsonRec))

}
