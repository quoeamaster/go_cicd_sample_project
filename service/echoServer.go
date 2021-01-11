package service

import (
	"encoding/json"
	"time"
)

// EchoServer - simply echoing what it received.
type EchoServer struct{}

// Serve - handle a request.
func (e *EchoServer) Serve(msg string) (document string, err error) {
	rP := new(EchoResponse)
	rP.Msg = msg
	rP.DateTime = time.Now()
	// marshal
	bContents, err2 := json.Marshal(rP)
	if err2 != nil {
		err = err2
		return
	}
	document = string(bContents)
	return
}

// EchoResponse - the object containing the echo msg.
type EchoResponse struct {
	Msg      string    `json:"echo_msg"`
	DateTime time.Time `json:"last_update_time"`
}
