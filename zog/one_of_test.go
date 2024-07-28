package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type OneOfTestSuite struct {
	suite.Suite
}

func (t *OneOfTestSuite) TestInvalidType() {
	schema := zog.OneOf(1, 2, 3)

	_, err := schema.Parse("1")
	t.Error(err)
}

func (t *OneOfTestSuite) TestValid() {
	schema := zog.OneOf(1, 2, 3)

	v, err := schema.Parse(1)
	t.Equal(1, v)
	t.Nil(err)

	// this passes and this is fine
	v, err = schema.Parse(1.0)
	t.Equal(1, v)
	t.Nil(err)

	// this also passes and is not as ideal
	v, err = schema.Parse(1.5)
	t.Equal(1, v)
	t.Nil(err)

	_, err = schema.Parse(6)
	t.Error(err)

	// allow custom types
	type MyEnum string
	var (
		a MyEnum = "a"
		b MyEnum = "b"
		c MyEnum = "c"
	)

	schema2 := zog.OneOf(a, b, c)

	// parse ones of the allowed values
	v2, err := schema2.Parse(a)
	t.Equal(a, v2)
	t.Nil(err)

	// parse string into correct type
	v2, err = schema2.Parse("a")
	t.Equal(a, v2)
	t.Nil(err)

	// don't parse incorrect string
	_, err = schema2.Parse("d")
	t.Error(err)

	var e MyEnum = "e"
	_, err = schema2.Parse(e)
	t.Error(err)
}

func TestOneOfSuite(t *testing.T) {
	suite.Run(t, new(OneOfTestSuite))
}
