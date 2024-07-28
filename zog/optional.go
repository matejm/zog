package zog

import (
	"reflect"
)

type optionalSchema[T any] struct {
	checks []func(*T) error
	// child schema
	schema SchemaDefinition[T]
}

// Optional creates a new optional schema.
// There are two possible usages of optional schema:
//  1. outer optional schema:
//     zog.Optional(zog.String())
//  2. chained optional schema:
//     zog.String().Optional()
//
// Both value and value pointers are valid.
func Optional[T any](schema SchemaDefinition[T]) *optionalSchema[T] {
	return &optionalSchema[T]{
		schema: schema,
	}
}

func (s *boolSchema) Optional() *optionalSchema[bool] {
	return Optional(s)
}

func (s *intSchema) Optional() *optionalSchema[int] {
	return Optional(s)
}

func (s *stringSchema) Optional() *optionalSchema[string] {
	return Optional(s)
}

func (s *objectSchema[T]) Optional() *optionalSchema[T] {
	return Optional(s)
}

func (s *arraySchema[T]) Optional() *optionalSchema[[]T] {
	return Optional(s)
}

func (s *pipeSchema[T, U]) Optional() *optionalSchema[U] {
	return Optional(s)
}

func (s *optionalSchema[T]) Parse(data any) (*T, error) {
	if data == nil {
		return nil, nil
	}

	// if data is a pointer, dereference it
	if reflect.TypeOf(data).Kind() == reflect.Ptr {
		data = reflect.ValueOf(data).Elem().Interface()
	}

	parsed, err := s.schema.Parse(data)
	if err != nil {
		return nil, err
	}
	return &parsed, check(&parsed, s.checks)
}
