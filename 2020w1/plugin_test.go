package _2020w1

import (
	"plugin"
	"testing"
)

func TestPlugin(t *testing.T) {
	p, err := plugin.Open("./plugin/plugin_name.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())()
}
