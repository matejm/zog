package main

import (
	"fmt"

	"github.com/matejm/zog/zog"
)

func main() {
	// define a custom struct
	type User struct {
		Name string
		Age  int
	}

	// or alternatively, to avoid writing map[string]any
	var schema = zog.Object[User]().AddField("Name", zog.String().NonEmpty()).AddField("Age", zog.Int().Gt(0).Lt(100))

	user, err := schema.Parse(map[string]any{
		"Name": "John",
		"Age":  18,
	})
	// user type is inferred to be User

	fmt.Println(user)
	fmt.Println(err)
}
