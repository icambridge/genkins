package genkins

import (
	"net/http"
	"testing"
	"fmt"
	"reflect"
)

func TestJobsService_GetAll(t *testing.T) {
	setUp()
	defer tearDown()
	hitApi := false
	mux.HandleFunc("/api/json", func(w http.ResponseWriter, r *http.Request) {
			if m := "GET"; m != r.Method {
				t.Errorf("Request method = %v, expected %v", r.Method, m)
			}

			if qs:= "jobs[name,url,color]"; qs != r.URL.Query().Get("tree") {
				t.Errorf("Query string tree = %v, expected %v", r.URL.Query().Get("tree"), qs)
			}

			hitApi = true
			fmt.Fprint(w, `
{"jobs":[{"name":"test","url":"http://localhost:8080/job/test/","color":"blue"},{"name":"test s","url":"http://localhost:8080/job/test%20s/","color":"blue"}]}`)
		})

	req, _ := client.Jobs.GetAll()

	if hitApi == false {
		t.Error("Didn't hit api")
	}

	expected := &JobView{
		Jobs: []Job{
			Job{Name: "test", Url: "http://localhost:8080/job/test/", Color: "blue"},
			Job{Name: "test s", Url: "http://localhost:8080/job/test%20s/", Color: "blue"},
		},
	}

	if !reflect.DeepEqual(req, expected) {
		t.Errorf("Response body = %v, expected %v", req, expected)
	}

}
