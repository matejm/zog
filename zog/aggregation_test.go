package zog_test

import (
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type AggregationTestSuite struct {
	suite.Suite
}

func (t *AggregationTestSuite) TestMatchAny() {
	schemaBool := zog.MatchAny(
		zog.Int(),
		zog.Array(zog.Int()),
	)

	v, err := schemaBool.Parse(3)
	t.Equal(3, v)
	t.Nil(err)

	v, err = schemaBool.Parse([]int{1, 2, 3})
	t.Equal([]int{1, 2, 3}, v)
	t.Nil(err)

	_, err = schemaBool.Parse([]string{"a", "b", "c"})
	t.NotNil(err)

	_, err = schemaBool.Parse(true)
	t.NotNil(err)
}

func (t *AggregationTestSuite) TestMatchAll() {
	schema := zog.MatchAll(
		zog.String().NonEmpty().Regex("^[a-z]+$"),
		zog.OneOf("a", "b", "c", "5"),
	)

	// match both
	v, err := schema.Parse("a")
	t.Equal("a", v)
	t.Nil(err)

	// match first only
	_, err = schema.Parse("e")
	t.Error(err)

	// match second only
	_, err = schema.Parse("5")
	t.Error(err)

	// match none
	_, err = schema.Parse("123")
	t.Error(err)
}

func TestAggregationSuite(t *testing.T) {
	suite.Run(t, new(AggregationTestSuite))
}
