package code

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	strutils "github.com/torden/go-strutil"
)

const (
	middlewarePort = 8200
)

func MiddlewareDemo() {
	fmt.Println("Hello code.MiddlewareDemo() world! (ch9)")

	middlewareDemo()
	otherMiddleware()

	fmt.Println()
}

func middlewareDemo() {
	fmt.Println("\n=-= demo http middleware =-=")

	serverAddress := fmt.Sprintf("%s:%d", remoteAddress, middlewarePort)
	go middlewareServer(serverAddress)
	time.Sleep(time.Millisecond * 321)
	someRequests(middlewarePort)
	time.Sleep(time.Second * 1)

}

func someRequest(port int, phrase string) {
	u, _ := url.ParseRequestURI(fmt.Sprintf("http://%s:%d", remoteAddress, port))
	u.Path = "/bla"
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Printf("prepare http request %q failed, %s \n", urlStr, err)
		return
	}

	req.SetBasicAuth(phrase, phrase)

	res, err := client.Do(req)

	fmt.Println()
	fmt.Printf("http get %q \n", urlStr)
	fmt.Printf("status: %d  ( %q ) \n", res.StatusCode, res.Status)
	fmt.Println()

	if err != nil {
		log.Printf("http request failed; %s \n", err)
		return
	}

	showHeaders("headers", res.Header)

	defer res.Body.Close()
	bodyBuf := bufio.NewScanner(res.Body)
	fmt.Println(" = body =")
	for bodyBuf.Scan() {
		fmt.Println(bodyBuf.Text())
	}
	fmt.Println()

}

func someRequests(port int) {
	words := []string{"Passenger Car", "Race Car"}
	//words = append(words, "😃🙃 😆🙃 😃")
	for _, phrase := range words {
		someRequest(port, phrase)
		time.Sleep(time.Millisecond * 468)
	}
}

type SomeHandler struct{}

func (sh *SomeHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	message := "Hello world."
	w.Write([]byte(message))
	return
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authType := strings.Split(header, " ")
		fmt.Println(authType)
		if len(authType) != 2 || authType[0] != "Basic" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		credentials, err := base64.StdEncoding.DecodeString(authType[1])
		if err != nil {
			fmt.Printf("(base64) decode failure %q \n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		parts := strings.Split(string(credentials), ":")
		if len(parts) != 2 || len(parts[0]) == 0 {
			fmt.Println("invalid credentials")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			firstPart := strings.ReplaceAll(parts[0], " ", "")
			nextPart := strings.ReplaceAll(parts[1], " ", "")
			fmt.Printf("firstPart: %q, nextPart: %q\n", firstPart, nextPart)
			strProc := strutils.NewStringProc()
			if strings.ToLower(firstPart) == strProc.ReverseUnicode(strings.ToLower(nextPart)) {
				next.ServeHTTP(w, r)
			} else {
				fmt.Println("invalid (user)name / password")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
	})
}

func middlewareServer(address string) {
	someHandler := &SomeHandler{}
	panic(http.ListenAndServe(address, AuthMiddleware(someHandler)))
}

////

type Middleware func(http.Handler) http.Handler

func ApplyMiddleware(h http.Handler, middleware ...Middleware) http.Handler {
	for _, next := range middleware {
		h = next(h)
	}
	return h
}

func SimpleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := w.Header().Get("simple")
		if value == "" {
			value = "X"
		} else {
			value = value + "X"
		}
		w.Header().Set("simple", value)
		next.ServeHTTP(w, r)
	})
}

func otherMiddleware() {
	fmt.Println("\n=-= other http middleware demo =-=")

	port := middlewarePort + 1
	serverAddress := fmt.Sprintf("%s:%d", remoteAddress, port)
	go otherMiddlewareServer(serverAddress)
	time.Sleep(time.Millisecond * 321)
	someOtherRequests(port)
	time.Sleep(time.Second * 1)

}

func otherMiddlewareServer(address string) {

	fmt.Printf(" = other middleware server @ %q = \n", address)

	sh := &SomeHandler{}
	http.Handle("/three", ApplyMiddleware(sh, SimpleMiddleware, SimpleMiddleware, SimpleMiddleware))
	http.Handle("/one", ApplyMiddleware(sh, SimpleMiddleware))
	panic(http.ListenAndServe(address, nil))
}

func someOtherRequests(port int) {
	paths := []string{"one", "three"}
	//paths = append(paths, "one", "three")
	for _, path := range paths {
		someOtherRequest(port, path)
		time.Sleep(time.Millisecond * 357)
	}
}

func someOtherRequest(port int, path string) {
	u, _ := url.ParseRequestURI(fmt.Sprintf("http://%s:%d", remoteAddress, port))
	u.Path = fmt.Sprintf("/%s", path)
	urlStr := fmt.Sprintf("%v", u)

	res, err := http.Get(urlStr)

	fmt.Println()
	fmt.Printf("http get %q \n", urlStr)
	fmt.Printf("status: %d  ( %q ) \n", res.StatusCode, res.Status)
	fmt.Println()

	if err != nil {
		log.Printf("http request failed; %s \n", err)
		return
	}

	showHeaders("headers", res.Header)

	defer res.Body.Close()
	bodyBuf := bufio.NewScanner(res.Body)
	fmt.Println(" = body =")
	for bodyBuf.Scan() {
		fmt.Println(bodyBuf.Text())
	}
	fmt.Println()

}
