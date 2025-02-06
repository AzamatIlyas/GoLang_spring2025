package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error){
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}


func main(){
	
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error{
			fmt.Println(req.Response.Status)
			fmt.Println("redirect")
			return nil
		},
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next: http.DefaultTransport,
		},
		Timeout: time.Second*15,
	}

	resp, err := client.Get("https://wsp.kbtu.kz/")

	if err != nil {
		log.Fatal()
	}
	defer resp.Body.Close()

	fmt.Println("response status: ", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(string(body))
}