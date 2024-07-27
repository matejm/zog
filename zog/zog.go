package zog

type SchemaDefinition[T any] interface {
	Parse(data any) (T, error)
}
