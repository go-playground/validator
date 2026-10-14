package validator

import (
	"context"
	"reflect"
	"testing"
)

func TestPartialSelectedSubtree(t *testing.T) {
	type Leaf struct {
		Name  string `validate:"required"`
		Other string `validate:"required"`
	}
	type Order struct {
		Item      *Leaf `validate:"required"`
		ItemExtra Leaf
		Items     []Leaf          `validate:"dive"`
		Lookup    map[string]Leaf `validate:"dive"`
	}
	v := New()
	input := Order{Item: &Leaf{}, Items: []Leaf{{}, {}}, Lookup: map[string]Leaf{"a": {}, "b": {}}}
	for _, tc := range []struct {
		name   string
		fields []string
		want   []string
	}{
		{"whole-child", []string{"Item"}, []string{"Order.Item.Name", "Order.Item.Other"}},
		{"child-leaf", []string{"Item.Name"}, []string{"Order.Item.Name"}},
		{"overlap", []string{"Item", "Item.Name", "Item"}, []string{"Order.Item.Name", "Order.Item.Other"}},
		{"slice-element", []string{"Items[0]"}, []string{"Order.Items[0].Name", "Order.Items[0].Other"}},
		{"map-element", []string{"Lookup[a]"}, []string{"Order.Lookup[a].Name", "Order.Lookup[a].Other"}},
		{"empty", nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := v.StructPartialCtx(context.Background(), input, tc.fields...)
			var got []string
			if err != nil {
				for _, field := range err.(ValidationErrors) {
					got = append(got, field.StructNamespace())
				}
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
	// Reuse the pool through other entry points, then make a narrow selection again.
	_ = v.Struct(input)
	_ = v.StructExcept(input, "Item")
	_ = v.StructFiltered(input, func([]byte) bool { return true })
	if err := v.StructPartial(input, "Item.Other"); err == nil || len(err.(ValidationErrors)) != 1 {
		t.Fatalf("pooled selection leaked: %v", err)
	}
}
