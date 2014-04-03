package genkins

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"
	//	"os"
)

//{"name":"test","url":"job/test/","build":{"number":27,"phase":"STARTED","url":"job/test/27/"}}

func TestGetHook(t *testing.T) {

	req := &http.Request{}
	req.ContentLength = 94
	req.Body = nopCloser{bytes.NewBufferString("{\"name\":\"test\",\"url\":\"job/test/\",\"build\":{\"number\":27,\"phase\":\"STARTED\",\"url\":\"job/test/27/\"}}")}

	h, err := GetHook(req)

	if err != nil {
		t.Errorf("Didn't expect an error, but got %v", err)
	}

	expected := &Hook{
		Name:  "test",
		Url:   "job/test/",
		Build: Build{Number: 27, Phase: "STARTED", Url: "job/test/27/"},
	}
	if !reflect.DeepEqual(h, expected) {
		t.Errorf("hook = %v, expected %v", h, expected)
	}
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
