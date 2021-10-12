package mailazy

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type SenderClient struct {
	Client *httpClient
	header http.Header
	defaultEndpoint string
}

func NewSenderClient(key, secret string, timeout time.Duration) *SenderClient {
	var httClient = &http.Client{
		Timeout: timeout,
	}
	header := http.Header{}
	header.Set(APIKeyHeaderKey, key)
	header.Set(APISecretHeaderKey, secret)

	cli := httpClient{client: *httClient}
	return &SenderClient{Client: &cli, header: header, defaultEndpoint: DefaultEndpoint}
}

type SenderClientOptions struct {
	Key        string
	Secret     string
	Endpoint   string
	timeout time.Duration
}

func NewSenderClientWithOptions(ops *SenderClientOptions) *SenderClient {
	client := NewSenderClient(ops.Secret, ops.Key, ops.timeout)
	return &SenderClient{Client: client.Client, header: client.header, defaultEndpoint: ops.Endpoint}
}

func (sc *SenderClient) Send(req *SendMailRequest) (*SendMailResponse, *SendMailError) {
	resp := new(SendMailResponse)
	errSendMail := new(SendMailError)

	if body , err := json.Marshal(req.Payload); err == nil {
		response, err := sc.Client.Post(sc.defaultEndpoint + req.Path,
			"application/json", bytes.NewBuffer(body), sc.header)
		if err != nil {
			//timeout reached
			return nil, nil
		}

		jsonData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, nil
		}

		if response.StatusCode != 202 {
			if err = json.Unmarshal(jsonData, errSendMail); err == nil {
				return nil, errSendMail
			}
		}

		if err = json.Unmarshal(jsonData, resp); err == nil {
			return resp, nil
		}
	}
	return nil, nil
}
