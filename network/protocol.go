package firewall

// Protocol represents a network protocol.
type Protocol struct {
	Name string
}

// NewProtocol creates a new instance of Protocol.
func NewProtocol(name string) *Protocol {
	return &Protocol{
		Name: name,
	}
}
