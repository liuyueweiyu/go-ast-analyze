package pkga

import (
	"fmt"
)

// A GetB
type A struct {
	Name string
	Age  int
}

var aEntity *A

// GetAEntity GetAEntity
func GetAEntity() *A {
	if aEntity == nil {
		aEntity = &A{Name: "hhh", Age: 10}
	}
	return aEntity
}
