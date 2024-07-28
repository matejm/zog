package zog_test

import (
	"fmt"
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type JSONTestSuite struct {
	suite.Suite
}

func (t *JSONTestSuite) TestSimple() {
	v, err := zog.Bool().ParseJSON([]byte(`true`))
	t.Equal(true, v)
	t.Nil(err)

	v2, err := zog.Int().ParseJSON([]byte(`1`))
	t.Equal(1, v2)
	t.Nil(err)

	v3, err := zog.String().ParseJSON([]byte(`"a"`))
	t.Equal("a", v3)
	t.Nil(err)

	v4, err := zog.String().Optional().ParseJSON([]byte(`null`))
	t.Nil(v4)
	t.Nil(err)

	_, err = zog.String().NonEmpty().ParseJSON([]byte(`""`))
	t.Error(err)

	v5, err := zog.String().NonEmpty().ParseJSON([]byte(`"a"`))
	t.Equal("a", v5)
	t.Nil(err)
}

func (t *JSONTestSuite) TestObject() {
	schemaObject := zog.Map().
		AddField("bool", zog.Bool()).
		AddField("int", zog.Int())

	v, err := schemaObject.ParseJSON([]byte(`{"bool": true, "int": 1, "str": "a"}`))
	t.Equal(true, v["bool"])
	t.Equal(1, v["int"])
	// str is not in schema
	t.Equal(nil, v["str"])

	t.Nil(err)

	v, err = schemaObject.ParseJSON([]byte(`{"bool": false, "invalid_key": 1}`))
	fmt.Printf("v: %v\n", v)
	t.Error(err)

	_, err = schemaObject.ParseJSON([]byte(`[true]`))
	t.Error(err)
}

func (t *JSONTestSuite) TestArray() {
	schema := zog.Array(
		zog.Int(),
	)

	v, err := schema.ParseJSON([]byte(`[1, 2, 3]`))
	t.Equal([]int{1, 2, 3}, v)
	t.Nil(err)

	v, err = schema.ParseJSON([]byte(`[]`))
	t.Equal([]int{}, v)
	t.Nil(err)

	_, err = schema.ParseJSON([]byte(`{"a": 1}`))
	t.Error(err)

	_, err = schema.ParseJSON([]byte(`["a"]`))
	t.Error(err)
}

func (t *JSONTestSuite) TestAggregation() {
	schema := zog.MatchAny(
		zog.Int(),
		zog.Map().AddField("a", zog.Int()),
	)

	v, err := schema.ParseJSON([]byte(`1`))
	t.Equal(1, v)
	t.Nil(err)

	v, err = schema.ParseJSON([]byte(`{"a": 1}`))
	t.Equal(map[string]any{"a": 1}, v)
	t.Nil(err)

	_, err = schema.ParseJSON([]byte(`"a"`))
	t.Error(err)

	_, err = schema.ParseJSON([]byte(`[1]`))
	t.Error(err)
}

func TestJsonSuite(t *testing.T) {
	suite.Run(t, new(JSONTestSuite))
}
