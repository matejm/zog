package main

import (
	"fmt"

	"github.com/matejm/zog/zog"
)

func main() {
	// Parse a list of users
	type User struct {
		Name string
		Age  int
	}

	var schema = zog.Array(zog.Object[User]().Fields(map[string]any{
		"Name": zog.String().NonEmpty(),
		"Age":  zog.Int().Gte(0).Lte(100),
	})).Min(1)

	users, err := schema.Parse([]map[string]any{{
		"Name": "John",
		"Age":  20,
	}, {
		"Name": "Jane",
		"Age":  25,
	}})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
}
