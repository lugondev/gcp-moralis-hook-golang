package notifier

type Notifier interface {
	Notify(message Message)
}

type Message struct {
	Content string
}
