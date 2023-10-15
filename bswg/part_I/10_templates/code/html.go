package code

import (
	"fmt"
	"html/template"
	"os"
)

func HtmlDemo() {
	fmt.Println("Hello code.HtmlDemo() world! (ch10)")

	firstHtmlDemo()

	fmt.Println()
}

const Page = `
<html>
<head>
	<title>{{.Name}}'s Languages</title>
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
	<section class="container">
		<ul class="fancy-styling">
		{{range .Languages}}
			<li class="item">{{print .}}</li>
		{{end}}
		</ul>
	</section>
</body>
</html>
`

type UserExperience struct {
	Name      string
	Languages []string
}

func firstHtmlDemo() {
	fmt.Println("\n=-= (first) html template =-=")

	languages := []string{"Go!", "\"Java\" || 'Kotlin'", "C++", " ⭾ Python ⭾ :", "Javascript && Typescript"}
	u := UserExperience{"John Doe", languages}
	t := template.Must(template.New("web-demo").Parse(Page))
	err := t.Execute(os.Stdout, u)
	if err != nil {
		fmt.Printf("html template failure: %q \n", err)
	}

}
