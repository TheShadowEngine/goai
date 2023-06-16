package listening

type Listenable interface {
	MessageNames() []string
	OnMessage(m *Message) error
}
