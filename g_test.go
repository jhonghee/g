package g_test

import (
	"testing"

	"github.com/jhonghee/g"
)

func TestVersion(t *testing.T) {
	expected := "G v1.1.0->F v1.1.0"
	if g.Version() != expected {
		t.Error("Expected ", expected, "but got", g.Version())
	}
}
