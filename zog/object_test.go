package zog_test

import (
	"fmt"
	"testing"

	"github.com/matejm/zog/zog"
	"github.com/stretchr/testify/suite"
)

type ObjectTestSuite struct {
	suite.Suite
}

func (t *ObjectTestSuite) TestInvalidType() {
	schema := zog.Map().AddField("foo", zog.String())

	_, err := schema.Parse("foo")
	t.Error(err)
	_, err = schema.Parse(map[string]interface{}{"bar": 3})
	t.Error(err)
	_, err = schema.Parse([]string{"foo"})
	t.Error(err)
	_, err = schema.Parse(struct{}{})
	t.Error(err)
}

func (t *ObjectTestSuite) TestValid() {
	schema := zog.Map().AddField("foo", zog.String())

	v, err := schema.Parse(map[string]interface{}{"foo": "bar"})
	t.Equal(map[string]interface{}{"foo": "bar"}, v)
	t.Nil(err)

	schema2 := zog.Map().AddField("foo", zog.String()).AddField("bar", zog.Int())
	v, err = schema2.Parse(map[string]interface{}{"foo": "bar", "bar": 3})
	t.Equal(map[string]interface{}{"foo": "bar", "bar": 3}, v)
	t.Nil(err)

	// nested object
	schema3 := zog.Map().AddField("foo", zog.Map().AddField("bar", zog.String()))
	v, err = schema3.Parse(map[string]interface{}{"foo": map[string]interface{}{"bar": "baz"}})
	t.Equal(map[string]interface{}{"foo": map[string]interface{}{"bar": "baz"}}, v)
	t.Nil(err)

	// array in object
	schema4 := zog.Map().AddField("foo", zog.Array(zog.String()))
	v, err = schema4.Parse(map[string]interface{}{"foo": []string{"bar", "baz"}})
	t.Equal(map[string]interface{}{"foo": []string{"bar", "baz"}}, v)
	t.Nil(err)
}

func (t *ObjectTestSuite) TestUnspecifiedFieldsAreOmitted() {
	schema := zog.Map().AddField("foo", zog.String())

	v, err := schema.Parse(map[string]interface{}{"foo": "bar", "bar": 3, "baz": 4})
	t.Equal(map[string]interface{}{"foo": "bar"}, v)
	t.Nil(err)
}

func (t *ObjectTestSuite) TestCustomTypes() {
	type Foo struct {
		Bar string
		Baz int
	}

	fooSchema := zog.Object[Foo]().Fields(map[string]any{
		"Bar": zog.String(),
		"Baz": zog.Int(),
	})

	v, err := fooSchema.Parse(map[string]any{"Bar": "asdf", "Baz": 3})
	t.Equal(Foo{Bar: "asdf", Baz: 3}, v)
	t.Nil(err)

	type Bar struct {
		Foo Foo
	}

	barSchema := zog.Object[Bar]().AddField("Foo", fooSchema)

	v2, err := barSchema.Parse(map[string]interface{}{"Foo": map[string]interface{}{"Bar": "bar", "Baz": 3}})
	t.Equal(v2, Bar{Foo: Foo{Bar: "bar", Baz: 3}})
	t.Nil(err)

	// object can be input
	v3, err := barSchema.Parse(Bar{Foo: Foo{Bar: "bar", Baz: 3}})
	t.Equal(v3, Bar{Foo: Foo{Bar: "bar", Baz: 3}})
	t.Nil(err)

	// also allow different types of maps (not just map[string]interface{})
	schema := zog.Object[map[string]int]().AddField("foo", zog.Int())

	v4, err := schema.Parse(map[string]interface{}{"foo": 3})
	t.Equal(map[string]int{"foo": 3}, v4)
	t.Nil(err)

	// also allow different types of maps for input
	v4, err = schema.Parse(map[string]int{"foo": 3})
	t.Equal(map[string]int{"foo": 3}, v4)
	t.Nil(err)
}

func (t *ObjectTestSuite) TestFieldExtensions() {
	initialSchema := zog.Map().Fields(map[string]any{
		"foo": zog.String(),
		"bar": zog.Int(),
	})

	// override foo
	schema := initialSchema.Extend(map[string]any{
		"foo": zog.Int(),
	})

	v, err := schema.Parse(map[string]interface{}{"foo": 3, "bar": 4})
	t.Equal(map[string]interface{}{"foo": 3, "bar": 4}, v)
	t.Nil(err)

	// override both
	schema = initialSchema.Extend(map[string]any{
		"foo": zog.Bool(),
		"bar": zog.String(),
	})

	v, err = schema.Parse(map[string]interface{}{"foo": true, "bar": "baz"})
	t.Equal(map[string]interface{}{"foo": true, "bar": "baz"}, v)
	t.Nil(err)

	// override with addField
	schema = initialSchema.AddField("foo", zog.Bool())

	fmt.Printf("schema = %+v\n", schema)

	v, err = schema.Parse(map[string]interface{}{"foo": true, "bar": 3})
	t.Equal(map[string]interface{}{"foo": true, "bar": 3}, v)
	t.Nil(err)
}

func TestObjectSuite(t *testing.T) {
	suite.Run(t, new(ObjectTestSuite))
}
