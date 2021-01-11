package test

import (
	"encoding/json"
	"testing"

	"github.com/quoeamaster/go_cicd_sample_project/service"
)

func TestEchoServer(t *testing.T) {
	svr := new(service.EchoServer)
	echoP := new(service.EchoResponse)
	msg := "hello World"

	if jDoc, err := svr.Serve(msg); err != nil {
		t.Fatalf("expect no problem, actual => [%v]", err)
	} else {
		// unmarshal
		json.Unmarshal([]byte(jDoc), echoP)
		if echoP.Msg != msg {
			t.Fatalf("expect [%v], actual [%v]", msg, echoP.Msg)
		}
	}
	msg = "What is 1 + 1?"
	if jDoc, err := svr.Serve(msg); err != nil {
		t.Fatalf("expect no problem, actual => [%v]", err)
	} else {
		// unmarshal
		json.Unmarshal([]byte(jDoc), echoP)
		if echoP.Msg != msg {
			t.Fatalf("expect [%v], actual [%v]", msg, echoP.Msg)
		}
	}

}
