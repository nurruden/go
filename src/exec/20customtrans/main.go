package main

import (
	"io"
	"net/http"
	"os"
)

type OurCustomTransport struct{
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response,error){
	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client{
	return &http.Client{Transport:t}
}

func main(){
	t := &OurCustomTransport{

	}
	c := t.Client()
	resp,err := c.Get("http://example.com")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)
}