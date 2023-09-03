package code

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	remoteAddress = "127.0.0.1"
	port          = 8000
)

func ServerDemo() {
	fmt.Println("Hello code.ServerDemo() world! (ch9)")

	demoServer()
	nextDemo()

	fmt.Println()
}

func demoServer() {
	fmt.Println("\n=-= demo http server =-=")
	serverAddress := fmt.Sprintf("%s:%d", remoteAddress, port)
	go server(serverAddress)
	time.Sleep(time.Millisecond * 321)
	requestInfo()
	time.Sleep(time.Second * 1)

}

func requestInfo() {
	params := url.Values{}
	params.Add("name", "Alice")
	params.Add("name", "Bob")
	u, _ := url.ParseRequestURI(fmt.Sprintf("http://%s:%d", remoteAddress, port))
	u.Path = "/info"
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	r, e := http.Get(urlStr)
	if e != nil {
		log.Printf("http get %q failed, %s \n", urlStr, e)
		return
	}

	fmt.Println()
	fmt.Printf("http get %q \n", urlStr)
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

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" =-= request at info =-=")
	for name, headers := range r.Header {
		fmt.Println(name, headers)
	}
	greet := "world!"
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
		if k == "name" {
			if len(v) > 0 {
				greet = mergeThem(v)
			}
		}
	}
	response := fmt.Sprintf("Hello %q\n", greet)
	w.Write([]byte(response))
	return
}

func server(address string) {
	http.HandleFunc("/info", info)
	panic(http.ListenAndServe(address, nil))
}

////

func nextDemo() {
	fmt.Println("\n=-= next http server =-=")
	nextPort := port + 1
	serverAddress := fmt.Sprintf("%s:%d", remoteAddress, nextPort)
	go nextServer(serverAddress)
	time.Sleep(time.Millisecond * 321)
	nextRequests(nextPort)
	time.Sleep(time.Second * 1)

}

type MyHandler struct{}

func (c *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.RequestURI {
	case "/hello":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("goodbye\n"))
	case "/goodbye":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello\n"))
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("have a nice day\n"))
	}
}

func nextServer(address string) {
	handler := MyHandler{}
	panic(http.ListenAndServe(address, &handler))
}

func nextRequest(port int, path string) {
	u, _ := url.ParseRequestURI(fmt.Sprintf("http://%s:%d", remoteAddress, port))
	u.Path = fmt.Sprintf("/%s", path)
	urlStr := fmt.Sprintf("%v", u)

	r, e := http.Get(urlStr)
	if e != nil {
		log.Printf("http get %q failed, %s \n", urlStr, e)
		return
	}

	fmt.Println()
	fmt.Printf("http get %q \n", urlStr)
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

func nextRequests(p int) {
	paths := []string{"hello", "goodbye", "howdy"}
	for _, path := range paths {
		nextRequest(p, path)
	}
}
