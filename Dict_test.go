package main

import (
	"fmt"
	"godict/lib"
	"testing"
)

func Test(t *testing.T) {

	rt := trans.Trans("test")
	fmt.Println(rt)
	t.Error("测试没通过")
}
