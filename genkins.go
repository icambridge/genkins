package genkins

import (
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"strings"
	"log"
)

const (
	libraryVersion = "0.1"
	userAgent      = "genkins/" + libraryVersion
)

type Client struct {
	BaseURL  *url.URL


	client *http.Client


	ApiKey  string

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

	v := u.Query()
	v.Set("token", c.ApiKey)
	u.RawQuery = v.Encode()

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	log.Print(u.String())
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
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


func NewClient(hostname string, apiKeyStr string) *Client {
	httpClient := http.DefaultClient


	baseURL, _ := url.Parse(hostname)

	c := &Client{
		client: httpClient,
		UserAgent: userAgent,
		BaseURL: baseURL,
		ApiKey: apiKeyStr,
	}
	c.Jobs = &JobsService{client: c}
	c.Builds = &BuildsService{client: c}
	return c
}
