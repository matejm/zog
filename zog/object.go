package zog

import (
	"fmt"
	"reflect"
)

type objectSchema[T any] struct {
	checks []func(T) error
	// schema map for children
	schemas map[string]any
}

func Object[T any]() *objectSchema[T] {
	schemaMap := make(map[string]any)
	return &objectSchema[T]{
		schemas: schemaMap,
	}
}

// Map is a shorthand for zog.Object[map[string]any]
func Map() *objectSchema[map[string]any] {
	return Object[map[string]any]()
}

// Specify the schema for the object
func (s *objectSchema[T]) Fields(schemas map[string]any) *objectSchema[T] {
	return &objectSchema[T]{
		schemas: schemas,
	}
}

// Extend is used to extend an existing object with new fields. Can be used to add new fields to an existing object. Existing fields will be overwritten.
func (s *objectSchema[T]) Extend(schemas map[string]any) *objectSchema[T] {
	schemaMap := make(map[string]any)
	// copy the original schemas (so we don't modify the original map)
	for key, value := range s.schemas {
		schemaMap[key] = value
	}
	// Add new fields to the schema map
	for key, value := range schemas {
		schemaMap[key] = value
	}
	return &objectSchema[T]{
		schemas: schemaMap,
	}
}

// AddField is used to add a new field to an existing object.
// If the field already exists, it will be overwritten.
//
// Warning: this function copies the original schema map, so it won't modify the original map,
// if adding a lot of fields, it is recommended to add them all at once using the Extend function.
func (s *objectSchema[T]) AddField(key string, schema any) *objectSchema[T] {
	return s.Extend(map[string]any{key: schema})
}

// TODO: improve this function, should at least be a bit more readable
func (s *objectSchema[T]) Parse(data any) (T, error) {
	var v reflect.Value

	res := reflect.New(reflect.TypeOf(*new(T))).Elem()
	if res.Kind() == reflect.Map {
		// initialize map of the correct type
		resMapType := res.Type()
		res = reflect.MakeMap(reflect.MapOf(resMapType.Key(), resMapType.Elem()))
	}

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
	for key, schema := range s.schemas {
		var fieldValue reflect.Value
		var ok bool

		if v.Kind() == reflect.Map {
			// Get the value from the map if data is a map
			fromMap := v.MapIndex(reflect.ValueOf(key))
			if fromMap.IsValid() {
				fieldValue = reflect.ValueOf(fromMap.Interface())
				ok = v.MapIndex(reflect.ValueOf(key)).IsValid()
			}
		} else if v.Kind() == reflect.Struct {
			// Get the value from the struct field if data is a struct
			fieldValue = v.FieldByName(key)
			ok = fieldValue.IsValid()
		} else {
			return res.Interface().(T), ErrInvalidType(data, fmt.Sprintf("%T", new(T)))
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

		if res.Kind() == reflect.Map {
			// Set the parsed value to the corresponding field in the result map
			res.SetMapIndex(reflect.ValueOf(key), parsed)
		} else if res.Kind() == reflect.Struct {
			// Set the parsed value to the corresponding field in the result struct
			field := res.FieldByName(key)
			if !field.CanSet() {
				return res.Interface().(T), ErrCannotSetField(key)
			}
			field.Set(reflect.ValueOf(parsed.Interface()))
		} else {
			return res.Interface().(T), ErrInvalidType(data, fmt.Sprintf("%T", new(T)))
		}

	}

	// Return the populated struct
	result := res.Interface().(T)
	return result, check(result, s.checks)
}
