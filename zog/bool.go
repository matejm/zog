package zog

type boolSchema struct {
	checks []func(bool) error
}

func Bool() *boolSchema {
	return &boolSchema{}
}

func (s *boolSchema) Parse(data any) (bool, error) {
	v, ok := data.(bool)
	if !ok {
		return false, ErrInvalidType(data, "bool")
	}
	return v, check(v, s.checks)
}

func (s *boolSchema) Exact(expected bool) *boolSchema {
	s.checks = append(s.checks, func(v bool) error {
		if v != expected {
			return ErrExact(v, expected)
		}
		return nil
	})
	return s
}

func (s *boolSchema) True() *boolSchema {
	return s.Exact(true)
}

func (s *boolSchema) False() *boolSchema {
	return s.Exact(false)
}
