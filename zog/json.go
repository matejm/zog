package zog

import (
	"encoding/json"
)

func ParseJSON[T any](schema SchemaDefinition[T], data []byte) (T, error) {
	var parsedData any

	err := json.Unmarshal(data, &parsedData)
	if err != nil {
		// continue parsing just get instance of T, no need to actually parse it
		res, _ := schema.Parse(parsedData)
		return res, ErrInvalidJSON(err)
	}

	// actually parse with zog
	return schema.Parse(parsedData)
}

func (s *boolSchema) ParseJSON(data []byte) (bool, error) {
	return ParseJSON(s, data)
}

func (s *stringSchema) ParseJSON(data []byte) (string, error) {
	return ParseJSON(s, data)
}

func (s *numberSchema[T]) ParseJSON(data []byte) (T, error) {
	return ParseJSON(s, data)
}

func (s *arraySchema[T]) ParseJSON(data []byte) ([]T, error) {
	return ParseJSON(s, data)
}

func (s *objectSchema[T]) ParseJSON(data []byte) (T, error) {
	return ParseJSON(s, data)
}

func (s *optionalSchema[T]) ParseJSON(data []byte) (*T, error) {
	return ParseJSON(s, data)
}

func (s *transformSchema[T, U]) ParseJSON(data []byte) (U, error) {
	return ParseJSON(s, data)
}

func (s *oneOfSchema[T]) ParseJSON(data []byte) (T, error) {
	return ParseJSON(s, data)
}

func (s *pipeSchema[T, U]) ParseJSON(data []byte) (U, error) {
	return ParseJSON(s, data)
}

func (s *aggregationSchema) ParseJSON(data []byte) (any, error) {
	return ParseJSON(s, data)
}
