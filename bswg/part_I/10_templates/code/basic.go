package code

import (
	"fmt"
	"os"
	"text/template"
)

func BasicDemo() {
	fmt.Println("Hello code.BasicDemo() world! (ch9)")

	firstDemo()

	fmt.Println()
}

type User struct {
	Name   string
	UserId string
	Email  string
}

const Msg = `
Dear {{.Name}},
You were registered with id: {{.UserId}}
and e-mail; {{.Email}} .

`

func firstDemo() {
	fmt.Println("\n=-= basic template =-=")

	u := User{"John Doe", "John1987", "john.doe@gmail.com"}
	t := template.Must(template.New("msg").Parse(Msg))
	err := t.Execute(os.Stdout, u)
	if err != nil {
		fmt.Printf("failure: %q \n", err)
	}

}
