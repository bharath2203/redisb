package resp

type Serializer interface {
	Serialize(input interface{}) (string, error)
}

type serializer struct {
}

func (s serializer) Serialize(input interface{}) (string, error) {
	panic("implement me")
}
