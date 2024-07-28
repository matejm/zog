package zog

import "reflect"

type matchType string

const (
	matchAny matchType = "any"
	matchAll matchType = "all"
)

type aggregationSchema struct {
	checks    []func(any) error
	schemas   []any
	matchType matchType
}

// MatchAny returns a schema that matches any of the given schemas.
// Currently, there is support only for two schemas so the type casting is easy (we use generics).
func MatchAny[T any, U any](schema1 SchemaDefinition[T], schema2 SchemaDefinition[U]) *aggregationSchema {
	return &aggregationSchema{
		schemas:   []any{schema1, schema2},
		matchType: matchAny,
	}
}

// MatchAll returns a schema that matches all of the given schemas.
// Currently, there is support only for two schemas so the type casting is easy (we use generics).
func MatchAll[T any, U any](schema1 SchemaDefinition[T], schema2 SchemaDefinition[U]) *aggregationSchema {
	return &aggregationSchema{
		schemas:   []any{schema1, schema2},
		matchType: matchAll,
	}
}

func (s *aggregationSchema) Parse(data any) (any, error) {
	var finalResult any

	switch s.matchType {
	case matchAny:
		foundAny := false

		// find the first schema that matches
		for _, schema := range s.schemas {
			// call the schema's parse function
			result, err := parse(schema, data)

			// if error is nil, we got a match
			if err == nil {
				finalResult = result
				foundAny = true
				break
			}
		}

		if !foundAny {
			return nil, ErrNotEvenOneMatch(len(s.schemas))
		}
	case matchAll:
		// expect all schemas to match
		for index, schema := range s.schemas {
			result, err := parse(schema, data)

			// if error is not nil, we got a mismatch
			if err != nil {
				return nil, ErrNotAllMatch(index, err)
			}

			// if we got a match, update the final result
			// TODO: should we check if all the results are exactly the same?
			// currently, this will return the last result.
			finalResult = result
		}
	}

	return finalResult, check(finalResult, s.checks)
}

func parse(schema any, data any) (any, error) {
	// call the schema's parse function
	parseFunc := reflect.ValueOf(schema).MethodByName("Parse")

	returned := parseFunc.Call([]reflect.Value{reflect.ValueOf(data)})
	parsed := returned[0].Interface()
	err := returned[1].Interface()

	if err != nil {
		return parsed, err.(error)
	}

	return parsed, nil
}
