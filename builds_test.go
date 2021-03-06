package genkins

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
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
		if qs := "testone"; qs != r.URL.Query().Get("parameter") {
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

func TestBuildsService_GetBuildInfo(t *testing.T) {
	setUp()
	defer tearDown()

	hitApi := false
	mux.HandleFunc("/job/test/10/api/json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		hitApi = true
		fmt.Fprint(w, `{"fullDisplayName":"test #24"}`)
	})

	b := &Build{
		Url: "job/test/10/",
	}

	info, err := client.Builds.GetInfo(b)

	if err != nil {
		t.Errorf("Didn't expect an error got : %v", err)
	}

	if hitApi == false {
		t.Error("Didn't hit api")
	}

	expected := &BuildInfo{
		FullDisplayName: "test #24",
	}

	if !reflect.DeepEqual(info, expected) {
		t.Errorf("Response body = %v, expected %v", info, expected)
	}
}

func TestBuildInfo_GetBranch_Unknown(t *testing.T) {

	b := BuildInfo{}

	name := "unknown"
	if b.GetBranchName() != name {
		t.Errorf("Expected %v got %v instead", name, b.GetBranchName())
	}

}

func TestBuildInfo_GetBranch_TwoPart(t *testing.T) {

	branches := []BuildBranch{
		BuildBranch{Name: "origin/master"},
	}

	lbr := BuildLastBuiltRevision{
		Branch: branches,
	}

	a := []BuildActions{
		BuildActions{LastBuiltRevision: lbr},
	}

	b := BuildInfo{Actions: a}

	name := "master"
	if b.GetBranchName() != name {
		t.Errorf("Expected %v got %v instead", name, b.GetBranchName())
	}

}

func TestBuildInfo_GetBranch_OnePart(t *testing.T) {

	branches := []BuildBranch{
		BuildBranch{Name: "develop"},
	}

	lbr := BuildLastBuiltRevision{
		Branch: branches,
	}

	a := []BuildActions{
		BuildActions{LastBuiltRevision: lbr},
	}

	b := BuildInfo{Actions: a}

	name := "develop"
	if b.GetBranchName() != name {
		t.Errorf("Expected %v got %v instead", name, b.GetBranchName())
	}
}
