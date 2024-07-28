package zog

type SchemaDefinition[T any] interface {
	Parse(data any) (T, error)
}

// JsonSchemaDefinition is a SchemaDefinition that can be parsed from JSON (object or array)
type JsonSchemaDefinition[T any] interface {
	SchemaDefinition[T]
	ParseJson(data []byte) (T, error)
}

func Parse[T any](data any, schema SchemaDefinition[T]) (T, error) {
	return schema.Parse(data)
}
