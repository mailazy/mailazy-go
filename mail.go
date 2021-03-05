package mailazy

type SendMailRequest struct {
	Path    string
	Payload *SendMailPayload
}

type SendMailPayload struct {
	To      []string                 `json:"to"`
	From    string                   `json:"from"`
	Subject string                   `json:"subject,omitempty"`
	Content []SendMailPayloadContent `json:"content,omitempty"`
}

type SendMailPayloadContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewSendMailRequest(to, from, subject, textContent, htmlContent string) *SendMailRequest {
	return &SendMailRequest{
		Path: DefaultVersion + "/" + SendMailPath,
		Payload: &SendMailPayload{
			To:      []string{to},
			From:    from,
			Subject: subject,
			Content: []SendMailPayloadContent{{
				Type:  PlainTextContentType,
				Value: textContent,
			},{
				Type:  HtmlContentType,
				Value: htmlContent,
			}},
		},
	}
}
