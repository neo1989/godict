package main

import (
	"fmt"
	"godict/lib"
	"testing"
)

func Test(t *testing.T) {

	rt := trans.Trans("test")
	fmt.Println(rt)

	t.Log("测试通过")
}
