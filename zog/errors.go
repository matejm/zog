package zog

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// common errors

func ErrInvalidType(got any, expected string) error {
	return fmt.Errorf("invalid type, got %T, expected %s", got, expected)
}

func ErrExact[T constraints.Ordered](got, exact T) error {
	return fmt.Errorf("expected %v, got %v", exact, got)
}

// string errors (some also used for arrays)

func ErrTooShort(got, min int) error {
	return fmt.Errorf("too short, got %d, min %d", got, min)
}

func ErrTooLong(got, max int) error {
	return fmt.Errorf("too long, got %d, max %d", got, max)
}

func ErrRegex(got, regex string) error {
	return fmt.Errorf("regex failed, got %s, regex %s", got, regex)
}

// int errors

func ErrTooSmall[T constraints.Ordered](got, min T, equalWasAllowed bool) error {
	if equalWasAllowed {
		return fmt.Errorf("too small, got %v, min %v", got, min)
	}
	return fmt.Errorf("too small, needs to be strictly greater than %v, got %v", min, got)
}

func ErrTooBig[T constraints.Ordered](got, max T, equalWasAllowed bool) error {
	if equalWasAllowed {
		return fmt.Errorf("too big, got %v, max %v", got, max)
	}
	return fmt.Errorf("too big, needs to be strictly less than %v, got %v", max, got)
}

// array errors

func ErrInArray(index int, err error) error {
	return fmt.Errorf("in array, index %d: %w", index, err)
}

func ErrInObject(field string, err error) error {
	return fmt.Errorf("in object, field %s: %w", field, err)
}

// object errors (invalid schema input or output struct)

// ErrCannotSetField is returned when a field cannot be set (object schema output struct is invalid)
func ErrCannotSetField(field string) error {
	return fmt.Errorf("cannot set field %s", field)
}

func ErrInvalidSchema(name string) error {
	return fmt.Errorf("invalid schema: %s", name)
}

func ErrMissingField(field string) error {
	return fmt.Errorf("missing field %s", field)
}
