package serializers

type Serializer interface {
	Serializer(data interface{}) ([]byte, error)
}
