package main

import (
	"fmt"
	"strings"
	"testing"
)

type niceMan interface {
	SayHi() string
}

type man struct {
	name string
}

func (m *man) SayHi() string {
	return fmt.Sprint("Hi" + m.name)
}

func (m *man) String() string {
	return fmt.Sprint("Hi" + m.name)
}

func (m man) ChangeName() {
	m.name = "huu"
}

func saiHi(a niceMan) {
	fmt.Println(a.SayHi())
}

func TestSayHi(t *testing.T) {
	var m man
	m.name = "lane"
	saiHi(&m)
}

func TestAny(t *testing.T) {
	var b strings.Builder
	bb := new(strings.Builder)
	t.Logf("%T %T", b, bb)
}
