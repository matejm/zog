# Zog

Schema validation for Go, heavily inspired by [Zod](https://github.com/colinhacks/zod).

## Object validation

Validate if all object fields are present, correct type and match the schema.

```go
var schema = zog.Object[map[string]any](map[string]any{
	"Name": zog.String().NonEmpty(),
	"Age":  zog.Int().Min(0).Max(100),
})

_, err := schema.Parse(map[string]any{
	"Name": "John",
	"Age":  "18",
})
```

## Casting and parsing any type

Zog can be used to parse any type. Build in types like `string`, `int`, `float64` and `bool`, as well as `map[string]any` and `[]any` are supported.

```go
var matrixSchema = zog.Array(zog.Array(zog.Int()))

var unknownType any = []any{[]any{1, 2, 3}, []any{4, 5, 6}}

matrix, err := matrixSchema.Parse(unknownType)
// matrix is now automatically cast to [][]int
```

Zog also supports casting to custom structs, using the `zog.Object` function.

```go
type User struct {
	Name string
	Age  int
}

var schema = zog.Object[User](map[string]any{
	"Name": zog.String().NonEmpty(),
	"Age":  zog.Int().Gte(0).Lte(100),
})

user, err := schema.Parse(map[string]any{
	"Name": "John",
	"Age":  18,
})
// user is now of type User
```

## Roadmap

- Add optional fields
- Add more validations (e.g. email, url)
- Add more types (e.g. date, support for int8, ...)
- Add code generation for custom types (if it is possible to infer the type from the schema)
