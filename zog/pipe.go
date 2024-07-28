package zog

import "reflect"

type pipeSchema[T any, U any] struct {
	checks []func(U) error
	first  SchemaDefinition[T]
	second SchemaDefinition[U]
}

// Pipe one schema to another.
// This method supports piping only two schemas to keep the type system simple.
func Pipe[T any, U any](first SchemaDefinition[T], second SchemaDefinition[U]) *pipeSchema[T, U] {
	return &pipeSchema[T, U]{first: first, second: second}
}

func (p *pipeSchema[T, U]) Parse(data any) (U, error) {
	first, err := p.first.Parse(data)
	if err != nil {
		// get an instance of the type U for the type checker
		u := reflect.New(reflect.TypeOf(*new(U))).Elem().Interface().(U)
		return u, err
	}

	second, err := p.second.Parse(first)
	if err != nil {
		return second, err
	}

	return second, check(second, p.checks)
}
