package models

import (
	"log"
	"reflect"
	"testing"
)

func TestT(t *testing.T) {
	cases := []struct {
		typ interface{}
	}{
		{&DummyScheme{}},
		{DummyScheme{}},
	}

	for _, testcase := range cases {
		expected := reflect.TypeOf(testcase.typ)

		log.Println("Creating", expected)
		result := T(testcase.typ)

		if expected != reflect.TypeOf(result) {
			t.Fatal("Expected type:", expected, "Got:", result)
		}
	}
}
