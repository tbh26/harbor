package code

import (
	"fmt"
	"os"
	"text/template"
)

func ActionDemo() {
	fmt.Println("Hello code.ActionDemo() world! (ch9)")

	firstActionDemo()
	nextActionDemo()
	listActionDemo()
	mapActionDemo()
	lastActionDemo()

	fmt.Println()
}

type GenderUser struct {
	Name   string
	Female bool
}

const FaMsg = `
{{if .Female}}Mrs.{{- else}}Mr.{{- end}} {{.Name}},
Your package is ready.
Thanks, the Package-Boss
`

func firstActionDemo() {
	fmt.Println("\n=-= first action template =-=")

	u1 := GenderUser{"John", false}
	u2 := GenderUser{"Mary", true}

	t := template.Must(template.New("message").Parse(FaMsg))

	for _, u := range []GenderUser{u2, u1} {
		err := t.Execute(os.Stdout, u)
		if err != nil {
			fmt.Printf("first failure: %q \n", err)
		}
	}

}

type ScoreUser struct {
	Name  string
	Score uint32
}

const ScoreMsg = `
{{.Name}} your score is {{.Score}}
your level is: {{if le .Score 50}}Amateur{{else if le .Score 80}}Professional{{else}}Expert
{{end}}
`

func nextActionDemo() {
	fmt.Println("\n=-= next action template =-=")

	u1 := ScoreUser{"John", 30}
	u2 := ScoreUser{"Mary", 80}
	u3 := ScoreUser{"Alice", 84}

	t := template.Must(template.New("hello").Parse(ScoreMsg))

	for _, u := range []ScoreUser{u1, u2, u3} {
		err := t.Execute(os.Stdout, u)
		if err != nil {
			fmt.Printf("next failure: %q \n", err)
		}
	}

}

const muskMsg = `
The musketeers are:
{{range .}} - {{print .}}
{{end}}
`

const muskMsg2 = `
The musketeers are:
{{range $index, $m := .}} {{$index}}) {{$m}}
{{end}}
`

type Pair struct {
	First string
	Next  string
}

const pairMsg = `
The pairs are:
{{range .}} {{print .First}} - {{.Next}}
{{end}}
`

func listActionDemo() {
	fmt.Println("\n=-= list action template =-=")

	musketeers := []string{"Athos", "Porthos", "Aramis", "D'Artagnan"}

	t := template.Must(template.New("list").Parse(muskMsg))
	err := t.Execute(os.Stdout, musketeers)
	if err != nil {
		fmt.Printf("list failure: %q \n", err)
	}

	t = template.Must(template.New("list2").Parse(muskMsg2))
	err = t.Execute(os.Stdout, musketeers)
	if err != nil {
		fmt.Printf("list(2) failure: %q \n", err)
	}

	pairs := []Pair{{First: "France", Next: "Paris"}, {First: "Norway", Next: "Oslo"}}
	t = template.Must(template.New("list3").Parse(pairMsg))
	err = t.Execute(os.Stdout, pairs)
	if err != nil {
		fmt.Printf("list(3) failure: %q \n", err)
	}

}

const ccMsg = `
{{range $key, $val := .}}  ==  {{$key}} :  {{$val}}  ==
{{end}}
`

func mapActionDemo() {
	fmt.Println("\n=-= map action template =-=")

	cc := map[string]string{"France": "Paris", "Netherlands": "Amsterdam", "Norway": "Oslo", "Ireland": "Dublin"}

	t := template.Must(template.New("map").Parse(ccMsg))
	err := t.Execute(os.Stdout, cc)
	if err != nil {
		fmt.Printf("map failure: %q \n", err)
	}
}

const dataMsg = `
{{range $key, $val := . }} == {{ $key }} ==
  - {{ $val.Next }} in {{ $val.First }} -
{{end}}
`

const dataMsg2 = `
{{range $key, $val := . }} == {{ $key }} ==
{{ range $val}} - {{ .Next }} in {{ .First }} -
{{ end }}
{{end}}
`

func lastActionDemo() {
	fmt.Println("\n=-= last action template =-=")

	data := map[string]Pair{"France": {First: "Paris", Next: "Place de la Concorde"}, "Netherlands": {First: "Amsterdam", Next: "Rijksmuseum"}, "Norway": {First: "Oslo", Next: "Vigelandsparken"}}
	t := template.Must(template.New("last").Parse(dataMsg))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Printf("last failure: %q \n", err)
	}

	data2 := map[string][]Pair{"France": {{First: "Paris", Next: "Place de la Concorde"}, {First: "Paris", Next: "Tour Eiffel"}}, "Netherlands": {{First: "Amsterdam", Next: "Rijksmuseum"}, {First: "Rotterdam", Next: "Europoort"}}, "Norway": {{First: "Oslo", Next: "Vigelandsparken"}}}
	t = template.Must(template.New("last").Parse(dataMsg2))
	err = t.Execute(os.Stdout, data2)
	if err != nil {
		fmt.Printf("last(2) failure: %q \n", err)
	}

}
