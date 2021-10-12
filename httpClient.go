package mailazy

import (
	"io"
	"net/http"
)

type httpClient struct {
	client   http.Client
	apiToken string
}

func (c *httpClient) Post(url, contentType string, body io.Reader, header http.Header) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err

	}
	req.Header = header
	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-api-key", c.apiToken)
	return c.client.Do(req)
}