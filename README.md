![Mailazy Logo](logo.png)

### Prerequisites ###
- Go version 1.14
- The Mailazy account, accessible [here](https://app.mailazy.com/signup?source=mailazy-go).

### Usage ###

```go
import "github.com/mailazy/mailazy-go"
```

### Example Usage ###

```go
senderClient := mailazy.NewSenderClient("access_key", "access_secret")
to := "recipient@mail7.io"
from := "Sender <no-reply@example.com>"
subject := "Sending with mailazy is easy!"
textContent := ""
htmlContent := ""
req := mailazy.NewSendMailRequest(to, from, subject, textContent, htmlContent)

resp, err := senderClient.Send(req)
```