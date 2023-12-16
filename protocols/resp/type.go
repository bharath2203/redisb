package resp

// MessageFormat Represent the RESP format used for redis communication
type MessageFormat interface {
	Type() (MessageType, error)
	Value() (interface{}, error)
}

// MessageType Represents the format underlying that is understood by go
type MessageType int

const (
	Integer MessageType = iota + 1
	String
	Boolean
	Double
	Array
)

type message struct {
	input     string
	inputType MessageType
}

func NewMessage(input string, messageType MessageType) MessageFormat {
	return &message{
		input:     input,
		inputType: messageType,
	}
}

func (m *message) Type() (MessageType, error) {
	return m.inputType, nil
}

func (m *message) Value() (interface{}, error) {
	return m.input, nil
}
