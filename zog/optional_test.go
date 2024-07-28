package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type OptionalTestSuite struct {
	suite.Suite
}

func (t *OptionalTestSuite) TestInvalidType() {
	schema := zog.Optional((zog.String()))

	_, err := schema.Parse(1)
	t.Error(err)

	number := 1
	_, err = schema.Parse(&number)
	t.Error(err)
}

func (t *OptionalTestSuite) TestValid() {
	schema := zog.Optional(zog.String())

	// both string and *string are valid
	str := "a"

	v, err := schema.Parse(str)
	t.Equal("a", *v)
	t.Nil(err)

	v, err = schema.Parse(&str)
	t.Equal("a", *v)
	t.Nil(err)

	// nil is valid
	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func (t *OptionalTestSuite) TestChained() {
	schema := zog.String().Optional()
	v, err := schema.Parse("a")
	t.Equal("a", *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)

	schema2 := zog.Bool().Optional()
	v2, err := schema2.Parse(true)
	t.Equal(true, *v2)
	t.Nil(err)

	v2, err = schema2.Parse(nil)
	t.Nil(v2)
	t.Nil(err)
}

func (t *OptionalTestSuite) TestPipe() {
	schema := zog.Pipe(
		zog.String().NonEmpty(),
		func(s string, err error) (string, error) {
			// ignore error
			return s + "!", nil
		},
	).Optional()

	v, err := schema.Parse("a")
	t.Equal("a!", *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func (t *OptionalTestSuite) TestOneOf() {
	schema := zog.OneOf("a", "b").Optional()

	v, err := schema.Parse("a")
	t.Equal("a", *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func TestOptionalSuite(t *testing.T) {
	suite.Run(t, new(OptionalTestSuite))
}
