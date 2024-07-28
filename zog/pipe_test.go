package zog_test

import (
	"strconv"
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type PipeTestSuite struct {
	suite.Suite
}

func (t *PipeTestSuite) TestBasic() {
	schemaBool := zog.Pipe(
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

	schemaIgnoreError := zog.Pipe(
		zog.String().NonEmpty(),
		func(s string, err error) (string, error) {
			return s + "!", nil
		},
	)

	v2, err := schemaIgnoreError.Parse("")
	t.Equal("!", v2)
	t.Nil(err)
}

func (t *PipeTestSuite) TestNested() {
	schema :=
		zog.Pipe(
			zog.Pipe(
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

func TestPipeSuite(t *testing.T) {
	suite.Run(t, new(PipeTestSuite))
}
