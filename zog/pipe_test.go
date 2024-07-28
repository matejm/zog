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
		zog.Bool().True(),
		zog.OneOf(true),
	)

	v, err := schemaBool.Parse(true)
	t.Equal(true, v)
	t.Nil(err)

	_, err = schemaBool.Parse(false)
	t.Error(err)
}

func (t *PipeTestSuite) TestPracticalUsage() {
	schema := zog.Pipe(
		zog.Transform(
			zog.String().NonEmpty(),
			func(s string, err error) (int, error) {
				if err != nil {
					return 0, err
				}
				return strconv.Atoi(s)
			},
		),
		zog.Int().Gte(0),
	)

	v, err := schema.Parse("0")
	t.Equal(0, v)
	t.Nil(err)

	_, err = schema.Parse("-1")
	t.Error(err)

	_, err = schema.Parse("")
	t.Error(err)

	v, err = schema.Parse("12312312")
	t.Equal(12312312, v)
	t.Nil(err)
}

func TestPipeSuite(t *testing.T) {
	suite.Run(t, new(PipeTestSuite))
}
