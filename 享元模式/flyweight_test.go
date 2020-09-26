package FlyWeight

import (
	"fmt"
	"testing"
)

func TestFlyWeightFactory_GetFlyWeight(t *testing.T) {
	factory := NewFlyWeightFactory()
	hong := factory.GetFlyWeight("hong beauty")
	xiang := factory.GetFlyWeight("xiang beauty")
	fmt.Println(hong)
	fmt.Println(xiang)
	fmt.Println(len(factory.pool))
}
