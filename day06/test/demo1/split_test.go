package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	data := Split("a:b:c", ":")
	array := []string{"a", "b", "c"}

	if ok := reflect.DeepEqual(data, array); !ok {
		t.Fatalf("期望得到:%v,实际得到:%v\n", array,data )
	}
}
