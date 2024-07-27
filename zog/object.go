package zog

import (
	"fmt"
	"reflect"
)

type objectSchema[T any] struct {
	// schema map for children
	schema map[string]any
}

func Object[T any](schema map[string]any) *objectSchema[T] {
	return &objectSchema[T]{
		schema: schema,
	}
}

func (s *objectSchema[T]) Parse(data any) (T, error) {
	var v reflect.Value

	// Create an instance of T
	res := reflect.New(reflect.TypeOf(*new(T))).Elem()

	// Check if data is a map or struct
	switch data := data.(type) {
	case map[string]interface{}:
		v = reflect.ValueOf(data)
	case T:
		v = reflect.ValueOf(data)
	default:
		return res.Interface().(T), ErrInvalidType(data, fmt.Sprintf("%T", new(T)))
	}

	// parse recursively (TODO: check if there is a more elegant way to do this)
	for key, schema := range s.schema {
		var fieldValue reflect.Value
		var ok bool

		if v.Kind() == reflect.Map {
			// Get the value from the map if data is a map
			fieldValue = reflect.ValueOf(v.MapIndex(reflect.ValueOf(key)).Interface())
			ok = v.MapIndex(reflect.ValueOf(key)).IsValid()
		} else {
			// Get the value from the struct field if data is a struct
			fieldValue = v.FieldByName(key)
			ok = fieldValue.IsValid()
		}

		if !ok {
			return res.Interface().(T), ErrMissingField(key)
		}

		parse := reflect.ValueOf(schema).MethodByName("Parse")

		// TODO: Just a sanity check for the Parse method, should check more
		if !parse.IsValid() || parse.Type().NumIn() != 1 || parse.Type().NumOut() != 2 {
			return res.Interface().(T), ErrInvalidSchema(key)
		}

		// Parse the field value using the schema
		result := parse.Call([]reflect.Value{fieldValue})
		parsed := result[0]
		err := result[1]

		if !err.IsNil() {
			return res.Interface().(T), ErrInObject(key, err.Interface().(error))
		}

		// Set the parsed value to the corresponding field in the result struct
		field := res.FieldByName(key)
		if !field.CanSet() {
			return res.Interface().(T), ErrCannotSetField(key)
		}
		field.Set(reflect.ValueOf(parsed.Interface()))
	}

	// Return the populated struct
	result := res.Interface().(T)
	return result, nil
}

// Extend is used to extend an existing object with new fields. Can be used to add new fields to an existing object. Existing fields will be overwritten.
func (s *objectSchema[T]) Extend(schema map[string]any) *objectSchema[T] {
	// Add fields to the schema map
	for key, value := range schema {
		s.schema[key] = value
	}
	return s
}
