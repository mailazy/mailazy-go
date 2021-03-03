package mailazy

import (
	"github.com/kunal-saini/httpman"
)

type SenderClient struct {
	Client *httpman.Httpman
}

func NewSenderClient(key, secret string) *SenderClient {
	client := httpman.New(DefaultEndpoint).SetHeader(APIKeyHeaderKey, key).SetHeader(APISecretHeaderKey, secret)
	return &SenderClient{Client: client}
}

type SenderClientOptions struct {
	Key        string
	Secret     string
	Endpoint   string
}

func NewSenderClientWithOptions(ops *SenderClientOptions) *SenderClient {
	client := httpman.New(ops.Endpoint).SetHeader(APIKeyHeaderKey, ops.Key).SetHeader(APISecretHeaderKey, ops.Secret)
	return &SenderClient{Client: client}
}

func (sc *SenderClient) Send(req *SendMailRequest) (*SendMailResponse, *SendMailError) {
	resp := new(SendMailResponse)
	err := new(SendMailError)
	request, _ := sc.Client.NewRequest().Post(req.Path).BodyJSON(req.Payload).Decode(resp, err)
	if request.StatusCode != 202 || len(err.Error) != 0 {
		return nil, err
	}
	return resp, nil
}
