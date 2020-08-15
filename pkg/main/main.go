package main

import (
	"fmt"
	"github.com/biubiu-biub/sakura/pkg/common"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	tp:=rand.Int()%4
	fmt.Println("我只是只 ",sayName(tp))
}
func sayName(tp int) string{
	switch  tp{
	case common.Dog:
		return "dog"
	case common.Cat:
		return "cat"
	case common.Bird:
		return "bird"
	case common.Fish:
		return "fish"
	default:
		return "abaaba"
	}
}