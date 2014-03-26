package genkins

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"strings"
)

const (
	libraryVersion = "0.1"
	userAgent      = "genkins/" + libraryVersion
)

type Client struct {
	BaseURL  *url.URL


	client *http.Client

	username string
	apiKey  string

	UserAgent string

	Jobs *JobsService
	Builds *BuildsService
}

func (c *Client) NewRequest(method string, urlString string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlString)

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(strings.TrimRight(c.BaseURL.String(), "/") + rel.String())

	if err != nil {
		return nil, err
	}



	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	if c.username != "" {
		req.SetBasicAuth(c.username, c.apiKey)
	}
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request, output interface {}) error {


	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(output)

	return nil
}


func NewClient(hostname string, usernameStr string, apiKeyStr string) *Client {
	httpClient := http.DefaultClient


	baseURL, _ := url.Parse(hostname)

	c := &Client{
		client: httpClient,
		UserAgent: userAgent,
		BaseURL: baseURL,
		apiKey: apiKeyStr,
		username: usernameStr,
	}
	c.Jobs = &JobsService{client: c}
	c.Builds = &BuildsService{client: c}
	return c
}
