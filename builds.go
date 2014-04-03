package genkins

import (
	"fmt"
	"net/url"
	"strings"
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

func (s BuildsService) GetInfo(b *Build) (*BuildInfo, error) {
	url := "/" + b.Url + "api/json"
	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var info BuildInfo

	err = s.client.Do(req, &info)

	if err != nil {
		return nil, err
	}

	return &info, nil
}

type Build struct {
	Number  int    `json:"number"`
	Phase   string `json:"phase"`
	Status  string `json:"status"`
	Url     string `json:"url"`
	FullUrl string `json:"full_url"`
}

type BuildInfo struct {
	Number          int            `json:"number"`
	FullDisplayName string         `json:"fullDisplayName"`
	Result          string         `json:"result"`
	Actions         []BuildActions `json:"actions"`
	Culpirts        []BuildCulprit `json:"culprits"`
}

type BuildCulprit struct {
	FullName string `json:"fullName"`
}

type BuildActions struct {
	LastBuiltRevision BuildLastBuiltRevision `json:"lastBuiltRevision"`
}

type BuildLastBuiltRevision struct {
	Branch []BuildBranch `json:"branch"`
}

type BuildBranch struct {
	Name string `json:"name"`
}

func (bi *BuildInfo) GetBranchName() string {

	lenActions := len(bi.Actions)

	for i := 0; i < lenActions; i++ {
		action := bi.Actions[i]
		if len(action.LastBuiltRevision.Branch) > 0 {
			parts := strings.SplitN(action.LastBuiltRevision.Branch[0].Name, "/", 2)

			if len(parts) == 2 {
				return parts[1]
			} else {
				return action.LastBuiltRevision.Branch[0].Name
			}
		}
	}

	return "unknown"
}
