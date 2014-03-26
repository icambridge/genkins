package genkins

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"io/ioutil"
	"testing"
	"fmt"
	"reflect"
)
var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)


func setUp() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// src.gobucket client configured to use test server
	client = NewClient("", "", "theApiKey")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

// tearDown closes the test HTTP server.
func tearDown() {
	server.Close()
}



func TestClientNewRequest(t *testing.T) {
	c := NewClient("http://localhost","", "apiKey")

	type Link struct {
		Href string `json:"href"`
	}

	inURL     := "/foo"
	outURL   :=  "http://localhost/foo"
	inBody, outBody := &Link{Href: "l"}, `{"href":"l"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v) Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test that default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}
}

func TestClientNewRequest_HttpAuth(t *testing.T) {
	c := NewClient("http://localhost","username", "apiKey")

	type Link struct {
		Href string `json:"href"`
	}

	inURL     := "/foo"
	inBody:= &Link{Href: "l"}
	req, _ := c.NewRequest("GET", inURL, inBody)

	userAgent := req.Header.Get("Authorization")
	expected := "Basic dXNlcm5hbWU6YXBpS2V5"
	if expected != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, expected)
	}
}


func TestClientNewRequest_QueryString(t *testing.T) {
	c := NewClient("http://localhost","", "apiKey")

	type Link struct {
		Href string `json:"href"`
	}

	inURL     := "/foo?zjobs=One"
	outURL   :=  "http://localhost/foo?zjobs=One"
	inBody, outBody := &Link{Href: "l"}, `{"href":"l"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v) Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test that default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, c.UserAgent)
	}
}

func TestDo_GET(t *testing.T) {
	setUp()
	defer tearDown()

	type Foo struct {
		Bar string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			fmt.Fprint(w, `{"Bar":"drink"}`)
		})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(Foo)
	client.Do(req, body)

	expected := &Foo{"drink"}

	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}

}

func TestDo_POST(t *testing.T) {
	setUp()
	defer tearDown()

	type Foo struct {
		Bar string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			fmt.Fprint(w, `{"Bar":"drink"}`)
		})

	req, _ := client.NewRequest("POST", "/", nil)
	body := new(Foo)
	client.Do(req, body)

	expected := &Foo{"drink"}

	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}

}
