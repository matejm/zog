package zog

type pipeSchema[T any, U any] struct {
	checks         []func(U) error
	schema         SchemaDefinition[T]
	transformation func(T, error) (U, error)
}

// Pipe creates a new SchemaDefinition that applies a transformation to the parsed data.
// Can be used to transform the data into a different type or capture errors.
// Pipe is used instead of something like .Transform() as chaining generic methods is not possible since
// methods in Go don't support generics.
func Pipe[T any, U any](schema SchemaDefinition[T], transformation func(T, error) (U, error)) *pipeSchema[T, U] {
	return &pipeSchema[T, U]{
		schema:         schema,
		transformation: transformation,
	}
}

func (s *pipeSchema[T, U]) Parse(data any) (U, error) {
	parsed, err := s.schema.Parse(data)

	// In all cases convert the data and error to the output type, even if we have an error.
	transformed, err := s.transformation(parsed, err)

	if err != nil {
		return transformed, err
	}

	return transformed, check(transformed, s.checks)
}
