package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type BoolTestSuite struct {
	suite.Suite
}

func (t *BoolTestSuite) TestInvalidType() {
	schema := zog.Bool()

	_, err := schema.Parse(0)
	t.Error(err)
	_, err = schema.Parse("false")
	t.Error(err)
}

func (t *BoolTestSuite) TestValid() {
	schema := zog.Bool()
	v, err := schema.Parse(true)
	t.Equal(true, v)
	t.Nil(err)

	v, err = schema.Parse(false)
	t.Equal(false, v)
	t.Nil(err)
}

func (t *BoolTestSuite) TestExact() {
	schema := zog.Bool().Exact(true)
	v, err := schema.Parse(true)
	t.Equal(true, v)
	t.Nil(err)

	_, err = schema.Parse(false)
	t.Error(err)
}

func (t *BoolTestSuite) TestRegex() {
	schema := zog.String().Regex(`[a-z]+`)

	v, err := schema.Parse("abcdefghijklmnopqrstuvwxyz")
	t.Equal("abcdefghijklmnopqrstuvwxyz", v)
	t.Nil(err)

	_, err = schema.Parse("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	t.Error(err)
}

func TestBoolSuite(t *testing.T) {
	suite.Run(t, new(BoolTestSuite))
}
