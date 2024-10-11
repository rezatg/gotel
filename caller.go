package gotel

import (
	"bytes"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type Poller interface {
	request_wrapper(url string, data *RequestData) (*Response, error)
}

type FastHTTPCaller struct{}

func (a FastHTTPCaller) request_wrapper(url string, data *RequestData) (*Response, error) {
	var req *fasthttp.Request = fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetContentType(data.ContentType)
	req.Header.SetMethod(fasthttp.MethodPost)

	req.SetRequestURI(url)
	req.SetBodyRaw(data.Buffer.Bytes())

	var resp *fasthttp.Response = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	if statusCode := resp.StatusCode(); statusCode >= fasthttp.StatusInternalServerError {
		return nil, fmt.Errorf("internal server error: %d", statusCode)
	}

	var res *Response = &Response{}

	var jsonC jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary
	// err = jsonC.Unmarshal(resp.Body(), res)

	return res, jsonC.Unmarshal(resp.Body(), res)
}

type RequestData struct {
	ContentType string
	Buffer      *bytes.Buffer
}
