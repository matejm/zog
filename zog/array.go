package zog

import "reflect"

type arraySchema[T any] struct {
	checks []func([]T) error
	// child schema
	schema SchemaDefinition[T]
}

func Array[T any](schema SchemaDefinition[T]) *arraySchema[T] {
	return &arraySchema[T]{
		schema: schema,
	}
}

func isArray(data any) bool {
	t := reflect.TypeOf(data)
	return t.Kind() == reflect.Array || t.Kind() == reflect.Slice
}

func (s *arraySchema[T]) Parse(data any) ([]T, error) {
	// check if this is an array
	// (simply cast to .([]interface{}) is not enough as it will fail for specific arrays like []int)
	if !isArray(data) {
		return nil, ErrInvalidType(data, "[]")
	}

	// parse recursively
	len := reflect.ValueOf(data).Len()
	res := make([]T, len)
	for i := 0; i < len; i++ {
		item := reflect.ValueOf(data).Index(i).Interface()

		parsedItem, err := s.schema.Parse(item)
		if err != nil {
			return nil, ErrInArray(i, err)
		}

		res[i] = parsedItem
	}

	// could've done checks before, but here is easier because of type inference
	return res, check(res, s.checks)
}

func (s *arraySchema[T]) Min(minLength int) *arraySchema[T] {
	s.checks = append(s.checks, func(v []T) error {
		if len(v) < minLength {
			return ErrTooShort(len(v), minLength)
		}
		return nil
	})
	return s
}

func (s *arraySchema[T]) Max(maxLength int) *arraySchema[T] {
	s.checks = append(s.checks, func(v []T) error {
		if len(v) > maxLength {
			return ErrTooLong(len(v), maxLength)
		}
		return nil
	})
	return s
}

func (s *arraySchema[T]) NonEmpty() *arraySchema[T] {
	return s.Min(1)
}
