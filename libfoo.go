// Package libfoo is a pretend Go library.
package libfoo

import (
	pb "github.com/hazaelsan/libfoo/proto/v1/foo"
)

// Foo is some dummy value.
const Foo string = "Some foo"

// MakeProto creates a proto message.
func MakeProto(s string) *pb.Foo {
	return &pb.Foo{Bar: s}
}
