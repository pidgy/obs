package profiler

import (
	"testing"
)

func TestProfilerNameStore(t *testing.T) {
	n, err := New()
	if err != nil {
		t.Fatal(err)
	}

	name, err := n.Store("testing")
	if err != nil {
		t.Fatal(err)
	}
	if name != "testing" {
		t.Fatalf("unexpected name store: %s", name)
	}

	err = n.Close()
	if err != nil {
		t.Fatal(err)
	}

	n2, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	if n2.IsNull() {
		t.Fatalf("unexpected profiler name store")
	}
}
