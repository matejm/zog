package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type IntTestSuite struct {
	suite.Suite
}

func (t *IntTestSuite) TestInvalidType() {
	schema := zog.Int()

	_, err := schema.Parse("1")
	t.Error(err)
	_, err = schema.Parse(1.0)
	t.Error(err)
	_, err = schema.Parse([]byte("1"))
	t.Error(err)
}

func (t *IntTestSuite) TestValid() {
	schema := zog.Int()

	v, err := schema.Parse(1)
	t.Equal(1, v)
	t.Nil(err)

	v, err = schema.Parse(-2132)
	t.Equal(-2132, v)
	t.Nil(err)

	var a any = 432
	v, err = schema.Parse(a)
	t.Equal(432, v)
	t.Nil(err)
}

func (t *IntTestSuite) TestComparison() {
	// gt and lt
	schema := zog.Int().Gt(50).Lt(100)

	v, err := schema.Parse(99)
	t.Nil(err)
	t.Equal(99, v)

	_, err = schema.Parse(100)
	t.Error(err)

	v, err = schema.Parse(51)
	t.Nil(err)
	t.Equal(51, v)

	_, err = schema.Parse(50)
	t.Error(err)

	// gte and lte
	schema = zog.Int().Gte(50).Lte(100)

	v, err = schema.Parse(99)
	t.Nil(err)
	t.Equal(99, v)

	v, err = schema.Parse(100)
	t.Nil(err)
	t.Equal(100, v)

	_, err = schema.Parse(101)
	t.Error(err)

	v, err = schema.Parse(51)
	t.Nil(err)
	t.Equal(51, v)

	v, err = schema.Parse(50)
	t.Nil(err)
	t.Equal(50, v)

	_, err = schema.Parse(49)
	t.Error(err)
}

func (t *IntTestSuite) TestSign() {
	schema := zog.Int().Positive()

	v, err := schema.Parse(1)
	t.Nil(err)
	t.Equal(1, v)

	_, err = schema.Parse(-1)
	t.Error(err)

	_, err = schema.Parse(0)
	t.Error(err)

	schema = zog.Int().Negative()

	v, err = schema.Parse(-1)
	t.Nil(err)
	t.Equal(-1, v)

	_, err = schema.Parse(1)
	t.Error(err)

	_, err = schema.Parse(0)
	t.Error(err)
}

func TestIntSuite(t *testing.T) {
	suite.Run(t, new(IntTestSuite))
}
