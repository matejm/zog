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

func (t *OptionalTestSuite) TestBasicTypes() {
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

func (t *OptionalTestSuite) TestObject() {
	schema := zog.Object[map[string]int]().AddField("foo", zog.Int()).Optional()

	v, err := schema.Parse(map[string]interface{}{"foo": 3})
	t.Equal(map[string]int{"foo": 3}, *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func (t *OptionalTestSuite) TestTransform() {
	schema := zog.Transform(
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

func (t *OptionalTestSuite) TestPipe() {
	schema := zog.Pipe(zog.String(), zog.OneOf("a", "b")).Optional()

	v, err := schema.Parse("a")
	t.Equal("a", *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func (t *OptionalTestSuite) TestMatchAny() {
	schema := zog.MatchAny(
		zog.Int().Gt(1),
		zog.OneOf(-10),
	).Optional()

	v, err := schema.Parse(3)
	t.Equal(3, *v)
	t.Nil(err)

	v, err = schema.Parse(-10)
	t.Equal(-10, *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func (t *OptionalTestSuite) TestMatchAll() {
	schema := zog.MatchAll(
		zog.Int().Lt(1),
		zog.OneOf(-10, -20),
	).Optional()

	v, err := schema.Parse(-20)
	t.Equal(-20, *v)
	t.Nil(err)

	v, err = schema.Parse(-10)
	t.Equal(-10, *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func (t *OptionalTestSuite) TestAnyOptional() {
	schema := zog.Any().NotNil().Optional()

	v, err := schema.Parse(1)
	t.Equal(1, *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)
}

func TestOptionalSuite(t *testing.T) {
	suite.Run(t, new(OptionalTestSuite))
}
