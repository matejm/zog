package zog

type transformSchema[T any, U any] struct {
	checks         []func(U) error
	schema         SchemaDefinition[T]
	transformation func(T, error) (U, error)
}

// Transform creates a new SchemaDefinition that applies a transformation to the parsed data.
// Can be used to transform the data into a different type or capture errors.
//
// Transform is used as a function instead of something like .Transform() as chaining generic methods is not possible since
// methods in Go don't support generics.
func Transform[T any, U any](schema SchemaDefinition[T], transformation func(T, error) (U, error)) *transformSchema[T, U] {
	return &transformSchema[T, U]{
		schema:         schema,
		transformation: transformation,
	}
}

func (s *transformSchema[T, U]) Parse(data any) (U, error) {
	parsed, err := s.schema.Parse(data)

	// In all cases convert the data and error to the output type, even if we have an error.
	transformed, err := s.transformation(parsed, err)

	if err != nil {
		return transformed, err
	}

	return transformed, check(transformed, s.checks)
}
