package internal

type Chat struct {
	entering chan Client
	leaving  chan Client
	msg      chan string
}

func NewChat() *Chat {
	return &Chat{
		leaving:  make(chan Client),
		msg:      make(chan string),
		entering: make(chan Client),
	}
}
