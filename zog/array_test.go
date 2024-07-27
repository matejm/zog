package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type ArrayTestSuite struct {
	suite.Suite
}

func (t *ArrayTestSuite) TestInvalidType() {
	schema := zog.Array(zog.Bool())

	_, err := schema.Parse([]int{1, 2, 3})
	t.Error(err)
	_, err = schema.Parse("a, b, c")
	t.Error(err)

	schema2 := zog.Array(zog.Array(zog.Bool()))

	_, err = schema2.Parse([][]int{{1, 2, 3}})
	t.Error(err)

	_, err = schema2.Parse([]bool{false, true})
	t.Error(err)
}

func (t *ArrayTestSuite) TestValid() {
	schema := zog.Array(zog.Bool())

	v, err := schema.Parse([]bool{true, false, true})
	t.Equal([]bool{true, false, true}, v)
	t.Nil(err)

	v, err = schema.Parse([]bool{})
	t.Equal([]bool{}, v)
	t.Nil(err)

	schema2 := zog.Array(zog.String())

	v2, err := schema2.Parse([]string{"a", "b", "c"})
	t.Equal(v2, []string{"a", "b", "c"})
	t.Nil(err)

	schema3 := zog.Array(zog.Array(zog.Array(zog.Int())))

	var tensor any = [][][]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	v3, err := schema3.Parse(tensor)
	t.Equal(v3, tensor)
	t.Nil(err)
}

func (t *ArrayTestSuite) TestLength() {
	schema := zog.Array(zog.Int()).Min(2).Max(4)

	_, err := schema.Parse([]int{})
	t.Error(err)
	_, err = schema.Parse([]int{1})
	t.Error(err)
	_, err = schema.Parse([]int{1, 2})
	t.Nil(err)
	_, err = schema.Parse([]int{1, 2, 3})
	t.Nil(err)
	_, err = schema.Parse([]int{1, 2, 3, 4})
	t.Nil(err)
	_, err = schema.Parse([]int{1, 2, 3, 4, 5})
	t.Error(err)
	_, err = schema.Parse([]int{1, 2, 3, 4, 5, 6})
	t.Error(err)

	schema2 := zog.Array(zog.Array(zog.Int())).NonEmpty()

	_, err = schema2.Parse([][]int{})
	t.Error(err)
	_, err = schema2.Parse([][]int{{}})
	t.Nil(err)
}

func TestArraySuite(t *testing.T) {
	suite.Run(t, new(ArrayTestSuite))
}
