package code

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func ClientDemo() {
	fmt.Println("Hello code.ClientDemo() world! (ch9)")

	demoGet()
	demoPost()
	demoPostForm()
	otherRequestDemo()

	fmt.Println()
}

func demoGet() {
	fmt.Println("\n=-= http get =-=")

	home := "https://home.tbhes.net"
	r, e := http.Get(home)

	if e != nil {
		log.Printf("http get %q failed, %s \n", home, e)
		return
	}

	fmt.Println()
	fmt.Printf("http get %q \n", home)
	fmt.Printf("status: %d  ( %q ) \n", r.StatusCode, r.Status)
	fmt.Println()

	fmt.Println(" = headers =")
	for k, v := range r.Header {
		fmt.Printf("%s: ", k)
		for _, w := range v {
			fmt.Printf("%s ", w)
		}
		fmt.Println()
	}
	fmt.Println()

	defer r.Body.Close()
	bodyBuf := bufio.NewScanner(r.Body)
	fmt.Println(" = body =")
	for bodyBuf.Scan() {
		fmt.Println(bodyBuf.Text())
	}

}

func demoPost() {
	fmt.Println("\n=-= http post =-=")

	requestPayload := []byte(`{"user": "John Doe", "email":"john.doe@gmail.com"}`)
	bodyBuf := bytes.NewBuffer(requestPayload)
	home := "https://home.tbhes.net"
	r, e := http.Post(home, "application/json", bodyBuf)

	fmt.Println()
	fmt.Printf("http post %q \n", home)
	fmt.Printf("status: %d  ( %q ) \n", r.StatusCode, r.Status)
	fmt.Println()

	if e != nil {
		log.Printf("http post %q failed, %s \n", home, e)
	} else {
		log.Println(" no error ")
	}

}

func demoPostForm() {
	fmt.Println("\n=-= http post form =-=")

	home := "https://home.tbhes.net"
	r, e := http.PostForm(home,
		url.Values{"user": {"John Doe"}, "email": {"john.doe@gmail.com"}})
	fmt.Println()
	fmt.Printf("http post form %q \n", home)
	fmt.Printf("status: %d  ( %q ) \n", r.StatusCode, r.Status)
	fmt.Println()

	if e != nil {
		log.Printf("http post form %q failed, %s \n", home, e)
	} else {
		log.Println(" no error ")
	}

}

func otherRequestDemo() {
	fmt.Println("\n=-= other http request =-=")

	requestPayload := []byte(`{"user": "John Doe", "email":"john.doe@gmail.com"}`)
	bodyBuf := bytes.NewBuffer(requestPayload)

	home := "https://home.tbhes.net"

	header := http.Header{}
	header.Add("Content-type", "application/json")
	header.Add("X-Custom-Header", "some_value")
	header.Add("User-Agent", "safe-the-world-with-go")

	request, err := http.NewRequest(http.MethodPut, home, bodyBuf)

	if err != nil {
		log.Printf("prepare new request to %s failed, %s \n", home, err)
	}

	request.Header = header

	client := http.Client{
		Timeout: time.Second * 5,
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println()
	fmt.Printf("http put request %q \n", home)
	fmt.Printf("status: %d  ( %q ) \n", response.StatusCode, response.Status)
	fmt.Println()

	if err != nil {
		log.Printf("http post form %q failed, %s \n", home, err)
	} else {
		log.Println(" no error ")
	}

}
