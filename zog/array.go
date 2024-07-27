package zog

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

func (s *arraySchema[T]) Parse(data any) ([]T, error) {
	v, ok := data.([]interface{})
	if !ok {
		return nil, ErrInvalidType(data, "[]any")
	}

	// parse recursively
	res := make([]T, len(v))
	for i, item := range v {
		item, err := s.schema.Parse(item)
		if err != nil {
			return nil, ErrInArray(i, err)
		}
		res[i] = item
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
