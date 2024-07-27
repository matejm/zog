package zog

type intSchema struct {
	checks []func(int) error
}

func Int() *intSchema {
	return &intSchema{
		checks: []func(int) error{},
	}
}

func (s *intSchema) Parse(data any) (int, error) {
	v, ok := data.(int)
	if !ok {
		return 0, ErrInvalidType(data, "int")
	}

	return v, check(v, s.checks)
}

func (s *intSchema) Gt(min int) *intSchema {
	s.checks = append(s.checks, func(v int) error {
		if v <= min {
			return ErrTooSmall(v, min, false)
		}
		return nil
	})
	return s
}

func (s *intSchema) Gte(min int) *intSchema {
	s.checks = append(s.checks, func(v int) error {
		if v < min {
			return ErrTooSmall(v, min, true)
		}
		return nil
	})
	return s
}

func (s *intSchema) Lt(max int) *intSchema {
	s.checks = append(s.checks, func(v int) error {
		if v >= max {
			return ErrTooBig(v, max, false)
		}
		return nil
	})
	return s
}

func (s *intSchema) Lte(max int) *intSchema {
	s.checks = append(s.checks, func(v int) error {
		if v > max {
			return ErrTooBig(v, max, true)
		}
		return nil
	})
	return s
}

func (s *intSchema) Positive() *intSchema {
	return s.Gt(0)
}

func (s *intSchema) Negative() *intSchema {
	return s.Lt(0)
}

func (s *intSchema) Exact(value int) *intSchema {
	s.checks = append(s.checks, func(v int) error {
		if v != value {
			return ErrExact(v, value)
		}
		return nil
	})
	return s
}
