package lib

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"
)

func DoRequest(request *http.Request, client *http.Client) (float64, error) {
	timers := NewTimers()
	trace := getHTTPTracer(timers)

	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))

	response, err := client.Do(request)
	if err != nil {
		// print error in a proper way
		fmt.Println(err.Error())
		return float64(client.Timeout.Milliseconds()), err
	}
	defer response.Body.Close()
	latency := timers.gotFirstResponseByte.Sub(timers.connectionStart)
	fmt.Printf("%v \n", latency)
	return float64(latency.Microseconds()) / 1000.0, nil
}

type Timers struct {
	dnsStart, dnsDone, connectionStart, connectionDone, tlsHandshakeStart, tlsHandshakeDone, gotFirstResponseByte time.Time
}

func NewTimers() *Timers {
	t := Timers{}
	return &t
}

func getHTTPTracer(timers *Timers) *httptrace.ClientTrace {
	trace := httptrace.ClientTrace {
		DNSStart: func(di httptrace.DNSStartInfo) {
			timers.dnsStart = time.Now()
		},
		DNSDone: func(di httptrace.DNSDoneInfo) {
			timers.dnsDone = time.Now()
		},
		ConnectStart: func(network, addr string) {
			timers.connectionStart = time.Now()
		},
		GotConn: func(gci httptrace.GotConnInfo) {
			timers.connectionStart = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			timers.connectionDone = time.Now()
		},
		TLSHandshakeStart: func() {
			timers.tlsHandshakeStart = time.Now()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, e error) {
			timers.tlsHandshakeDone = time.Now()
		},
		GotFirstResponseByte: func() {
			timers.gotFirstResponseByte = time.Now()
		},
	}
	return &trace
}
