package genkins

import (
	"fmt"
	"net/url"
)

type BuildsService struct {
	client *Client
}

func (s BuildsService) Trigger(job string) error {

	url := fmt.Sprintf("/job/%s/build", job)

	req, err := s.client.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	err = s.client.Do(req, nil)

	if err != nil {
		return err
	}
	// Could do return err, but this seems clearer.
	return nil
}

func (s BuildsService) TriggerWithParameters(job string, parameters map[string]string) error {

	v := url.Values{}

	for key, value := range parameters {
		v.Add(key, value)
	}

	url := fmt.Sprintf("/job/%s/buildWithParameters?%s", job, v.Encode())

	req, err := s.client.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	err = s.client.Do(req, nil)

	if err != nil {
		return err
	}
	// Could do return err, but this seems clearer.
	return nil
}

type Hook struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Build Build  `json:"build"`
}

type Build struct {
	Number  int    `json:"number"`
	Phase   string `json:"phase"`
	Status  string `json:"status"`
	Url     string `json:"url"`
	FullUrl string `json:"full_url"`
}
