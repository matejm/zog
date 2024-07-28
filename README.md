# Zog

Schema validation for Go, heavily inspired by [Zod](https://github.com/colinhacks/zod).

## Basic usage

Validate if all object fields are present, correct type and match the schema.

```go
// create a new schema
var schema = zog.String().NonEmpty().Max(255).Regex("^[a-zA-Z0-9]+$")

// parse a value from any type
value, err := schema.Parse("John")
```

More complex schema example:

```go
var userSchema = zog.Map().Fields(map[string]any{
	"Name": zog.String().NonEmpty(),
	"Age":  zog.Int().Gte(0).Lte(100),
	"Permissions": zog.Array(zog.OneOf(
		"read", "write", "admin",
	)).NonEmpty(),
	"PhoneNumber": zog.String().Regex("^[0-9]{10}$").Max(10).Optional(),
	"Address": zog.MatchAny(
		zog.String().NonEmpty(),
		zog.Array(zog.String().NonEmpty()).NonEmpty(),
	),
})
```

## Parsing JSON

Zog can parse JSON into any type defined by the schema.

```go
data := `{
  "Name": "John",
  "Age": 18,
  "Permissions": ["read", "write", "admin"],
  "PhoneNumber": "1234567890",
  "Address": ["123 Main St", "Anytown, USA"]
}`

user, err := userSchema.ParseJSON(data)
// user was validated and parsed
```

## Type inference

Go typing system is not nearly as capable as TypeScript, so Zog is unable to match the Zod's type inference. However, type inference is still present for all the built-in types such as `string`, `int`, `bool`, maps and arrays of any type.

```go
var matrixSchema = zog.Array(zog.Array(zog.Float()))

var unknownType any = []any{[]any{1, 2, 3}, []any{4, 5, 6}}
matrix, err := matrixSchema.Parse(unknownType)
// matrix type is inferred to be [][]float64

var optionalSchema = zog.String().NonEmpty().Optional()

value, err := optionalSchema.Parse("John")
// value type is inferred to be *string

var transformedSchema = zog.Transform(
	zog.String().NonEmpty(),
	func (s string, err error) (int, error) {
		return len(s), nil
	},
)

value, err := transformedSchema.Parse("John")
// value type is inferred to be int
```

In case of object schemas, Zog currently provides two options. You can either use `map[string]any` or types defined by your own structs. In the future, we might add support for code generation to automatically create structs from the schema.

### 1. use `zog.Map` to get a `map[string]any` schema

```go
var schema = zog.Map().Fields(map[string]any{
	"Name": zog.String().NonEmpty(),
	"Age":  zog.Int().Gte(0).Lte(100),
})

user, err := schema.Parse(map[string]any{
	"Name": "John",
	"Age":  18,
})
// user type is inferred to be map[string]any
```

### 2. use `zog.Object` to get a custom struct schema

```go
// define a custom struct
type User struct {
	Name string
	Age  int
}

// pass the struct to the Object function
var schema = zog.Object[User]().Fields(map[string]any{
	"Name": zog.String().NonEmpty(),
	"Age":  zog.Int().Gte(0).Lte(100),
})
// or alternatively, to avoid writing map[string]any
var schema = zog.Object[User]()
	.AddField("Name", zog.String().NonEmpty())
	.AddField("Age", zog.Int().Gte(0).Lte(100))

user, err := schema.Parse(map[string]any{
	"Name": "John",
	"Age":  18,
})
// user type is inferred to be User
```

## Future plans

- Add more useful types (e.g. date, ...)
- Check if it is possible to infer more types from the schema
- Improve README and add documentation
- Add code generation for custom types (if it is possible to infer the type from the schema)
