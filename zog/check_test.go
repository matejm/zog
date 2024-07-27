package zog_test

import (
	"errors"
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type CheckTestSuite struct {
	suite.Suite
}

func (t *CheckTestSuite) TestBool() {
	schemaBool := zog.Bool().Check(func(v bool) error {
		if v {
			return nil
		}
		return errors.New("invalid")
	})

	v, err := schemaBool.Parse(true)
	t.Equal(true, v)
	t.Nil(err)

	_, err = schemaBool.Parse(false)
	t.Error(err)
}

func (t *CheckTestSuite) TestInt() {
	schemaInt := zog.Int().Check(func(v int) error {
		if v > 0 {
			return nil
		}
		return errors.New("invalid")
	})
	v, err := schemaInt.Parse(1)
	t.Equal(1, v)
	t.Nil(err)

	_, err = schemaInt.Parse(0)
	t.Error(err)
}

func (t *CheckTestSuite) TestString() {
	schemaString := zog.String().Check(func(v string) error {
		if len(v) > 0 {
			return nil
		}
		return errors.New("invalid")
	})
	v, err := schemaString.Parse("a")
	t.Equal("a", v)
	t.Nil(err)

	_, err = schemaString.Parse("")
	t.Error(err)
}

func (t *CheckTestSuite) TestObject() {
	schemaObject := zog.Object[map[string]int]().AddField("foo", zog.Int()).Check(func(v map[string]int) error {
		if v["foo"] > 0 {
			return nil
		}
		return errors.New("invalid")
	})

	v, err := schemaObject.Parse(map[string]interface{}{"foo": 3})
	t.Equal(map[string]int{"foo": 3}, v)
	t.Nil(err)

	_, err = schemaObject.Parse(map[string]interface{}{"foo": -1})
	t.Error(err)
}

func (t *CheckTestSuite) TestArray() {
	schemaArray := zog.Array(zog.Int()).Check(func(v []int) error {
		if len(v) > 0 {
			return nil
		}
		return errors.New("invalid")
	})

	v, err := schemaArray.Parse([]int{1, 2, 3})
	t.Equal([]int{1, 2, 3}, v)
	t.Nil(err)

	_, err = schemaArray.Parse([]int{})
	t.Error(err)
}

func (t *CheckTestSuite) TestOptional() {
	schema := zog.String().Optional().Check(func(v *string) error {
		if v != nil && *v == "invalid" {
			return errors.New("invalid")
		}
		return nil
	})

	v, err := schema.Parse("a")
	t.Equal("a", *v)
	t.Nil(err)

	v, err = schema.Parse(nil)
	t.Nil(v) // nil is valid
	t.Nil(err)

	_, err = schema.Parse("invalid")
	t.Error(err)
}

func TestCheckSuite(t *testing.T) {
	suite.Run(t, new(CheckTestSuite))
}
