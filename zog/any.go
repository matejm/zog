package zog

type anySchema struct {
	checks []func(any) error
}

func Any() *anySchema {
	return &anySchema{}
}

func (s *anySchema) Parse(data any) (any, error) {
	return data, check(data, s.checks)
}

func (s *anySchema) Exact(expected any) *anySchema {
	s.checks = append(s.checks, func(v any) error {
		if v != expected {
			return ErrExact(v, expected)
		}
		return nil
	})
	return s
}

func (s *anySchema) NotNil() *anySchema {
	s.checks = append(s.checks, func(v any) error {
		if v == nil {
			return ErrNotNil()
		}
		return nil
	})
	return s
}
