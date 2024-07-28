package zog

// Maybe there is a better way to do this while still keeping the type safety

func (s *boolSchema) Check(check func(bool) error) *boolSchema {
	s.checks = append(s.checks, check)
	return s
}

func (s *numberSchema[T]) Check(check func(T) error) *numberSchema[T] {
	s.checks = append(s.checks, check)
	return s
}

func (s *stringSchema) Check(check func(string) error) *stringSchema {
	s.checks = append(s.checks, check)
	return s
}

func (s *objectSchema[T]) Check(check func(T) error) *objectSchema[T] {
	s.checks = append(s.checks, check)
	return s
}

func (s *arraySchema[T]) Check(check func([]T) error) *arraySchema[T] {
	s.checks = append(s.checks, check)
	return s
}

func (s *optionalSchema[T]) Check(check func(*T) error) *optionalSchema[T] {
	s.checks = append(s.checks, check)
	return s
}

func (s *transformSchema[T, U]) Check(check func(U) error) *transformSchema[T, U] {
	s.checks = append(s.checks, check)
	return s
}

func (s *oneOfSchema[T]) Check(check func(T) error) *oneOfSchema[T] {
	s.checks = append(s.checks, check)
	return s
}

func (s *pipeSchema[T, U]) Check(check func(U) error) *pipeSchema[T, U] {
	s.checks = append(s.checks, check)
	return s
}

func (s *aggregationSchema) Check(check func(any) error) *aggregationSchema {
	s.checks = append(s.checks, check)
	return s
}
