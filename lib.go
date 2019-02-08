package gomodlib

import (
	"github.com/travisjeffery/gomodlib/subpkga"
	"github.com/travisjeffery/gomodlib/subpkgb"
)

type Service struct {
	A *subpkga.Service
	B *subpkgb.Service
}
