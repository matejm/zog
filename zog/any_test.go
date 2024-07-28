package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type AnyTestSuite struct {
	suite.Suite
}

func (t *AnyTestSuite) TestInvalidType() {
	schema := zog.Any().Exact(1)

	_, err := schema.Parse("1")
	t.Error(err)

	_, err = schema.Parse([]int{1})
	t.Error(err)

	schema = zog.Any().NotNil()

	_, err = schema.Parse(nil)
	t.Error(err)
}

func (t *AnyTestSuite) TestValid() {
	schema := zog.Any()

	v, err := schema.Parse(1)
	t.Equal(1, v)
	t.Nil(err)

	v, err = schema.Parse("abc")
	t.Equal("abc", v)
	t.Nil(err)

	v, err = schema.Parse([]int{1})
	t.Equal([]int{1}, v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v)
	t.Nil(err)

	schema = zog.Any().NotNil()

	_, err = schema.Parse(nil)
	t.Error(err)
}

func TestAnySuite(t *testing.T) {
	suite.Run(t, new(AnyTestSuite))
}
