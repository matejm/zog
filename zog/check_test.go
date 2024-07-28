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

func (t *CheckTestSuite) TestFloat() {
	schemaFloat := zog.Float().Check(func(v float64) error {
		if v > 0 {
			return nil
		}
		return errors.New("invalid")
	})
	v, err := schemaFloat.Parse(1.5)
	t.Equal(1.5, v)
	t.Nil(err)

	_, err = schemaFloat.Parse(0.0)
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

func (t *CheckTestSuite) TestTransform() {
	schema := zog.Transform(
		zog.String(),
		func(s string, err error) (string, error) {
			return s + "!", err
		},
	).Check(func(v string) error {
		if v == "!" {
			return nil
		}
		return errors.New("invalid")
	})

	v, err := schema.Parse("")
	t.Equal("!", v)
	t.Nil(err)

	_, err = schema.Parse("a")
	t.Error(err)
}

func (t *CheckTestSuite) TestOneOf() {
	schema := zog.OneOf(1, 2, 3).Check(func(v int) error {
		if v >= 2 {
			return nil
		}
		return errors.New("invalid")
	})

	v, err := schema.Parse(2)
	t.Equal(2, v)
	t.Nil(err)

	_, err = schema.Parse(1)
	t.Error(err)
}

func (t *CheckTestSuite) TestPipe() {
	schema := zog.Pipe(
		zog.Int().Gt(1).Lt(10),
		zog.OneOf(1, 2, 3),
	)

	// first schema fails
	_, err := schema.Parse(1)
	t.Error(err)

	// second schema fails
	_, err = schema.Parse(4)
	t.Error(err)

	// both schemas pass
	v, err := schema.Parse(2)
	t.Equal(2, v)
	t.Nil(err)
}

func (t *CheckTestSuite) TestMatchAny() {
	schema := zog.MatchAny(
		zog.Int().Gte(1),
		zog.OneOf(-10),
	).Check(func(v any) error {
		if v.(int) >= 2 {
			return nil
		}
		return errors.New("invalid")
	})

	v, err := schema.Parse(3)
	t.Equal(3, v)
	t.Nil(err)

	// doesn't match anymore because of the check
	_, err = schema.Parse(-10)
	t.Error(err)
}

func (t *CheckTestSuite) TestMatchAll() {
	schema := zog.MatchAll(
		zog.Int().Lt(1),
		zog.OneOf(-10, -20),
	).Check(func(v any) error {
		if v != -10 {
			return nil
		}
		return errors.New("invalid")
	})

	v, err := schema.Parse(-20)
	t.Equal(-20, v)
	t.Nil(err)

	// doesn't match anymore because of the check
	_, err = schema.Parse(-10)
	t.Error(err)
}

func (t *CheckTestSuite) TestAnyCheck() {
	schema := zog.Any().Check(func(v any) error {
		if v.(int) < 0 {
			return zog.ErrExact(v, 0)
		}
		return nil
	})

	v, err := schema.Parse(0)
	t.Equal(0, v)
	t.Nil(err)

	_, err = schema.Parse(-1)
	t.Error(err)
}

func TestCheckSuite(t *testing.T) {
	suite.Run(t, new(CheckTestSuite))
}
