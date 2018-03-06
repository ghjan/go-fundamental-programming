package main

import "testing"

type Foo int

func main() {

}

// 要性能
func (self *Foo) Get() Foo {
	return *self
}

// 要安全
func (self Foo) GetConst() Foo {
	return self
}

func TestGet(t *testing.T) {
	var foo Foo = 0
	if v := foo.Get(); v != 0 {
		t.Errorf("Bad Get. Need=%v, Got=%v", 0, v)
	}
}
func TestGetConst(t *testing.T) {
	var foo Foo = 0
	if v := foo.GetConst(); v != 0 {
		t.Errorf("Bad GetConst. Need=%v, Got=%v", 0, v)
	}
}

func BenchmarkGet(b *testing.B) {
	var foo Foo = 0
	for i := 0; i < b.N; i++ {
		_ = foo.Get()
	}
}
func BenchmarkGetConst(b *testing.B) {
	var foo Foo = 0
	for i := 0; i < b.N; i++ {
		_ = foo.GetConst()
	}
}
