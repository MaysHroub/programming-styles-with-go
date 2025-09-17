package actor

type Message []any

type Actor interface {
	Run()
	AddToMailbox(Message)
}

func Send(actor Actor, message Message) {
	actor.AddToMailbox(message)
}
