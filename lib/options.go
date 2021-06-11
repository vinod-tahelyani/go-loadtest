package lib

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	urlLib "net/url"
	"os"
	"time"
)

type Options struct {
	Method string
	Body string
	File string
	MaxRequests int
	Concurrency int
	RequestPerSecond int
	URL string
	KeepAlive bool
	RequestHeaders Headers
	Cookies 
}

func NewOptions() Options {
	return Options{}
}

func (options *Options) BuildRequest() *http.Request {
	var request *http.Request
	var err error
	if options.Method == http.MethodGet {
		request, err = http.NewRequest(options.Method, options.URL, nil)
	} else if options.Body == "" {
		var f *os.File
		f, err = os.Open(options.File)
		if err != nil {
			HelpAndExit("error in opening file: ", options.File)
		}
		request, err = http.NewRequest(options.Method, options.URL, f)
	} else {
		request, err = http.NewRequest(options.Method, options.URL, bytes.NewReader([]byte(options.Body)))
	}
	if err != nil {
		HelpAndExit(err.Error())
	}
	return request
}

func (options *Options) GetUrl() *urlLib.URL {
	u, err := urlLib.Parse(options.URL)
	if err != nil {
		HelpAndExit(err.Error())
	} else if u.Host == "" {
		HelpAndExit("Invalid url: ", options.URL)
	}
	return u
}

func (options *Options) ExecuteTest() Result {
	now := time.Now()
	res := NewResult()
	var clients []http.Client
	for i := 0; i < options.Concurrency; i++ {
		clients = append(clients, http.Client{
			Timeout: time.Duration(time.Second * 10),
		})
	}
	for j, n := 0, options.MaxRequests; n > 0; {
		now := time.Now()
		for i := 0; i < options.RequestPerSecond && n > 0; i++ {
			go time.AfterFunc(time.Duration(rand.Int63n(time.Second.Nanoseconds())), func(){
				latency, err := DoRequest(options.BuildRequest(), &clients[j])
				if err != nil {
					res.AddResult(latency, false)
					return
				}
				res.AddResult(latency, true)
			})
			j = (j+1) % options.Concurrency
			n--
		}
		time.Sleep(time.Until(now.Add(time.Second)))
	}
	for res.GetNoOfRequestcompleted() < options.MaxRequests {
		time.Sleep(time.Second)
	}
	
	res.PrintResult()
	fmt.Println("\nTime Taken", time.Since(now).Seconds(), "s")
	return *res
}

type Headers http.Header

func (headers Headers) String() string {
	return headers.String()
}

func (headers Headers) Set(header string) error {
	return nil
}

type Cookies []http.Cookie

func (cookies Cookies) String() string {
	return cookies.String()
}

func (cookies Cookies) Set(header string) error {
	return nil
}