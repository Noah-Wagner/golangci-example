package internal

import (
	"time"
)

//sumtype:decl
type ExampleSum interface {
	isSum()
}

type TypeA struct{}
type TypeB struct{}
type TypeC struct {
	Value time.Duration
}

func (r TypeA) isSum() {}
func (r TypeB) isSum() {}
func (r TypeC) isSum() {}
