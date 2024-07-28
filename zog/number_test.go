package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type NumberTestSuite struct {
	suite.Suite
}

func (t *NumberTestSuite) TestInvalidType() {
	schema := zog.Int()

	_, err := schema.Parse("1")
	t.Error(err)
	_, err = schema.Parse([]byte("1"))
	t.Error(err)

	schema2 := zog.Float()

	_, err = schema2.Parse("1.0")
	t.Error(err)
	_, err = schema2.Parse([]byte("1.0"))
	t.Error(err)
}

func (t *NumberTestSuite) TestValid() {
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

	// float gets cast to int
	v, err = schema.Parse(1.3)
	t.Equal(1, v)
	t.Nil(err)

	schema2 := zog.Float()

	v2, err := schema2.Parse(1.2)
	t.Equal(1.2, v2)
	t.Nil(err)

	v2, err = schema2.Parse(-2132.6)
	t.Equal(-2132.6, v2)
	t.Nil(err)

	var a2 any = 432.0
	v2, err = schema2.Parse(a2)
	t.Equal(432.0, v2)
	t.Nil(err)
}

func (t *NumberTestSuite) TestComparison() {
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

func (t *NumberTestSuite) TestSign() {
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

func TestNumberSuite(t *testing.T) {
	suite.Run(t, new(NumberTestSuite))
}
