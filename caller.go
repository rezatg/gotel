package gotel

import (
	"bytes"
)

type Poller interface {
	request_wrapper(url string, data *RequestData) (*Response, error)
}

type FastHTTPCaller struct{}

func (a FastHTTPCaller) request_wrapper(url string, data *RequestData) (*Response, error) {
	var res *Response = &Response{}

	return res, nil
}

type RequestData struct {
	ContentType string
	Buffer      *bytes.Buffer
}
