package zog

import "regexp"

type stringSchema struct {
	checks []func(string) error
}

func String() *stringSchema {
	return &stringSchema{}
}

func (s *stringSchema) Parse(data any) (string, error) {
	v, ok := data.(string)
	if !ok {
		return "", ErrInvalidType(data, "string")
	}

	return v, check(v, s.checks)
}

func (s *stringSchema) Min(minLength int) *stringSchema {
	s.checks = append(s.checks, func(v string) error {
		if len(v) < minLength {
			return ErrTooShort(len(v), minLength)
		}
		return nil
	})
	return s
}

func (s *stringSchema) Max(maxLength int) *stringSchema {
	s.checks = append(s.checks, func(v string) error {
		if len(v) > maxLength {
			return ErrTooLong(len(v), maxLength)
		}
		return nil
	})
	return s
}

func (s *stringSchema) NonEmpty() *stringSchema {
	return s.Min(1)
}

func (s *stringSchema) Exact(value string) *stringSchema {
	s.checks = append(s.checks, func(v string) error {
		if v != value {
			return ErrExact(v, value)
		}
		return nil
	})
	return s
}

func (s *stringSchema) Regex(regex string) *stringSchema {
	s.checks = append(s.checks, func(v string) error {
		if !regexp.MustCompile(regex).MatchString(v) {
			return ErrRegex(v, regex)
		}
		return nil
	})
	return s
}
