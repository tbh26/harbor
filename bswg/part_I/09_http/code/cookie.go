package code

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"
)

const (
	cookiePort = 8100
)

func CookieDemo() {
	fmt.Println("Hello code.CookieDemo() world! (ch9)")

	demoCookie()
	demoCookieJar()

	fmt.Println()
}

func demoCookie() {
	fmt.Println("\n=-= demo http cookie =-=")

	serverAddress := fmt.Sprintf("%s:%d", remoteAddress, cookiePort)
	go cookieServer(serverAddress)
	time.Sleep(time.Millisecond * 321)
	requestWithCookie(cookiePort)
	time.Sleep(time.Second * 1)

	// notice, cookie-server will keep running!
}

func requestWithCookie(port int) {
	params := url.Values{}
	params.Add("name", "Alice")
	params.Add("name", "Bob")
	u, _ := url.ParseRequestURI(fmt.Sprintf("http://%s:%d", remoteAddress, port))
	u.Path = "/cookie"
	u.RawQuery = params.Encode()

	urlStr := fmt.Sprintf("%v", u)

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{}

	c := http.Cookie{
		Name: "counter", Value: "1", Domain: remoteAddress,
		Path: "/", Expires: time.Now().AddDate(1, 0, 0)}
	req.AddCookie(&c)

	fmt.Println(" = request headers (before) =")
	for k, v := range req.Header {
		fmt.Printf("%s: %s \n", k, spaceMerge(v))
	}
	fmt.Println()

	r, e := client.Do(req)
	if e != nil {
		log.Printf("http get %q failed, %s \n", urlStr, e)
		return
	}

	fmt.Println("\n = client request / response (after) =")
	fmt.Printf("http get %q \n", urlStr)
	fmt.Printf("status: %d  ( %q ) \n", r.StatusCode, r.Status)
	fmt.Println()

	showHeaders("headers", r.Header)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("deferred body close failure;", err)
		}
	}(r.Body)
	bodyBuf := bufio.NewScanner(r.Body)
	fmt.Println(" = body =")
	for bodyBuf.Scan() {
		fmt.Println(bodyBuf.Text())
	}

}

func handleCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" =-= handle cookie request (cookie-server) =-=")

	showHeaders("server request headers", r.Header)

	counter, err := r.Cookie("counter")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value, err := strconv.Atoi(counter.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value = value + 1
	newCookie := http.Cookie{
		Name:  "counter",
		Value: strconv.Itoa(value),
	}
	http.SetCookie(w, &newCookie)
	w.WriteHeader(http.StatusOK)

	response := fmt.Sprintln("Hello cookie world.")
	w.Write([]byte(response))
	return
}

func cookieServer(address string) {
	http.HandleFunc("/cookie", handleCookie)
	panic(http.ListenAndServe(address, nil))
}

////

func demoCookieJar() {
	fmt.Println("\n=-= demo http cookie jar =-=")

	port := cookiePort
	// NO need to start cookie-server, still running, we will re-use it
	time.Sleep(time.Millisecond * 321)
	requestsWithCookieJar(port)
	time.Sleep(time.Second * 1)

}

func requestsWithCookieJar(p int) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	cookies := []*http.Cookie{
		&http.Cookie{Name: "counter", Value: "1"},
	}

	urlStr := fmt.Sprintf("http://%s:%d/cookie", remoteAddress, p)
	u, _ := url.Parse(urlStr)
	jar.SetCookies(u, cookies)

	client := http.Client{Jar: jar}

	fmt.Printf("\nurl: %v\njar: %v\n\n", u, jar)

	for i := 0; i < 5; i++ {
		r, err := client.Get(urlStr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("response %d, cookies: %v \n", r.StatusCode, r.Cookies())
		fmt.Println("Client cookie: ", jar.Cookies(u))
		//fmt.Printf("Cookie jar: %v \n", jar)
	}
}
