package main

import (
	"reflect"
	"testing"
)

func TestSay(t *testing.T) {
	var actual interface{} = Say()
	var expect interface{} = "hello"
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf(`Say() => %#v; want %#v`, actual, expect)
	}
}
