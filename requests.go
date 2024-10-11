package gotel

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

const ContentTypeJSON string = "application/json"

func (c *Client) buildRequestJson(parameters any) (*RequestData, error) {
	data := &RequestData{
		ContentType: ContentTypeJSON,
		Buffer:      &bytes.Buffer{},
	}

	err := json.NewEncoder(data.Buffer).Encode(parameters)
	if err != nil {
		return nil, fmt.Errorf("encode json: %w", err)
	}

	return data, nil
}

func (c *Client) rawRequest(method string, parameters any, vrs any) error {
	data, err := c.buildRequestJson(parameters)
	if err != nil {
		return err
	}

	resp, err := c.poller.request_wrapper(fmt.Sprintf("%s/%s", c.baseUrl, method), data)
	if err != nil {
		return err
	}

	if !resp.Ok {
		return errors.New(string(resp.Description))
	}

	var jsonC jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary
	var unmarshalErr error = jsonC.Unmarshal(resp.Result, &vrs)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil
}
