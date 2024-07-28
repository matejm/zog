package zog

import (
	"reflect"
)

type oneOfSchema[T comparable] struct {
	checks []func(T) error
	oneOf  []T
}

// OneOf checks if value is one of the allowed values.
// Automatically casts value to type T.
// Examples of automatic casting:
//   - string to type MyEnum string
//   - int and float to type int
//   - etc.
//
// Warning: this means that the following code will pass:
//
//	schema := zog.OneOf(1, 2, 3)
//	v, err := schema.Parse(1)    // will pass as expected
//	v, err := schema.Parse(1.5)  // will be casted to int 1 so it will pass
func OneOf[T comparable](oneOf ...T) *oneOfSchema[T] {
	return &oneOfSchema[T]{oneOf: oneOf}
}

func (s *oneOfSchema[T]) Parse(data any) (T, error) {
	// check if value is of type T
	// casting with data.(T) will fail for custom types
	// (for example we wan't to allow automatic parsing string into type MyEnum string)
	reflectValue := reflect.ValueOf(data)
	targetType := reflect.TypeOf(*new(T))

	if !reflectValue.Type().ConvertibleTo(targetType) {
		return reflect.New(reflect.TypeOf(*new(T))).Elem().Interface().(T), ErrInvalidType(data, reflect.TypeOf(*new(T)).Name())
	}

	v := reflectValue.Convert(targetType).Interface().(T)

	// check if value is one of the allowed values
	for _, o := range s.oneOf {
		if v == o {
			// success
			return v, check(v, s.checks)
		}
	}

	return v, ErrNoOneOf(v)
}
