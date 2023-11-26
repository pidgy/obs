package dll

import "testing"

func TestDirectories(t *testing.T) {
	file, dir, err := Core("obs.dll")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("core dir:", dir)
	t.Log("core file:", file)

	file, dir, err = Module("win-capture.dll")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("module dir:", dir)
	t.Log("module file:", file)
}
