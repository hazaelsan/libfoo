package libfoo

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	m := MakeProto(Foo)
	if got := m.Bar; got != "Some foo" {
		t.Errorf("MakeProto() got %v", got)
	}
}
