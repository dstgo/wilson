package base

import (
	"os"
	"testing"
)

func TestModuleVersion(t *testing.T) {
	v, err := ModuleVersion("golang.org/x/mod")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestModulePath(t *testing.T) {
	if err := os.Mkdir("/tmp/test_mod", os.ModePerm); err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("/tmp/test_mod/go.mod")
	if err != nil {
		t.Fatal(err)
	}

	mod := `module github.com/dstgo/kratosx/

go 1.19`
	_, err = f.WriteString(mod)
	if err != nil {
		t.Fatal(err)
	}

	p, err := ModulePath("/tmp/test_mod/go.mod")
	if err != nil {
		t.Fatal(err)
	}
	if p != "github.com/dstgo/kratosx/" {
		t.Fatalf("want: %s, got: %s", "github.com/dstgo/kratosx/", p)
	}

	t.Cleanup(func() { os.RemoveAll("/tmp/test_mod") })
}
