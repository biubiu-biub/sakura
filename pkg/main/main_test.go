package main

import (
	"github.com/biubiu-biub/sakura/pkg/entry"
	"github.com/biubiu-biub/sakura/pkg/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntry(t *testing.T) {
	assert.Equal(t, entry.Ping(), "ping", "The two words should be the same.")
}
func TestSayName(t *testing.T){
	expected:=map[int]string{
		common.Cat:"cat",
		common.Bird:"bird",
		common.Dog:"dog",
		common.Fish:"fish",
		114514:"abaaba",
	}

    for tp,name:=range expected {
    	assert.Equal(t,sayName(tp),name,"returned value should be identical to excepted")
	}
}