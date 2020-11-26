package amplitude

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const apiUrl = "https://api.amplitude.com/2/httpapi"

type Client struct {
	key                 string
	client              *http.Client
	requestOptions      *Options
	extraRequestHeaders map[string]string
}

func New(key string) *Client {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		key:    key,
		client: &client,
	}
}

func (c *Client) SetClient(client *http.Client) {
	c.client = client
}

func (c *Client) SetOptions(options *Options) {
	c.requestOptions = options
}

func (c *Client) SetExtraHeadersForRequests(headers map[string]string) {
	c.extraRequestHeaders = headers
}

func (c *Client) send(reqData request) (*Response, error) {
	valueString, marshalErr := json.Marshal(reqData)
	if marshalErr != nil {
		return nil, marshalErr
	}

	req, reqErr := http.NewRequest("POST", apiUrl, bytes.NewBuffer(valueString))
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Add("Content-type", "application/json; charset=utf-8")

	if c.extraRequestHeaders != nil && len(c.extraRequestHeaders) > 0 {
		for header, value := range c.extraRequestHeaders {
			req.Header.Add(header, value)
		}
	}

	resp, respErr := c.client.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	var amplitudeResponse Response
	respDecodeErr := json.NewDecoder(resp.Body).Decode(&amplitudeResponse)
	if respDecodeErr != nil {
		return nil, respDecodeErr
	}

	return &amplitudeResponse, nil
}

func (c *Client) SendEvents(events []Event) (*Response, error) {
	return c.send(request{
		ApiKey: c.key,
		Events: events,
	})
}
