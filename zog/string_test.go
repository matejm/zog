package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type StringTestSuite struct {
	suite.Suite
}

func (t *StringTestSuite) TestInvalidType() {
	schema := zog.String()

	_, err := schema.Parse(1)
	t.Error(err)
	_, err = schema.Parse('a')
	t.Error(err)
	_, err = schema.Parse([]byte("a"))
	t.Error(err)
}

func (t *StringTestSuite) TestValid() {
	schema := zog.String()
	v, err := schema.Parse("a")
	t.Equal("a", v)
	t.Nil(err)

	v, err = schema.Parse("")
	t.Equal("", v)
	t.Nil(err)

	var alphabet any = "abcdefghijklmnopqrstuvwxyz"
	v, err = schema.Parse(alphabet)
	t.Equal("abcdefghijklmnopqrstuvwxyz", v)
	t.Nil(err)
}

func (t *StringTestSuite) TestLength() {
	schema := zog.String().Min(5).Max(10)

	v, err := schema.Parse("abcde")
	t.Equal("abcde", v)
	t.Nil(err)

	v, err = schema.Parse("abcdefg")
	t.Equal("abcdefg", v)
	t.Nil(err)

	v, err = schema.Parse("abcdefghij")
	t.Equal("abcdefghij", v)
	t.Nil(err)

	_, err = schema.Parse("abcd")
	t.Error(err)

	_, err = schema.Parse("abcdefghijk")
	t.Error(err)

	_, err = schema.Parse("abcdefghijklmnopqrstuvwxyz")
	t.Error(err)
}

func (t *StringTestSuite) TestNonEmpty() {
	schema := zog.String().NonEmpty()

	v, err := schema.Parse("a")
	t.Equal("a", v)
	t.Nil(err)

	_, err = schema.Parse("")
	t.Error(err)
}

func (t *StringTestSuite) TestExact() {
	schema := zog.String().Exact("a")

	v, err := schema.Parse("a")
	t.Equal("a", v)
	t.Nil(err)

	_, err = schema.Parse("b")
	t.Error(err)
}

func (t *StringTestSuite) TestRegex() {
	schema := zog.String().Regex(`[a-z]+`)

	v, err := schema.Parse("abcdefghijklmnopqrstuvwxyz")
	t.Equal("abcdefghijklmnopqrstuvwxyz", v)
	t.Nil(err)

	_, err = schema.Parse("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	t.Error(err)
}

func TestStringSuite(t *testing.T) {
	suite.Run(t, new(StringTestSuite))
}
