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

func NewSenderClient(key, secret string) *SenderClient {
	header := http.Header{}
	header.Set(APIKeyHeaderKey, key)
	header.Set(APISecretHeaderKey, secret)

	cli := httpClient{client: http.Client{}}

	return &SenderClient{Client: &cli, header: header, defaultEndpoint: DefaultEndpoint}
}

type SenderClientOptions struct {
	Key        string
	Secret     string
	Endpoint   string
	timeout *time.Duration
}

func NewSenderClientWithOptions(ops *SenderClientOptions) *SenderClient {
	client := NewSenderClient(ops.Secret, ops.Key)

	if ops.timeout != nil {
		var httClient = &http.Client{
			Timeout: *ops.timeout,
		}
		cli := httpClient{client: *httClient}
		client.Client = &cli
	}

	if ops.Endpoint != "" {
		client.defaultEndpoint = ops.Endpoint
	}
	return client
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
