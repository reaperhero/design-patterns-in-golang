package Prototype

import (
	"fmt"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "name1"
	proto := ConcretePrototype{name: name}
	newProto := proto.Clone()
	newProto.SetName("name2")
	fmt.Printf("%p,%p", &proto, &newProto)           // 0xc0000f6000,0xc0000f6010 新对象
	fmt.Println(proto.GetName(), newProto.GetName()) // name1 name2
}
