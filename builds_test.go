package genkins

import (
	"net/http"
	"testing"
	"fmt"
)

func TestBuildsService_Trigger(t *testing.T) {
	setUp()
	defer tearDown()
	hitApi := false
	mux.HandleFunc("/job/test/build", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			hitApi = true
			fmt.Fprint(w, `{"status":"success"}`)
		})

	err := client.Builds.Trigger("test")

	if err != nil {
		t.Errorf("Didn't expect an error got : %v", err)
	}

	if hitApi == false {
		t.Error("Didn't hit api")
	}
}
func TestBuildsService_TriggerWithParameters(t *testing.T) {
	setUp()
	defer tearDown()
	hitApi := false
	mux.HandleFunc("/job/test/buildWithParameters", func(w http.ResponseWriter, r *http.Request) {
			if m := "POST"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}
			if qs:= "testone"; qs != r.URL.Query().Get("parameter") {
				t.Errorf("Query string tree = %v, expected %v", r.URL.Query().Get("parameter"), qs)
			}
			hitApi = true
			fmt.Fprint(w, `{"status":"success"}`)
		})

	m := map[string]string{
		"parameter": "testone",
	}

	err := client.Builds.TriggerWithParameters("test", m)

	if err != nil {
		t.Errorf("Didn't expect an error got : %v", err)
	}

	if hitApi == false {
		t.Error("Didn't hit api")
	}
}