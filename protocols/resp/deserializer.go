package resp

import (
	"fmt"
	"strings"
)

type Deserializer interface {
	Deserialize(input string) (interface{}, error)
}

type deserializer struct{}

const CRLF = "\r\n"

func (d deserializer) Deserialize(input string) (interface{}, error) {
	// check if by default the input is of array type from first byte
	message := NewMessage(input, Array)

}

func processArrayType(input string) ([]*MessageFormat, error) {

}

func deserializeSimpleString(input string) (string, error) {
	// of the format "+<str>\r\n"

	if len(input) == 0 {
		return "", fmt.Errorf("empty string in input")
	}

	if input[0] != '+' {
		return "", fmt.Errorf("not a simple string")
	}

	// Trim suffix
	input = strings.TrimSuffix(input, CRLF)
	input = strings.TrimPrefix(input, "+")

	return input, nil
}
