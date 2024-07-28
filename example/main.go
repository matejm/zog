package main

import (
	"fmt"

	"github.com/matejm/zog/zog"
)

func main() {
	var usersSchema = zog.Map().Fields(map[string]any{
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

	user, err := usersSchema.Parse(map[string]any{
		"Name": "John",
		"Age":  18,
		"Permissions": []any{
			"read",
			"write",
		},
		"PhoneNumber": "1234567890",
		"Address": []any{
			"123 Main St",
			"Apt 1",
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}
