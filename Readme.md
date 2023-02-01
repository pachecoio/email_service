# Email service

This is an example of a simple email service showcasing a design that handles different integrations of email providers through abstractions.

It also includes support to use Mailgun as an email provider as well as some examples.

The service is built using [Go](https://golang.org/)

## Prerequisites

- [Go](https://golang.org/)

## Running the service

- Clone the repository
- Create a `.env` file with the configuration of your desired service
  - By default, the main file is running an example with Mailgun so it requires the mailgun API key and domain
    ```
    MAILGUN_API_KEY=your-mailgun-api-key
    MAILGUN_DOMAIN=your-mailgun-domain
    ```
- Run `go run main.go`
- The service will be running on port 5000

### Sending email
- To send an email, make a POST request to `http://localhost:5000` with the following payload:
    ```json
  {
  	"from": "EMAIL FROM",
  	"to": "EMAIL TO",
  	"subject": "Test subject",
  	"body": "Test body"
  }
   ```
  
### Sending batch of emails
- To send a batch of emails, make a POST request to `http://localhost:5000/batch` with the following payload:
  ```json
  {
    "emails": [
      {
        "from": "EMAIL FROM",
        "to": "EMAIL TO",
        "subject": "Test subject",
        "body": "Test body"
      }
    ]
  }
  ```

### Extending the service

To add a new email provider, you can create a new `Email Client` adapter that implements the `AbstractEmailClient` interface.

E.g.
    
```go
type CustomEmailClient struct {
    // Add any required fields
}

// Add required methods to initiate your client

func (c *CustomEmailClient) Send(payload *domain.EmailPayload) error {
	// Implement the logic to send the email
    return nil
}

```


## Running the tests

- Clone the repository
- Run `go test ./...`
