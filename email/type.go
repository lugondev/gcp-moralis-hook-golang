package email

type Client interface {
	Send(message Message) error
}

type Sender struct {
	Name  string
	Email string
}

type Recipient struct {
	Name  string
	Email string
}

type Message struct {
	Sender      Sender
	Recipient   Recipient
	Subject     string
	HtmlContent string
	TextContent string
	CCs         *[]Recipient
}
