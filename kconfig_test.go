package kconfig

import (
	"os"
	"reflect"
	"testing"
)


// BenchmarkGetStringMap-4   	25624730	        44.3 ns/op
func BenchmarkGetStringMap(b *testing.B) {
	InitConfig("test.yaml",false)
	for i := 0; i < b.N; i++ {
		GetStringMap("arr.0.a")
	}
}

func TestYamlConfig(t *testing.T) {
	InitConfig("test.yaml",false)
	got := GetStringArray("arr2.arr3")
	want := []string{"aa","bb","cc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v",got,want )
	}

	got2 := GetStringMap("arr.0.a")
	want2 := map[string]string{"bb":"nnn","zz":"zzz"}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got2:%v, want2:%v",got2,want2)
	}

	got3 := GetInt64Array("arr4.arr5")
	want3 := []int64{1,2,3}
	if !reflect.DeepEqual(got3, want3) {
		t.Errorf("got3:%v, want3:%v",got3,want3)
	}

	got4 := GetInt64Map("arr6")
	want4 := map[string]int64{
		"aa":1,
		"bb":2,
		"cc":3,
	}
	if !reflect.DeepEqual(got4, want4) {
		t.Errorf("got4:%v, want4:%v",got4,want4)
	}
}

func TestJsonConfig(t *testing.T) {
	InitConfig("test.json",true)
	got := GetInt64("bb")
	var want int64= 11
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v",got,want )
	}
}

func TestEnvConfig(t *testing.T) {
	os.Setenv("testpath","test.yaml")
	InitConfig("xxx",false,"testpath")
	got := GetString("arr.0.a.bb")
	want := "nnn"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v",got,want )
	}
}

