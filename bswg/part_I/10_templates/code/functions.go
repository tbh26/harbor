package code

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func FunctionsDemo() {
	fmt.Println("Hello code.FunctionsDemo() world! (ch10)")

	firstFuncDemo()
	nextFuncDemo()
	thirdFuncDemo()

	fmt.Println()
}

const funMsg = `
The fourth musketeer is:
{{slice . 3}}
`

func firstFuncDemo() {
	fmt.Println("\n=-= first function template =-=")

	musketeers := []string{"Athos", "Porthos", "Aramis", "D'Artagnan"}
	t := template.Must(template.New("fun").Parse(funMsg))
	err := t.Execute(os.Stdout, musketeers)
	if err != nil {
		fmt.Printf("first fun failure: %q \n", err)
	}
}

const funMsg2 = `
The musketeers are:
{{join . ", "}}
`

func nextFuncDemo() {
	fmt.Println("\n=-= next function template =-=")

	musketeers := []string{"Athos", "Porthos", "Aramis", "D'Artagnan"}
	funcs := template.FuncMap{"join": strings.Join}
	t, err := template.New("msg").Funcs(funcs).Parse(funMsg2)
	if err != nil {
		fmt.Printf("next fun prepare failure: %q \n", err)
	}
	err = t.Execute(os.Stdout, musketeers)
	if err != nil {
		fmt.Printf("next fun execute failure: %q \n", err)
	}
}

// block == define-template & template-execute
const funHeader = `
{{block "hello" .}}Hello and welcome{{end}}!

`

const funWelcome = `
{{define "musketeers"}} ------------
{{range .}} - {{print .}}
{{end}}{{end}}
{{template "hello" .}}
{{template "musketeers" .}}
{{define "divider"}} =-=-=-=-={{end}}
{{template "divider" .}}
bla bla bla
{{template "divider" .}}

{{template "musketeers" .}}
`

func thirdFuncDemo() {
	fmt.Println("\n=-= third function template =-=")

	musketeers := []string{"Athos", "Porthos", "Aramis", "D'Artagnan"}
	helloMsg, err := template.New("hello-welcome").Parse(funHeader)
	if err != nil {
		fmt.Printf("third initial prepare failure: %q \n", err)
	}
	welcomeMsg, err := template.Must(helloMsg.Clone()).Parse(funWelcome)
	if err != nil {
		fmt.Printf("third prepare extend failure: %q \n", err)
	}

	if err = helloMsg.Execute(os.Stdout, musketeers); err != nil {
		fmt.Printf("third execute hello failure: %q \n", err)
	}
	fmt.Println(" ~~~~~ ")
	if err = welcomeMsg.Execute(os.Stdout, musketeers); err != nil {
		fmt.Printf("third execute welcome failure: %q \n", err)
	}
}
