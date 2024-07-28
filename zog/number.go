package zog

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

type numberSchema[T constraints.Integer | constraints.Float] struct {
	checks []func(T) error
}

func Int() *numberSchema[int] {
	return &numberSchema[int]{}
}

func Float() *numberSchema[float64] {
	return &numberSchema[float64]{}
}

func (s *numberSchema[T]) Parse(data any) (T, error) {
	// cast if we can (simple .(T) is not enough as we can get json data which are always float64)
	vType := reflect.TypeOf(data)
	requestedType := reflect.TypeOf(T(0))

	if vType != requestedType {
		// check if we can convert
		if !vType.ConvertibleTo(requestedType) {
			return 0, ErrInvalidType(data, requestedType.String())
		}

		// cast
		data = reflect.ValueOf(data).Convert(requestedType).Interface()
	}

	// now we can safely cast to T
	v, ok := data.(T)
	if !ok {
		return 0, ErrInvalidType(data, requestedType.String())
	}

	return v, check(v, s.checks)
}

func (s *numberSchema[T]) Gt(min T) *numberSchema[T] {
	s.checks = append(s.checks, func(v T) error {
		if v <= min {
			return ErrTooSmall(v, min, false)
		}
		return nil
	})
	return s
}

func (s *numberSchema[T]) Gte(min T) *numberSchema[T] {
	s.checks = append(s.checks, func(v T) error {
		if v < min {
			return ErrTooSmall(v, min, true)
		}
		return nil
	})
	return s
}

func (s *numberSchema[T]) Lt(max T) *numberSchema[T] {
	s.checks = append(s.checks, func(v T) error {
		if v >= max {
			return ErrTooBig(v, max, false)
		}
		return nil
	})
	return s
}

func (s *numberSchema[T]) Lte(max T) *numberSchema[T] {
	s.checks = append(s.checks, func(v T) error {
		if v > max {
			return ErrTooBig(v, max, true)
		}
		return nil
	})
	return s
}

func (s *numberSchema[T]) Positive() *numberSchema[T] {
	return s.Gt(0)
}

func (s *numberSchema[T]) Negative() *numberSchema[T] {
	return s.Lt(0)
}

func (s *numberSchema[T]) Exact(value T) *numberSchema[T] {
	s.checks = append(s.checks, func(v T) error {
		if v != value {
			return ErrExact(v, value)
		}
		return nil
	})
	return s
}
