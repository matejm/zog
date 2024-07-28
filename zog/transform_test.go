package zog_test

import (
	"strconv"
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type TransformTestSuite struct {
	suite.Suite
}

func (t *TransformTestSuite) TestBasic() {
	schemaBool := zog.Transform(
		zog.Bool(),
		func(b bool, err error) (bool, error) {
			return !b, err
		},
	)

	v, err := schemaBool.Parse(true)
	t.Equal(false, v)
	t.Nil(err)

	v, err = schemaBool.Parse(false)
	t.Equal(true, v)
	t.Nil(err)

	schemaIgnoreError := zog.Transform(
		zog.String().NonEmpty(),
		func(s string, err error) (string, error) {
			return s + "!", nil
		},
	)

	v2, err := schemaIgnoreError.Parse("")
	t.Equal("!", v2)
	t.Nil(err)
}

func (t *TransformTestSuite) TestNested() {
	schema :=
		zog.Transform(
			zog.Transform(
				zog.String(),
				func(s string, err error) (int, error) {
					return len(s), err
				},
			),
			func(i int, err error) (string, error) {
				return strconv.Itoa(i), err
			},
		)

	v, err := schema.Parse("")
	t.Equal("0", v)
	t.Nil(err)

	v, err = schema.Parse("a")
	t.Equal("1", v)
	t.Nil(err)
}

func TestTransformSuite(t *testing.T) {
	suite.Run(t, new(TransformTestSuite))
}
