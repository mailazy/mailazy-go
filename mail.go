package mailazy

type SendMailRequest struct {
	Path    string
	Payload *SendMailPayload
}

type SendMailPayload struct {
	To      []string                 `json:"to"`
	From    string                   `json:"from"`
	ReplyTo string                   `json:"reply_to"`
	CC      []string                 `json:"cc"`
	BCC     []string                 `json:"bcc"`
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
			}, {
				Type:  HtmlContentType,
				Value: htmlContent,
			}},
		},
	}
}

func NewSendMailRequestWithParams(to, from, subject, textContent, htmlContent string, replyTo *string, cc, bcc []string) *SendMailRequest {
	r := NewSendMailRequest(to, from, subject, textContent, htmlContent)
	if replyTo != nil {
		r.Payload.ReplyTo = *replyTo
	}
	if len(cc) > 0 {
		r.Payload.CC = cc
	}
	if len(bcc) > 0 {
		r.Payload.BCC = bcc
	}

	return r
}
