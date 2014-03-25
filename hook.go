package genkins

import (
	"net/http"
	"encoding/json"
)

func GetHook(r *http.Request) (*Hook, error) {
	p := make([]byte, r.ContentLength)

	_, err := r.Body.Read(p)

	if err != nil {
		return nil, err
	}

	var h Hook
	err = json.Unmarshal(p, &h)

	if err != nil {
		return nil, err
	}

	return &h, nil
}

type Hook struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Build Build  `json:"build"`
}

